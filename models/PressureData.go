package models

import (
	"fmt"
	"strconv"
	"time"
)

type PressureData struct {
	Pressure *FloatData `json:"pressure"`
	Time     time.Time  `json:"time"` //Measurement time
}

func (p *PressureData) ParseFromString(val string, measurementTime time.Time) error {
	value, err := strconv.ParseFloat(val[:len(val)-2], 32)
	if err != nil {
		return fmt.Errorf("cannot parse PRS in string: %s. Reason: %v", val, err)
	}

	p.Pressure = &FloatData{
		Value: float32(value),
	}
	p.Time = measurementTime

	return nil
}
