package models

import (
	"fmt"
	"strconv"
	"time"
)

type AltitudeData struct {
	Altitude *FloatData `json:"altitude"`
	Unit     string     `json:"unit"`
	Time     time.Time  `json:"time"` //Measurement time
}

func (a *AltitudeData) ParseFromString(val string, measurementTime time.Time) error {
	value, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return fmt.Errorf("cannot parse ALT in string: %s. Reason: %v", val, err)
	}

	a.Altitude = &FloatData{
		Value: float64(value),
		Type:  "float",
	}
	a.Unit = "M"
	a.Time = measurementTime

	return nil
}

func (a *AltitudeData) ToArguments() []any {
	return []any{
		strconv.FormatFloat(a.Altitude.Value, 'f', 6, 64),
		a.Altitude.Type,
		a.Unit,
		a.Time.Format(time.RFC3339),
	}
}
