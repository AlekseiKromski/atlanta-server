package v1

import (
	"alekseikromski.com/atlanta/modules/storage"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (v *V1) GetKVRecord(store storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("wrong id"))
			return
		}

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

		records, err := store.GetValues(id, val)
		if err != nil {
			v.log("cannot get value", err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse("cannot get value"))
			return
		}

		c.JSON(http.StatusOK, records)
	}
}

type kv_record_upsert_request struct {
	Key   string `json:"key"`
	Value string `json:"Value"`
}

func (v *V1) UpsertKVRecord(store storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer c.Request.Body.Close()

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

		kvrur := &kv_record_upsert_request{}
		if err := json.NewDecoder(c.Request.Body).Decode(&kvrur); err != nil {
			v.log("cannot decode kv record request", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("wrong payload"))
			return
		}

		if err := store.UpsertValue(kvrur.Key, kvrur.Value, val); err != nil {
			v.log("cannot upsert kv record", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("cannot upsert record"))
			return
		}

		c.JSON(http.StatusOK, struct {
			Key string
		}{
			Key: kvrur.Key,
		})
	}
}
