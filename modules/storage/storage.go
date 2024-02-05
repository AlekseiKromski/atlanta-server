package storage

import "alekseikromski.com/atlanta/models"

type Storage interface {
	SaveDatapoints([]models.DataPoints) error // Save datapoint to database
}
