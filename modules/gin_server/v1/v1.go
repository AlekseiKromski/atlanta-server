package v1

import (
	"alekseikromski.com/atlanta/modules/gin_server/guard"
	"alekseikromski.com/atlanta/modules/storage"
	"embed"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"path/filepath"
	"time"
)

type V1 struct {
	router  *gin.Engine
	storage storage.Storage
	log     func(messages ...string)
	secret  []byte
	guard   *guard.Guard
}

func NewV1Api(storage storage.Storage, secret []byte, cookieDomain string, log func(messages ...string)) *V1 {
	return &V1{
		router:  gin.Default(),
		storage: storage,
		log:     log,
		secret:  secret,
		guard:   guard.NewGuard(secret, storage, cookieDomain),
	}
}

func (v *V1) RegisterRoutes(resources embed.FS) error {
	v.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3001"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Content-Type, Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v.router.GET("/", v.application(resources))
	v.router.POST("/api/auth", v.guard.Auth)
	v.router.Static("/static", filepath.Join("front-end", "build", "static"))

	api := v.router.Group("/api").Use(v.guard.Check)
	{
		api.GET("/healthz", v.Healthz)
		api.GET("/datapoints/get-all", v.GetAllDatapoints(v.storage)) //TODO: remove
		api.GET("/datapoints/info/labels", v.GetAllLabels(v.storage))
		api.GET("/datapoints/info/devices", v.GetAllDevices(v.storage))
		api.POST("/datapoints/find", v.FindDatapoints(v.storage))
		api.POST("/users/create", v.CreateUser(v.storage))
	}

	return nil
}

func (v *V1) GetEngine() *gin.Engine {
	return v.router
}

func (v *V1) GetGuard() *guard.Guard {
	return v.guard
}

func (v *V1) application(resources embed.FS) func(c *gin.Context) {
	return func(c *gin.Context) {
		content, err := resources.ReadFile("front-end/build/index.html")
		if err != nil {
			log.Printf("cannot return index.html: %v", err)
			c.Status(500)
			return
		}

		c.Writer.Write(content)
	}
}
