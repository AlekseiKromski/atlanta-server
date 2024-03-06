package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type HumidityData struct {
	Label    string     `json:"label"`
	Humidity *FloatData `json:"humidity"`
	Unit     string     `json:"unit"`
	Time     time.Time  `json:"time"` //Measurement time
	Flags    []string   `json:"flags"`
}

func (h *HumidityData) ParseFromString(val string, measurementTime time.Time) error {
	value, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return fmt.Errorf("cannot parse PRS in string: %s. Reason: %v", val, err)
	}

	h.Label = "Humidity"
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
		strings.Join(h.Flags, ","),
		h.Label,
	}
}

func (h *HumidityData) Validate() {
	if h.Humidity.Value > 100.00 || h.Humidity.Value < 0 {
		h.Flags = append(h.Flags, "ignored")
	}
}
