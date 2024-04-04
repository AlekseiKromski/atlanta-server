package v1

import (
	"alekseikromski.com/atlanta/modules/storage"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (v *V1) GetAllRoles(store storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, err := store.GetAllRoles()
		if err != nil {
			v.log("cannot get roles", err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse("cannot get roles"))
			return
		}

		c.JSON(http.StatusOK, roles)
	}
}

type upsertRoleRequest struct {
	Id        *string             `json:"id"`
	Name      string              `json:"name"`
	Endpoints []*storage.Endpoint `json:"endpoints"`
	DeletedAt *string             `json:"deleted_at"`
}

func (v *V1) UpsertRole(store storage.Storage) func(c *gin.Context) {
	return func(c *gin.Context) {
		defer c.Request.Body.Close()

		urr := &upsertRoleRequest{}
		if err := json.NewDecoder(c.Request.Body).Decode(&urr); err != nil {
			v.log("cannot decode upsert role request", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("wrong payload"))
			return
		}

		if urr.Id != nil && urr.DeletedAt == nil {
			// update
			endpointIds, err := store.GetEndpointIdsByRoleId(*urr.Id)
			if err != nil {
				v.log("cannot get permissions by role id", err.Error())
				c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse("cannot get permissions by role id"))
				return
			}

			for _, eid := range endpointIds {
				if err := store.DeletePermission(*urr.Id, eid); err != nil {
					v.log("cannot delete permission", err.Error())
					c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse("cannot permission"))
					return
				}
			}

			for _, eid := range urr.Endpoints {
				if err := store.CreatePermission(*urr.Id, eid.Id); err != nil {
					v.log("cannot create permission", err.Error())
					c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse("cannot create permission"))
					return
				}
			}

			if err := store.UpdateRole(*urr.Id, urr.Name); err != nil {
				v.log("cannot update role", err.Error())
				c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse("cannot update role"))
				return
			}

			c.Status(200)

		} else if urr.Id != nil && urr.DeletedAt != nil {
			// delete
			endpointIds, err := store.GetEndpointIdsByRoleId(*urr.Id)
			if err != nil {
				v.log("cannot get permissions by role id", err.Error())
				c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse("cannot get permissions by role id"))
				return
			}

			for _, eid := range endpointIds {
				if err := store.DeletePermission(*urr.Id, eid); err != nil {
					v.log("cannot delete permission", err.Error())
					c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse("cannot permission"))
					return
				}
			}

			if err := store.DeleteRole(*urr.Id); err != nil {
				v.log("cannot delete role", err.Error())
				c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse("cannot delete role"))
				return
			}

			c.Status(200)
		} else {
			// create
			roleId, err := store.CreateRole(urr.Name)
			if err != nil {
				v.log("cannot create role", err.Error())
				c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse("cannot create role"))
				return
			}

			for _, endpoint := range urr.Endpoints {
				if err := store.CreatePermission(*roleId, endpoint.Id); err != nil {
					v.log("cannot create permission", err.Error())
					c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrorResponse("cannot create permission"))
					return
				}
			}

			c.Status(200)
		}
	}
}
