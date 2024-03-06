package postgres

import (
	"alekseikromski.com/atlanta/models"
	"alekseikromski.com/atlanta/modules/storage"
	"fmt"
)

func (p *Postgres) SaveDatapoints(deviceUuid string, datapoints []models.DataPoints) error {
	query := "INSERT INTO datapoints (deviceUuid, value, type, unit, measurement_time, flags, label) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	for _, datapoint := range datapoints {
		arguments := []any{deviceUuid}
		arguments = append(arguments, datapoint.ToArguments()...)
		if _, err := p.db.Exec(query, arguments...); err != nil {
			return fmt.Errorf("cannot save datapoint: %v", err)
		}
	}

	return nil
}

func (p *Postgres) GetAllDatapoints() ([]*storage.Datapoint, error) {
	query := "SELECT id, deviceUuid, value, type, unit, label, flags, measurement_time, created_at, updated_at FROM datapoints ORDER BY created_at DESC"
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("cannot send request to check migrations tables: %v", err)
	}
	defer rows.Close()

	dps := []*storage.Datapoint{}
	for rows.Next() {
		dp := &storage.Datapoint{}
		err := rows.Scan(&dp.ID, &dp.DeviceId, &dp.Value, &dp.ValueType, &dp.Unit, &dp.Label, &dp.Flags, &dp.MeasurementTime, &dp.CreatedAt, &dp.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("cannot read response from database: %v", err)
		}
		dps = append(dps, dp)
	}

	return dps, nil
}
