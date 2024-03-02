package datapoints_parser

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser(t *testing.T) {
	deviceUuidExpected := "3cc76ff4-cbaa-436c-b727-45d526facfc7"
	data := fmt.Sprintf("DEVICE::%s;TIME::2019-10-12T07:20:50.52Z;TEMP::14", deviceUuidExpected)
	parser := NewDataPointsParser()
	deviceUuid, datapoints, err := parser.Parse(data)
	if err != nil {
		t.Fatalf("cannot parser data points: %v", err)
		return
	}

	actual, err := json.Marshal(datapoints[0])
	if err != nil {
		t.Fatalf("cannot marshal: %v", err)
		return
	}
	expected := "{\"temperature\":{\"value\":14,\"type\":\"float\"},\"unit\":\"C\",\"time\":\"2019-10-12T07:20:00Z\"}"

	assert.Equal(t, expected, string(actual))
	assert.Equal(t, deviceUuidExpected, deviceUuid)
}
