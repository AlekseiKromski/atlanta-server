package v1

import (
	"alekseikromski.com/atlanta/modules/storage"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (v *V1) CreateUser(store storage.Storage) func(c *gin.Context) {
	return func(c *gin.Context) {
		user := &storage.User{}

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.Status(400)
			return
		}

		user.Password = string(hash)

		if err := c.ShouldBindJSON(&user); err != nil {
			c.Status(400)
			return
		}

		if err := store.CreateUser(user); err != nil {
			c.Status(400)
			return
		}

		c.Status(http.StatusOK)
	}
}
