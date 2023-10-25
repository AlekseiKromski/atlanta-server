package models

type GeoData struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func (gd *GeoData) Validate() bool {
	return true
}
