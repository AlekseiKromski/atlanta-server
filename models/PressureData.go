package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type PressureData struct {
	Label    string     `json:"label"`
	Pressure *FloatData `json:"pressure"`
	Unit     string     `json:"unit"`
	Time     time.Time  `json:"time"` //Measurement time
	Flags    []string   `json:"flags"`
}

func (p *PressureData) ParseFromString(val string, measurementTime time.Time) error {
	value, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return fmt.Errorf("cannot parse PRS in string: %s. Reason: %v", val, err)
	}

	p.Label = "Pressure"
	p.Pressure = &FloatData{
		Value: float64(value),
		Type:  "float",
	}
	p.Unit = "Pa"
	p.Time = measurementTime

	return nil
}

func (p *PressureData) ToArguments() []any {
	return []any{
		strconv.FormatFloat(p.Pressure.Value, 'f', 6, 64),
		p.Pressure.Type,
		p.Unit,
		p.Time.Format(time.RFC3339),
		strings.Join(p.Flags, ","),
		p.Label,
	}
}

func (p *PressureData) Validate() {}
