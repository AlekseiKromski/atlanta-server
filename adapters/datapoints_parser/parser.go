package datapoints_parser

import (
	"alekseikromski.com/atlanta/models"
	"fmt"
	"log"
	"strings"
	"time"
)

type DataPointsParser struct {
	data        string
	telemetries []models.DataPoints
}

func NewDataPointsParser(data string) *DataPointsParser {
	return &DataPointsParser{
		data: data,
	}
}

// TIME::2019-10-12T07:20:50.52Z;TEMP::14;PRS::1000PA
func (dpp *DataPointsParser) Parse() ([]models.DataPoints, error) {

	dataPoints := dpp.parseStringToMap()

	measurementTime, err := time.Parse(time.RFC3339, dataPoints["TIME"])
	if err != nil {
		return nil, fmt.Errorf("cannot parse time in data: %s. Reason: %v", dpp.data, err)
	}

	for key, val := range dataPoints {
		switch key {
		case "TIME":
			continue
		case "TEMP":
			td := &models.TemperatureData{}
			if err := td.ParseFromString(val, measurementTime); err != nil {
				return dpp.telemetries, fmt.Errorf("cannot parse temperature string: %v", err)
			}

			dpp.telemetries = append(dpp.telemetries, td)
		case "GEO":
			gd := &models.GeoDataPoint{}
			if err := gd.ParseFromString(val, measurementTime); err != nil {
				return dpp.telemetries, fmt.Errorf("cannot parse geo string: %v", err)
			}

			dpp.telemetries = append(dpp.telemetries, gd)
		case "PRS":
			pd := &models.PressureData{}
			if err := pd.ParseFromString(val, measurementTime); err != nil {
				return dpp.telemetries, fmt.Errorf("cannot parse pressure string: %v", err)
			}

			dpp.telemetries = append(dpp.telemetries, pd)
		case "ALT":
			ad := &models.AltitudeData{}
			if err := ad.ParseFromString(val, measurementTime); err != nil {
				return dpp.telemetries, fmt.Errorf("cannot parse altitude string: %v", err)
			}

			dpp.telemetries = append(dpp.telemetries, ad)
		case "HUM":
			hd := &models.HumidityData{}
			if err := hd.ParseFromString(val, measurementTime); err != nil {
				return dpp.telemetries, fmt.Errorf("cannot parse humidity string: %v", err)
			}

			dpp.telemetries = append(dpp.telemetries, hd)
		default:
			log.Printf("cannot parse key-value: %s/%s. Incoming string: %s", key, val, dpp.data)
		}
	}

	return dpp.telemetries, nil
}

func (dpp *DataPointsParser) parseStringToMap() map[string]string {
	dataPointList := strings.Split(dpp.data, ";")

	dataPoints := make(map[string]string)

	for _, dp := range dataPointList {
		dataPoint := strings.Split(dp, "::")
		dataPoints[dataPoint[0]] = dataPoint[1]
	}

	return dataPoints
}
