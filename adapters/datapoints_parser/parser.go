package datapoints_parser

import (
	"alekseikromski.com/atlanta/models"
	"fmt"
	"log"
	"strings"
	"time"
)

type DataPointsParser struct {
	Datapoints []models.DataPoints
}

func NewDataPointsParser() *DataPointsParser {
	return &DataPointsParser{
		Datapoints: []models.DataPoints{},
	}
}

// TIME::2019-10-12T07:20:50.52Z;TEMP::14;PRS::1000PA
func (dpp *DataPointsParser) Parse(data string) error {

	dataPoints := dpp.parseStringToMap(data)

	measurementTime, err := time.Parse(time.RFC3339, dataPoints["TIME"])
	if err != nil {
		return fmt.Errorf("cannot parse time in data: %s. Reason: %v", data, err)
	}

	for key, val := range dataPoints {
		switch key {
		case "TIME":
			continue
		case "TEMP":
			td := &models.TemperatureData{}
			if err := td.ParseFromString(val, measurementTime); err != nil {
				return fmt.Errorf("cannot parse temperature string: %v", err)
			}

			dpp.Datapoints = append(dpp.Datapoints, td)
		case "GEO":
			gd := &models.GeoDataPoint{}
			if err := gd.ParseFromString(val, measurementTime); err != nil {
				return fmt.Errorf("cannot parse geo string: %v", err)
			}

			dpp.Datapoints = append(dpp.Datapoints, gd)
		case "PRS":
			pd := &models.PressureData{}
			if err := pd.ParseFromString(val, measurementTime); err != nil {
				return fmt.Errorf("cannot parse pressure string: %v", err)
			}

			dpp.Datapoints = append(dpp.Datapoints, pd)
		case "ALT":
			ad := &models.AltitudeData{}
			if err := ad.ParseFromString(val, measurementTime); err != nil {
				return fmt.Errorf("cannot parse altitude string: %v", err)
			}

			dpp.Datapoints = append(dpp.Datapoints, ad)
		case "HUM":
			hd := &models.HumidityData{}
			if err := hd.ParseFromString(val, measurementTime); err != nil {
				return fmt.Errorf("cannot parse humidity string: %v", err)
			}

			dpp.Datapoints = append(dpp.Datapoints, hd)
		default:
			log.Printf("cannot parse key-value: %s/%s. Incoming string: %s", key, val, data)
		}
	}

	return nil
}

func (dpp *DataPointsParser) parseStringToMap(data string) map[string]string {
	dataPointList := strings.Split(data, ";")

	dataPoints := make(map[string]string)

	for _, dp := range dataPointList {
		dataPoint := strings.Split(dp, "::")
		dataPoints[dataPoint[0]] = dataPoint[1]
	}

	return dataPoints
}
