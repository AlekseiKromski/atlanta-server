package models

import (
	"fmt"
	"strconv"
	"time"
)

type HumidityData struct {
	Humidity *IntegerData `json:"humidity"`
	Unit     string       `json:"unit"`
	Time     time.Time    `json:"time"` //Measurement time
}

func (h *HumidityData) ParseFromString(val string, measurementTime time.Time) error {
	value, err := strconv.ParseInt(val, 0, 64)
	if err != nil {
		return fmt.Errorf("cannot parse HUM in string: %s. Reason: %v", val, err)
	}

	h.Humidity = &IntegerData{
		Value: value,
		Type:  "int",
	}
	h.Unit = "percentage"
	h.Time = measurementTime

	return nil
}

func (h *HumidityData) ToArguments() []any {
	return []any{
		strconv.FormatInt(h.Humidity.Value, 2),
		h.Humidity.Type,
		h.Unit,
		h.Time.Format(time.RFC3339),
	}
}
