package v1

import (
	"alekseikromski.com/atlanta/modules/storage"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (v *V1) GetDevices(store storage.Storage) func(c *gin.Context) {
	return func(c *gin.Context) {
		dps, err := store.GetDevices()
		if err != nil {
			v.log("cannot get devices from database", err.Error())
			c.JSON(400, NewErrorResponse("cannot get devices"))
			return
		}

		c.JSON(200, dps)
	}
}

type upsertDeviceRequest struct {
	Id          *string `json:"id"`
	Description string  `json:"description"`
	Status      bool    `json:"status"`
}

func (v *V1) UpsertDevice(store storage.Storage) func(c *gin.Context) {
	return func(c *gin.Context) {
		defer c.Request.Body.Close()

		upr := &upsertDeviceRequest{}
		if err := json.NewDecoder(c.Request.Body).Decode(&upr); err != nil {
			v.log("cannot decode upsert device request", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("wrong payload"))
			return
		}

		if len(upr.Description) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("wrong payload"))
			return
		}

		var err error
		if upr.Id == nil {
			//create
			err = store.CreateDevice(upr.Description, upr.Status)
		} else {
			//update
			err = store.UpdateDevice(*upr.Id, upr.Description, upr.Status)
		}

		if err != nil {
			v.log("cannot upsert device to db", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("cannot upsert device"))
			return
		}

		c.Status(http.StatusOK)
	}
}

func (v *V1) DeleteDevice(store storage.Storage) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("wrong id"))
			return
		}

		if err := store.DeleteDevice(id); err != nil {
			v.log("cannot delete device", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("cannot delete device"))
			return
		}

		c.Status(http.StatusOK)
	}
}
