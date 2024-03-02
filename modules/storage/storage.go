package storage

import "alekseikromski.com/atlanta/models"

type Storage interface {
	SaveDatapoints(deviceUuid string, datapoints []models.DataPoints) error // SaveDatapoints - Save datapoint to database

	StorageUser      // StorageUser - crud interface for working with users
	StorageDatapoint // StorageDatapoint - all commands related to datapoints
}
