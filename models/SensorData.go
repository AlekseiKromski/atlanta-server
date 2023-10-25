package models

import "time"

type SensorData struct {
	Pressure    FloatData `json:"pressure"`
	Temperature FloatData `json:"temperature"`
}

type FloatData struct {
	Value *float64  `json:"value"`
	Time  time.Time `json:"time"` //Measurement time
}

func (sd *SensorData) Validate() bool {
	return sd.Pressure.Value != nil && sd.Temperature.Value != nil
}
