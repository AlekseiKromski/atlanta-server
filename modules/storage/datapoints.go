package storage

import (
	"alekseikromski.com/atlanta/models"
	"github.com/google/uuid"
	"time"
)

type Datapoint struct {
	ID              uuid.UUID `json:"id"`
	DeviceId        uuid.UUID `json:"device_id"`
	Value           string    `json:"value"`
	ValueType       string    `json:"type"`
	Unit            string    `json:"unit"`
	Label           *string   `json:"label"`
	Flags           *string   `json:"flags"`
	MeasurementTime time.Time `json:"measurement_time"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type StorageDatapoint interface {
	SaveDatapoints(deviceUuid string, datapoints []models.DataPoints) ([]*Datapoint, error) // SaveDatapoints - Save datapoint to database
	FindDatapoints(fd *FindDatapointsRequest) ([]*Datapoint, []string, error)
	FindAllLabels() ([]string, error)
	FindAllDevices() ([]*Device, error)
}
