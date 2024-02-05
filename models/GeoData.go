package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type GeoDataPoint struct {
	Lat  *FloatData `json:"lat"`
	Lon  *FloatData `json:"lon"`
	Unit string     `json:"unit"`
	Time time.Time  `json:"time"` //Measurement time
}

// 56.098,23.090 - lat, lon
func (g *GeoDataPoint) ParseFromString(val string, measurementTime time.Time) error {
	value := strings.Split(val, ",")

	lat, err := strconv.ParseFloat(value[0], 64)
	if err != nil {
		return fmt.Errorf("cannot parse LAT in string: %s. Reason: %v", val, err)
	}

	lon, err := strconv.ParseFloat(value[1], 64)
	if err != nil {
		return fmt.Errorf("cannot parse LON in string: %s. Reason: %v", val, err)
	}

	geoType := "float"
	g.Lat = &FloatData{
		Value: float64(lat),
		Type:  geoType,
	}
	g.Lon = &FloatData{
		Value: float64(lon),
		Type:  geoType,
	}
	g.Unit = "coordinate"
	g.Time = measurementTime

	return nil
}

func (g *GeoDataPoint) ToArguments() []any {
	return []any{
		fmt.Sprintf("%f,%f", g.Lat.Value, g.Lon.Value),
		g.Lat.Type, // same type for both values
		g.Unit,
		g.Time.Format(time.RFC3339),
	}
}
