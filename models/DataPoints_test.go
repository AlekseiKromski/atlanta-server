package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTemperatureData_ParseFromStringCelsius(t *testing.T) {
	str := "30.0C"
	mtime := time.Now()

	td := TemperatureData{}
	err := td.ParseFromString(str, mtime)
	assert.NoError(t, err)

	assert.Equal(t, float32(30.0), td.Temperature.Value)
	assert.Equal(t, "C", td.Temperature.Type)
	assert.Equal(t, mtime, td.Time)
}

func TestTemperatureData_ParseFromStringFahrenheit(t *testing.T) {
	str := "30.1F"
	mtime := time.Now()

	td := TemperatureData{}
	err := td.ParseFromString(str, mtime)
	assert.NoError(t, err)

	assert.Equal(t, float32(30.1), td.Temperature.Value)
	assert.Equal(t, "F", td.Temperature.Type)
	assert.Equal(t, mtime, td.Time)
}
