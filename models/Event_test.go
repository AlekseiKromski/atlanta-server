package models

import (
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"os"
	"testing"
	"time"
)

func TestEventDecoding(t *testing.T) {
	file, err := os.ReadFile("../examples/example1.json")
	if err != nil {
		t.Fatalf("cannot open file: %v", err)
		return
	}

	var event Event
	if err := json.Unmarshal(file, &event); err != nil {
		t.Fatalf("cannot decode file: %v", err)
		return
	}

	assert.Equal(t, event.Device.Uuid, "6b2ad11d-f83a-4b40-bb9f-f5d316e4871d")
	assert.Equal(t, event.Device.ApiKey, "123")
	assert.Equal(t, event.Device.Battery, 43.0)

	assert.Equal(t, event.GeoData.Lat, 0.0)
	assert.Equal(t, event.GeoData.Lon, 0.0)

	stime, err := time.Parse(time.RFC3339, "2019-10-12T07:20:50.52Z")
	if err != nil {
		t.Fatalf("cannot decode string to date: %v", err)
		return
	}
	assert.Equal(t, event.SensorData.Pressure.Value, 100.0)
	assert.Equal(t, event.SensorData.Pressure.Time, stime)
	assert.Equal(t, event.SensorData.Temperature.Value, 28.0)
	assert.Equal(t, event.SensorData.Temperature.Time, stime)

	mtime, err := time.Parse(time.RFC3339, "2019-10-12T07:20:51.52Z")
	if err != nil {
		t.Fatalf("cannot decode string to date: %v", err)
		return
	}

	assert.Equal(t, event.Time, mtime)

	assert.Equal(t, true, event.Validate())
}
