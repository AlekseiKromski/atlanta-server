package v1

import (
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
}

func NewV1Api(storage storage.Storage, secret []byte, log func(messages ...string)) *V1 {
	return &V1{
		router:  gin.Default(),
		storage: storage,
		log:     log,
		secret:  secret,
	}
}

func (v *V1) RegisterRoutes(resources embed.FS) error {
	v.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v.router.GET("/", v.application(resources))
	v.router.Static("/static", filepath.Join("front-end", "build", "static"))

	api := v.router.Group("/api")
	{
		api.GET("/healthz", v.Healthz)
		api.GET("/datapoint/get-all", v.GetAllDatapoints(v.storage)) //TODO: remove
		api.GET("/datapoints/info/labels", v.GetAllLabels(v.storage))
		api.POST("/datapoints/find", v.FindDatapoints(v.storage))
		api.POST("/auth", v.Auth)
	}
	users := v.router.Group("/users")
	{
		users.POST("/create", v.CreateUser(v.storage))
		users.GET("/get/:id", v.GetUser(v.storage))
	}
	return nil
}

func (v *V1) GetEngine() *gin.Engine {
	return v.router
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
