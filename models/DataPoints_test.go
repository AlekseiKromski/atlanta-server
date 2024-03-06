package models

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestTemperatureData_ParseFromString(t *testing.T) {
	str := "30.0"
	mtime := time.Now().UTC()

	td := TemperatureData{}
	err := td.ParseFromString(str, mtime)
	assert.NoError(t, err)

	assert.Equal(t, 30.0, td.Temperature.Value)
	assert.Equal(t, "C", td.Unit)
	assert.Equal(t, mtime, td.Time)
}

func TestTemperatureData_ToArguments(t *testing.T) {
	str := "30.0"
	mtime := time.Now().UTC()

	td := TemperatureData{}
	err := td.ParseFromString(str, mtime)
	require.NoError(t, err)

	arguments := td.ToArguments()
	assert.Equal(t, 4, len(arguments))
}

func TestGeoDataPoint_ParseFromString(t *testing.T) {
	str := "69.3455,45.455"
	mtime := time.Now().UTC()

	gd := GeoDataPoint{}
	err := gd.ParseFromString(str, mtime)
	assert.NoError(t, err)

	assert.Equal(t, 69.3455, gd.Lat.Value)
	assert.Equal(t, 45.455, gd.Lon.Value)
	assert.Equal(t, "coordinate", gd.Unit)
	assert.Equal(t, mtime, gd.Time)
}

func TestGeoDataPoint_ToArguments(t *testing.T) {
	str := "69.3455,45.455"
	mtime := time.Now().UTC()

	gd := GeoDataPoint{}
	err := gd.ParseFromString(str, mtime)
	assert.NoError(t, err)

	arguments := gd.ToArguments()
	assert.Equal(t, 4, len(arguments))
}

func TestAltitudeData_ParseFromString(t *testing.T) {
	str := "32.5"
	mtime := time.Now().UTC()

	ad := AltitudeData{}
	err := ad.ParseFromString(str, mtime)
	assert.NoError(t, err)

	assert.Equal(t, 32.5, ad.Altitude.Value)
	assert.Equal(t, "M", ad.Unit)
	assert.Equal(t, mtime, ad.Time)
}

func TestAltitudeData_negative(t *testing.T) {
	str := "-32.5"
	mtime := time.Now().UTC()

	ad := AltitudeData{}
	err := ad.ParseFromString(str, mtime)
	assert.NoError(t, err)

	assert.Equal(t, 32.5, ad.Altitude.Value)
	assert.Equal(t, "M", ad.Unit)
	assert.Equal(t, mtime, ad.Time)
}

func TestAltitudeData_ToArguments(t *testing.T) {
	str := "32.5"
	mtime := time.Now().UTC()

	ad := AltitudeData{}
	err := ad.ParseFromString(str, mtime)
	assert.NoError(t, err)

	arguments := ad.ToArguments()
	assert.Equal(t, 4, len(arguments))
}

func TestHumidityData_ParseFromString(t *testing.T) {
	str := "90"
	mtime := time.Now().UTC()

	hd := HumidityData{}
	err := hd.ParseFromString(str, mtime)
	assert.NoError(t, err)

	assert.Equal(t, int64(90), hd.Humidity.Value)
	assert.Equal(t, "percentage", hd.Unit)
	assert.Equal(t, mtime, hd.Time)
}

func TestHumidityData_ToArguments(t *testing.T) {
	str := "90"
	mtime := time.Now().UTC()

	hd := HumidityData{}
	err := hd.ParseFromString(str, mtime)
	assert.NoError(t, err)

	arguments := hd.ToArguments()
	assert.Equal(t, 4, len(arguments))
}

func TestPressureData_ParseFromString(t *testing.T) {
	str := "10000"
	mtime := time.Now().UTC()

	pd := PressureData{}
	err := pd.ParseFromString(str, mtime)
	assert.NoError(t, err)

	assert.Equal(t, 10000.0, pd.Pressure.Value)
	assert.Equal(t, "Pa", pd.Unit)
	assert.Equal(t, mtime, pd.Time)
}

func TestPressureData_ToArguments(t *testing.T) {
	str := "10000"
	mtime := time.Now().UTC()

	pd := PressureData{}
	err := pd.ParseFromString(str, mtime)
	assert.NoError(t, err)

	arguments := pd.ToArguments()
	assert.Equal(t, 4, len(arguments))
}
