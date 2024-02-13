package storage

import "alekseikromski.com/atlanta/models"

type Storage interface {
	SaveDatapoints(deviceUuid string, datapoints []models.DataPoints) error // Save datapoint to database
}
