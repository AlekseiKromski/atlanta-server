package gin_server

import "github.com/gin-gonic/gin"

type Api interface {
	RegisterRoutes() error
	GetEngine() *gin.Engine
}
