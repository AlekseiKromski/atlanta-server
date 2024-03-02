package v1

import (
	"alekseikromski.com/atlanta/modules/storage"
	"github.com/gin-gonic/gin"
)

func (v *V1) GetAllDatapoints(store storage.Storage) func(c *gin.Context) {
	return func(c *gin.Context) {
		dps, err := store.GetAllDatapoints()
		if err != nil {
			c.Status(400)
			return
		}

		c.JSON(200, dps)
	}
}
