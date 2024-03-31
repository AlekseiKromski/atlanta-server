package storage

import "time"

type Device struct {
	Id          string     `json:"id"`
	Description string     `json:"description"`
	Status      bool       `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type StorageDevices interface {
	GetDevices() ([]*Device, error)
	CreateDevice(description string, status bool) error
	UpdateDevice(id, description string, status bool) error
	DeleteDevice(id string) error
}
