package v1

import (
	"alekseikromski.com/atlanta/modules/storage"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (v *V1) GetAllEndpoints(store storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		endpoints, err := store.GetAllEndpoints()
		if err != nil {
			v.log("cannot get endpoints", err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse("cannot get endpoints"))
			return
		}

		c.JSON(http.StatusOK, endpoints)
	}
}

type upsertEndpointRequest struct {
	Id          *string `json:"id"`
	Urn         string  `json:"urn"`
	Description string  `json:"description"`
	DeletedAt   *string `json:"deleted_at"`
}

func (v *V1) UpsertEndpoint(store storage.Storage) func(c *gin.Context) {
	return func(c *gin.Context) {
		defer c.Request.Body.Close()

		uer := &upsertEndpointRequest{}
		if err := json.NewDecoder(c.Request.Body).Decode(&uer); err != nil {
			v.log("cannot decode upsert endpoint request", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("wrong payload"))
			return
		}

		if uer.Id != nil && uer.DeletedAt == nil {
			// update
			if err := store.UpdateEndpoint(*uer.Id, uer.Urn, uer.Description); err != nil {
				v.log("cannot update endpoint", err.Error())
				c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse("cannot update endpoint"))
				return
			}

			c.Status(200)

		} else if uer.Id != nil && uer.DeletedAt != nil {
			// delete
			if err := store.DeleteEndpoint(*uer.Id); err != nil {
				v.log("cannot delete endpoint", err.Error())
				c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse("cannot delete endpoint"))
				return
			}

			c.Status(200)
		} else {
			// create
			if err := store.CreateEndpoint(uer.Urn, uer.Description); err != nil {
				v.log("cannot create endpoint", err.Error())
				c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse("cannot create endpoint"))
				return
			}

			c.Status(200)
		}
	}
}
