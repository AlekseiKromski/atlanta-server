package storage

import (
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
	GetAllDatapoints() ([]*Datapoint, error)
}
