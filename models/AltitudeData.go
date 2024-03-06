package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type AltitudeData struct {
	Label    string     `json:"label"`
	Altitude *FloatData `json:"altitude"`
	Unit     string     `json:"unit"`
	Time     time.Time  `json:"time"` //Measurement time
	Flags    []string   `json:"flags"`
}

func (a *AltitudeData) ParseFromString(val string, measurementTime time.Time) error {
	if val[0] == '-' {
		val = val[1:len(val)]
	}
	value, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return fmt.Errorf("cannot parse ALT in string: %s. Reason: %v", val, err)
	}

	a.Label = "Altitude"
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
		strings.Join(a.Flags, ","),
		a.Label,
	}
}

func (a *AltitudeData) Validate() {}
