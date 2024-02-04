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
	Time time.Time  `json:"time"` //Measurement time
}

// 56.098,23.090 - lat, lon
func (g *GeoDataPoint) ParseFromString(val string, measurementTime time.Time) error {
	value := strings.Split(val, ",")

	lat, err := strconv.ParseFloat(value[0], 32)
	if err != nil {
		return fmt.Errorf("cannot parse LAT in string: %s. Reason: %v", val, err)
	}

	lon, err := strconv.ParseFloat(value[1], 32)
	if err != nil {
		return fmt.Errorf("cannot parse LON in string: %s. Reason: %v", val, err)
	}

	g.Lat = &FloatData{
		Value: float32(lat),
	}
	g.Lon = &FloatData{
		Value: float32(lon),
	}
	g.Time = measurementTime

	return nil
}
