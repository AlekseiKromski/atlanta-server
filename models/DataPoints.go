package models

import "time"

type DataPoints interface {
	ParseFromString(val string, measurementTime time.Time) error
}
