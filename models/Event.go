package models

import "time"

type Event struct {
	GeoData     *GeoData    `json:"geo_data"`
	Device      *Device     `json:"device"`
	SensorData  *SensorData `json:"sensor_data"`
	Time        time.Time   `json:"time"`         // Measurement time (RFC3339 UTC+0)
	ReceiveTime *string     `json:"receive_time"` // Receiving time on the server (RFC 3339 UTC+0)
}
