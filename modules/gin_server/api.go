package gin_server

import (
	"alekseikromski.com/atlanta/modules/gin_server/guard"
	"embed"
	"github.com/gin-gonic/gin"
)

type Api interface {
	RegisterRoutes(resources embed.FS) error
	GetEngine() *gin.Engine
	GetGuard() *guard.Guard
}
