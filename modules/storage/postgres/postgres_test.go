package postgres

import (
	"alekseikromski.com/atlanta/core"
	"alekseikromski.com/atlanta/models"
	"testing"
	"time"
)

func TestPostgres_SaveDatapoints(t *testing.T) {
	postgres := NewPostgres(
		NewConfig(
		// TODO: import credits
		),
	)

	notifyChannel := make(chan struct{})

	go postgres.Start(notifyChannel, map[string]core.Module{})
	defer postgres.Stop()
	<-notifyChannel

	err := postgres.SaveDatapoints([]models.DataPoints{
		&models.AltitudeData{
			Altitude: &models.FloatData{
				Value: 1,
				Type:  "float",
			},
			Unit: "M",
			Time: time.Now(),
		},
	})

	if err != nil {
		t.Fatalf("error during saving datapoints: %v", err)
		return
	}
}
