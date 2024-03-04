package models

import (
	"fmt"
	"strconv"
	"time"
)

type HumidityData struct {
	Humidity *FloatData `json:"humidity"`
	Unit     string     `json:"unit"`
	Time     time.Time  `json:"time"` //Measurement time
}

func (h *HumidityData) ParseFromString(val string, measurementTime time.Time) error {
	value, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return fmt.Errorf("cannot parse PRS in string: %s. Reason: %v", val, err)
	}

	h.Humidity = &FloatData{
		Value: float64(value),
		Type:  "float",
	}
	h.Unit = "percentage"
	h.Time = measurementTime

	return nil
}

func (h *HumidityData) ToArguments() []any {
	return []any{
		strconv.FormatFloat(h.Humidity.Value, 'f', 6, 64),
		h.Humidity.Type,
		h.Unit,
		h.Time.Format(time.RFC3339),
	}
}
