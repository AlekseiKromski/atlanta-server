package gin_server

import (
	"embed"
	"github.com/gin-gonic/gin"
)

type Api interface {
	RegisterRoutes(resources embed.FS) error
	GetEngine() *gin.Engine
}
