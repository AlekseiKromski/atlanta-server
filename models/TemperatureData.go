package models

import (
	"fmt"
	"strconv"
	"time"
)

type TemperatureData struct {
	Temperature *FloatData `json:"temperature"`
	Unit        string     `json:"unit"`
	Time        time.Time  `json:"time"` //Measurement time
}

func (t *TemperatureData) ParseFromString(val string, measurementTime time.Time) error {
	index := len(val) - 1
	tempType := val[index]
	tempValue, err := strconv.ParseFloat(string(val[0:index]), 64)
	if err != nil {
		return fmt.Errorf("cannot parse TEMP in string: %s. Reason: %v", val, err)
	}

	t.Temperature = &FloatData{
		Value: float64(tempValue),
		Type:  "float",
	}
	t.Unit = string(tempType)
	t.Time = measurementTime

	return nil
}

func (t *TemperatureData) ToArguments() []any {
	return []any{
		strconv.FormatFloat(t.Temperature.Value, 'f', 6, 64),
		t.Temperature.Type,
		t.Unit,
		t.Time.Format(time.RFC3339),
	}
}
