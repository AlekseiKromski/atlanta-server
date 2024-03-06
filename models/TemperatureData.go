package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type TemperatureData struct {
	Label       string     `json:"label"`
	Temperature *FloatData `json:"temperature"`
	Unit        string     `json:"unit"`
	Time        time.Time  `json:"time"` //Measurement time
	Flags       []string   `json:"flags"`
}

func (t *TemperatureData) ParseFromString(val string, measurementTime time.Time) error {
	tempValue, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return fmt.Errorf("cannot parse TEMP in string: %s. Reason: %v", val, err)
	}

	t.Label = "Temperature"
	t.Temperature = &FloatData{
		Value: float64(tempValue),
		Type:  "float",
	}
	t.Unit = "C"
	t.Time = measurementTime

	return nil
}

func (t *TemperatureData) ToArguments() []any {
	return []any{
		strconv.FormatFloat(t.Temperature.Value, 'f', 6, 64),
		t.Temperature.Type,
		t.Unit,
		t.Time.Format(time.RFC3339),
		strings.Join(t.Flags, ","),
		t.Label,
	}
}

func (t *TemperatureData) Validate() {
	if t.Temperature.Value >= 50 {
		t.Flags = append(t.Flags, "ignored")
	}
}
