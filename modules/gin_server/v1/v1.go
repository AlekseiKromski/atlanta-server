package v1

import (
	"alekseikromski.com/atlanta/modules/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

type V1 struct {
	router  *gin.Engine
	storage storage.Storage
}

func NewV1Api(storage storage.Storage) *V1 {
	return &V1{
		router:  gin.Default(),
		storage: storage,
	}
}

func (v *V1) RegisterRoutes() error {
	v.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	v.router.GET("/healthz", v.Healthz)
	v.router.GET("/datapoint/get-all", v.GetAllDatapoints(v.storage))
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
