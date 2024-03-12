package storage

import "time"

type FindDatapointsRequest struct {
	Start   *time.Time `json:"start"`
	End     *time.Time `json:"end"`
	Select  []string   `json:"select"`
	Ignored bool       `json:"ignored"`
}

func (fd *FindDatapointsRequest) Validate() bool {
	if fd.Start == nil {
		return false
	}
	if fd.End == nil {
		return false
	}
	if fd.Select == nil {
		return false
	}

	return true
}
