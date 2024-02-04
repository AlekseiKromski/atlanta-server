package models

import (
	"fmt"
	"strconv"
	"time"
)

type AltitudeData struct {
	Altitude *FloatData `json:"altitude"`
	Time     time.Time  `json:"time"` //Measurement time
}

func (a *AltitudeData) ParseFromString(val string, measurementTime time.Time) error {
	value, err := strconv.ParseFloat(val, 32)
	if err != nil {
		return fmt.Errorf("cannot parse ALT in string: %s. Reason: %v", val, err)
	}

	a.Altitude = &FloatData{
		Value: float32(value),
	}
	a.Time = measurementTime

	return nil
}
