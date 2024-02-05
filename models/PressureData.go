package models

import (
	"fmt"
	"strconv"
	"time"
)

type PressureData struct {
	Pressure *FloatData `json:"pressure"`
	Unit     string     `json:"unit"`
	Time     time.Time  `json:"time"` //Measurement time
}

func (p *PressureData) ParseFromString(val string, measurementTime time.Time) error {
	value, err := strconv.ParseFloat(val[:len(val)-2], 64)
	if err != nil {
		return fmt.Errorf("cannot parse PRS in string: %s. Reason: %v", val, err)
	}

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
	}
}
