package datapoints_parser

import (
	"alekseikromski.com/atlanta/models"
	"fmt"
	"log"
	"strings"
	"time"
)

type DataPointsParser struct{}

func NewDataPointsParser() *DataPointsParser {
	return &DataPointsParser{}
}

// DEVICE::3cc76ff4-cbaa-436c-b727-45d526facfc7;TIME::2019-10-12T07:20:50.52Z;TEMP::14C;PRS::1000PA
func (dpp *DataPointsParser) Parse(data string) (string, []models.DataPoints, error) {
	var datapoints []models.DataPoints
	incomingDatapoints := dpp.parseStringToMap(data)

	deviceUuid := incomingDatapoints["DEVICE"]

	measurementTime, err := time.Parse(time.RFC3339, incomingDatapoints["TIME"])
	if err != nil {
		return deviceUuid, nil, fmt.Errorf("cannot parse time in data: %s. Reason: %v", data, err)
	}

	for key, val := range incomingDatapoints {
		switch key {
		case "TIME": // Measurement timestamp
			continue
		case "DEVICE": // Device uuid
			continue
		case "TEMP":
			td := &models.TemperatureData{}
			if err := td.ParseFromString(val, measurementTime); err != nil {
				return deviceUuid, nil, fmt.Errorf("cannot parse temperature string: %v", err)
			}

			datapoints = append(datapoints, td)
		case "GEO":
			gd := &models.GeoDataPoint{}
			if err := gd.ParseFromString(val, measurementTime); err != nil {
				return deviceUuid, nil, fmt.Errorf("cannot parse geo string: %v", err)
			}

			datapoints = append(datapoints, gd)
		case "PRS":
			pd := &models.PressureData{}
			if err := pd.ParseFromString(val, measurementTime); err != nil {
				return deviceUuid, nil, fmt.Errorf("cannot parse pressure string: %v", err)
			}

			datapoints = append(datapoints, pd)
		case "ALT":
			ad := &models.AltitudeData{}
			if err := ad.ParseFromString(val, measurementTime); err != nil {
				return deviceUuid, nil, fmt.Errorf("cannot parse altitude string: %v", err)
			}

			datapoints = append(datapoints, ad)
		case "HUM":
			hd := &models.HumidityData{}
			if err := hd.ParseFromString(val, measurementTime); err != nil {
				return deviceUuid, nil, fmt.Errorf("cannot parse humidity string: %v", err)
			}

			datapoints = append(datapoints, hd)
		default:
			log.Printf("cannot parse key-value: %s/%s. Incoming string: %s", key, val, data)
		}
	}

	return deviceUuid, datapoints, nil
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
