package v1

import (
	"alekseikromski.com/atlanta/modules/storage"
	"github.com/gin-gonic/gin"
)

func (v *V1) CreateUser(store storage.Storage) func(c *gin.Context) {
	return func(c *gin.Context) {
		user := &storage.User{}

		if err := c.ShouldBindJSON(&user); err != nil {
			c.Status(400)
			return
		}

		if err := store.CreateUser(user); err != nil {
			c.Status(400)
			return
		}
	}
}

func (v *V1) GetUser(store storage.Storage) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")

		user, err := store.GetUser(id)
		if err != nil {
			c.Status(400)
			return
		}

		user.Password = nil

		c.JSON(200, user)
	}
}
