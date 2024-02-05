package postgres

import (
	"alekseikromski.com/atlanta/models"
	"fmt"
)

func (p *Postgres) SaveDatapoints(datapoints []models.DataPoints) error {
	query := "INSERT INTO datapoints (value, type, unit, measurement_time) VALUES ($1, $2, $3, $4)"
	for _, datapoint := range datapoints {
		if _, err := p.db.Exec(query, datapoint.ToArguments()...); err != nil {
			return fmt.Errorf("cannot save datapoint: %v", err)
		}
	}

	return nil
}
