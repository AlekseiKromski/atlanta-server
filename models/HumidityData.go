package models

import (
	"fmt"
	"strconv"
	"time"
)

type HumidityData struct {
	Humidity *FloatData `json:"humidity"`
	Time     time.Time  `json:"time"` //Measurement time
}

func (h *HumidityData) ParseFromString(val string, measurementTime time.Time) error {
	value, err := strconv.ParseFloat(val, 32)
	if err != nil {
		return fmt.Errorf("cannot parse PRS in string: %s. Reason: %v", val, err)
	}

	h.Humidity = &FloatData{
		Value: float32(value),
	}
	h.Time = measurementTime

	return nil
}
