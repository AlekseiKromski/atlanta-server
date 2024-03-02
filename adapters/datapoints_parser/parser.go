package datapoints_parser

import (
	"alekseikromski.com/atlanta/models"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type DataPointsParser struct{}

func NewDataPointsParser() *DataPointsParser {
	return &DataPointsParser{}
}

// DEVICE::3cc76ff4-cbaa-436c-b727-45d526facfc7;TIME::2019-10-12T07:20:50.52Z;TEMP::14C;PRS::1000PA
func (dpp *DataPointsParser) Parse(data string) (string, []models.DataPoints, error) {
	data = strings.TrimSuffix(data, "\n")
	data = strings.TrimSuffix(data, "\r")
	var datapoints []models.DataPoints
	incomingDatapoints := dpp.parseStringToMap(data)

	deviceUuid := incomingDatapoints["DEVICE"]

	incomingTime := incomingDatapoints["TIME"]
	parsedTime, err := dpp.parseTime(incomingTime)
	if err != nil {
		return deviceUuid, nil, fmt.Errorf("cannot parse time. Reason: %v", data)
	}
	measurementTime, err := time.Parse(time.RFC3339, parsedTime)
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

func (dpp *DataPointsParser) parseTime(timeData string) (string, error) {
	dividedTime := strings.Split(timeData, "T")

	if len(dividedTime) != 2 {
		return "", fmt.Errorf("cannot parse time data: %s", timeData)
	}

	dates := strings.Split(dividedTime[0], "-")
	times := strings.Split(dividedTime[1], ":")

	year, _ := strconv.Atoi(dates[0])
	month, _ := strconv.Atoi(dates[1])
	day, _ := strconv.Atoi(dates[2])

	hours, _ := strconv.Atoi(times[0])
	minutes, _ := strconv.Atoi(times[1])
	seconds, _ := strconv.Atoi((times[2])[:len(times[2])-1])

	fullDate := time.Date(year, time.Month(month), day, hours, minutes, seconds, 0, time.UTC)
	return fullDate.Format(time.RFC3339), nil
}

func (dpp *DataPointsParser) parseStringToMap(data string) map[string]string {
	dataPointList := strings.Split(data, ";")

	dataPoints := make(map[string]string)

	for _, dp := range dataPointList {
		dataPoint := strings.Split(dp, "::")
		if len(dataPoint) != 2 {
			return dataPoints
		}
		dataPoints[dataPoint[0]] = dataPoint[1]
	}

	return dataPoints
}
