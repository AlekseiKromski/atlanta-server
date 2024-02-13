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
			"localhost",
			"atlanta",
			"postgres",
			"postgres",
			5432,
		),
	)

	notifyChannel := make(chan struct{})

	go postgres.Start(notifyChannel, map[string]core.Module{})
	defer postgres.Stop()
	<-notifyChannel

	err := postgres.SaveDatapoints("3cc76ff4-cbaa-436c-b727-45d526facfc7", []models.DataPoints{
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

func TestPostgres_migrations(t *testing.T) {
	postgres := NewPostgres(
		NewConfig(
			"localhost",
			"atlanta",
			"postgres",
			"postgres",
			5432,
		),
	)

	notifyChannel := make(chan struct{})

	go postgres.Start(notifyChannel, map[string]core.Module{})
	defer postgres.Stop()
	<-notifyChannel

	if err := postgres.migrations(); err != nil {
		t.Fatalf("cannot make migrations: %v", err)
	}
}
