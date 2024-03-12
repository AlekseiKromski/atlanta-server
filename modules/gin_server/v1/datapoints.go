package v1

import (
	"alekseikromski.com/atlanta/modules/storage"
	"encoding/json"
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

func (v *V1) FindDatapoints(store storage.Storage) func(c *gin.Context) {
	return func(c *gin.Context) {
		defer c.Request.Body.Close()

		findReq := &storage.FindDatapoints{}
		if err := json.NewDecoder(c.Request.Body).Decode(findReq); err != nil {
			c.Status(400)
			return
		}

		if ok := findReq.Validate(); !ok {
			c.Status(400)
			return
		}

		dps, labels, err := store.FindDatapoints(findReq)
		if err != nil {
			c.Status(400)
			return
		}

		//Group by label
		grouped := map[string][]*storage.Datapoint{}
		for _, datapoint := range dps {
			if grouped[*datapoint.Label] == nil {
				grouped[*datapoint.Label] = []*storage.Datapoint{datapoint}
				continue
			}
			grouped[*datapoint.Label] = append(grouped[*datapoint.Label], datapoint)
		}

		c.JSON(200, struct {
			Datapoints map[string][]*storage.Datapoint `json:"datapoints"`
			Labels     []string                        `json:"labels"`
		}{
			Datapoints: grouped,
			Labels:     labels,
		})
	}
}

func (v *V1) GetAllLabels(store storage.Storage) func(c *gin.Context) {
	return func(c *gin.Context) {
		labels, err := store.FindAllLabels()
		if err != nil {
			c.Status(400)
			return
		}

		c.JSON(200, labels)
	}
}
