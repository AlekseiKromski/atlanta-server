package postgres

import (
	"alekseikromski.com/atlanta/models"
	"fmt"
)

func (p *Postgres) SaveDatapoints(deviceUuid string, datapoints []models.DataPoints) error {
	query := "INSERT INTO datapoints (deviceUuid, value, type, unit, measurement_time) VALUES ($1, $2, $3, $4, $5)"
	for _, datapoint := range datapoints {
		arguments := []any{deviceUuid}
		arguments = append(arguments, datapoint.ToArguments()...)
		if _, err := p.db.Exec(query, arguments...); err != nil {
			return fmt.Errorf("cannot save datapoint: %v", err)
		}
	}

	return nil
}
