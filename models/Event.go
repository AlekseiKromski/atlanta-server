package models

import "time"

type Validator interface {
	Validate() bool
}

type Event struct {
	GeoData    GeoData    `json:"geo_data"`
	Device     Device     `json:"device"`
	SensorData SensorData `json:"sensor_data"`
	Time       time.Time  `json:"time"` // Measurement time (RFC3339 UTC+0)
}

func (e *Event) Validate() bool {
	return e.GeoData.Validate() && e.Device.Validate() && e.SensorData.Validate()
}
