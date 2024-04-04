package v1

import (
	"alekseikromski.com/atlanta/modules/storage"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (v *V1) GetCurrentUser(store storage.Storage) func(c *gin.Context) {
	return func(c *gin.Context) {
		userId, exist := c.Get("uid")
		if !exist {
			v.log("cannot get current user form request")
			c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse("cannot get current user form request"))
			return
		}

		val, ok := userId.(string)
		if !ok {
			v.log("wrong user id")
			c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse("wrong user id"))
			return
		}

		user, err := store.GetUserByIdWithEndpoints(val)
		if err != nil {
			v.log("cannot find user", err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse("cannot find user"))
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
func (v *V1) GetAllUsers(store storage.Storage) func(c *gin.Context) {
	return func(c *gin.Context) {
		users, err := store.GetAllUsers()
		if err != nil {
			v.log("cannot get users", err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse("cannot get users"))
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

type upsertUserRequest struct {
	Id          *string `json:"id"`
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	First_name  string  `json:"first_name"`
	Second_name string  `json:"second_name"`
	Password    string  `json:"password,omitempty"`
	Role        string  `json:"role"`
	DeletedAt   *string `json:"deleted_at"`
}

func (v *V1) UpsertUser(store storage.Storage) func(c *gin.Context) {
	return func(c *gin.Context) {
		defer c.Request.Body.Close()

		uur := &upsertUserRequest{}
		if err := json.NewDecoder(c.Request.Body).Decode(&uur); err != nil {
			v.log("cannot decode upsert user request", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("wrong payload"))
			return
		}

		if uur.Id != nil && uur.DeletedAt == nil {
			// update

			user, err := store.GetUserById(*uur.Id)
			if err != nil {
				v.log("cannot get user", err.Error())
				c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse("cannot get user"))
				return
			}

			hash_password := user.Password
			if len(uur.Password) != 0 {
				hash, err := bcrypt.GenerateFromPassword([]byte(uur.Password), bcrypt.DefaultCost)
				if err != nil {
					v.log("cannot generate password", err.Error())
					c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse("cannot generate password"))
					return
				}
				hash_password = string(hash)
			}

			if err := store.UpdateUser(*uur.Id, uur.Username, uur.Email, uur.First_name, uur.Second_name, hash_password, uur.Role); err != nil {
				v.log("cannot update user", err.Error())
				c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse("cannot update user"))
				return
			}

			c.Status(200)

		} else if uur.Id != nil && uur.DeletedAt != nil {
			// delete
			if err := store.DeleteUser(*uur.Id); err != nil {
				v.log("cannot delete user", err.Error())
				c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse("cannot create user"))
				return
			}

			c.Status(200)
		} else {
			// create
			hash, err := bcrypt.GenerateFromPassword([]byte(uur.Password), bcrypt.DefaultCost)
			if err != nil {
				c.Status(400)
				return
			}

			if err := store.CreateUser(uur.Username, uur.Email, uur.First_name, uur.Second_name, string(hash), uur.Role); err != nil {
				v.log("cannot create user", err.Error())
				c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse("cannot delete user"))
				return
			}

			c.Status(200)
		}
	}
}
