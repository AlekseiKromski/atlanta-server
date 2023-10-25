package models

import "time"

type SensorData struct {
	Pressure    *FloatData `json:"pressure"`
	Temperature *FloatData `json:"temperature"`
}

type FloatData struct {
	Value float32   `json:"value"`
	Time  time.Time `json:"time"` //Measurement time
}
