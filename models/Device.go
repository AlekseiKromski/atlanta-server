package models

type Device struct {
	Uuid    string   `json:"uuid"`
	ApiKey  string   `json:"api_key"`
	Battery *float64 `json:"battery"` //Percentage
}

func (d *Device) Validate() bool {
	return d.Battery != nil
}
