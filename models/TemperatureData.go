package models

import (
	"fmt"
	"strconv"
	"time"
)

type TemperatureData struct {
	Temperature *FloatData `json:"temperature"`
	Time        time.Time  `json:"time"` //Measurement time
}

func (t *TemperatureData) ParseFromString(val string, measurementTime time.Time) error {
	index := len(val) - 1
	tempType := val[index]
	tempValue, err := strconv.ParseFloat(string(val[0:index]), 32)
	if err != nil {
		return fmt.Errorf("cannot parse TEMP in string: %s. Reason: %v", val, err)
	}

	t.Temperature = &FloatData{
		Value: float32(tempValue),
		Type:  string(tempType),
	}
	t.Time = measurementTime

	return nil
}
