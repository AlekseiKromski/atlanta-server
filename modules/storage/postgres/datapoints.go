package postgres

import (
	"alekseikromski.com/atlanta/models"
	"alekseikromski.com/atlanta/modules/storage"
	"fmt"
	"time"
)

func (p *Postgres) SaveDatapoints(deviceUuid string, datapoints []models.DataPoints) ([]*storage.Datapoint, error) {
	query := "INSERT INTO datapoints (deviceUuid, value, type, unit, measurement_time, flags, label) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, deviceUuid, value, type, unit, label, flags, measurement_time, created_at, updated_at"

	dps := []*storage.Datapoint{}
	for _, datapoint := range datapoints {
		arguments := []any{deviceUuid}
		arguments = append(arguments, datapoint.ToArguments()...)
		rows, err := p.db.Query(query, arguments...)
		if err != nil {
			return nil, fmt.Errorf("cannot save datapoint: %v", err)
		}

		for rows.Next() {
			dp := &storage.Datapoint{}
			err := rows.Scan(&dp.ID, &dp.DeviceId, &dp.Value, &dp.ValueType, &dp.Unit, &dp.Label, &dp.Flags, &dp.MeasurementTime, &dp.CreatedAt, &dp.UpdatedAt)
			if err != nil {
				return nil, fmt.Errorf("cannot read response from database: %v", err)
			}
			dps = append(dps, dp)
		}
	}

	return dps, nil
}

func (p *Postgres) GetAllDatapoints() ([]*storage.Datapoint, error) {
	query := "SELECT id, deviceUuid, value, type, unit, label, flags, measurement_time, created_at, updated_at FROM datapoints ORDER BY created_at DESC LIMIT 3000"
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("cannot get all datapoints: %v", err)
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

func (p *Postgres) FindDatapoints(fd *storage.FindDatapointsRequest) ([]*storage.Datapoint, []string, error) {
	query := "SELECT DISTINCT ON (measurement_time, label) id, deviceUuid, value, type, unit, label, flags, measurement_time, created_at, updated_at FROM datapoints WHERE measurement_time BETWEEN $1 and $2 "

	items := ""
	for index, sel := range fd.Select {
		if index == len(fd.Select)-1 {
			items += fmt.Sprintf("'%s'", sel)
			continue
		}
		items += fmt.Sprintf("'%s',", sel)
	}

	query += fmt.Sprintf("AND label IN (%s) ", items)

	if !fd.Ignored {
		query += "AND flags != 'ignored' "
	}

	query += fmt.Sprintf("AND deviceUuid = '%s'", fd.Device)

	rows, err := p.db.Query(query, fd.Start.Format(time.RFC3339), fd.End.Format(time.RFC3339))
	if err != nil {
		return nil, nil, fmt.Errorf("cannot find datapoints: %v", err)
	}
	defer rows.Close()

	dps := []*storage.Datapoint{}
	for rows.Next() {
		dp := &storage.Datapoint{}
		err := rows.Scan(&dp.ID, &dp.DeviceId, &dp.Value, &dp.ValueType, &dp.Unit, &dp.Label, &dp.Flags, &dp.MeasurementTime, &dp.CreatedAt, &dp.UpdatedAt)
		if err != nil {
			return nil, nil, fmt.Errorf("cannot read response from database: %v", err)
		}
		dps = append(dps, dp)
	}

	query = "SELECT measurement_time FROM datapoints WHERE measurement_time BETWEEN $1 and $2 "

	items = ""
	for index, sel := range fd.Select {
		if index == len(fd.Select)-1 {
			items += fmt.Sprintf("'%s'", sel)
			continue
		}
		items += fmt.Sprintf("'%s',", sel)
	}

	query += fmt.Sprintf("AND label IN (%s) ", items)

	query += "GROUP BY measurement_time ORDER by measurement_time ASC"

	rows, err = p.db.Query(query, fd.Start.Format(time.RFC3339), fd.End.Format(time.RFC3339))
	if err != nil {
		return nil, nil, fmt.Errorf("cannot get all unique measurement time: %v", err)
	}
	defer rows.Close()

	labels := []string{}
	for rows.Next() {
		time := ""
		err := rows.Scan(&time)
		if err != nil {
			return nil, nil, fmt.Errorf("cannot read response from database: %v", err)
		}
		labels = append(labels, time)
	}

	return dps, labels, nil
}

func (p *Postgres) FindAllLabels() ([]string, error) {
	query := "SELECT DISTINCT label FROM datapoints WHERE label IS NOT NULL"

	rows, err := p.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("cannot get all unique labels: %v", err)
	}
	defer rows.Close()

	labels := []string{}
	for rows.Next() {
		time := ""
		err := rows.Scan(&time)
		if err != nil {
			return nil, fmt.Errorf("cannot read response from database: %v", err)
		}
		labels = append(labels, time)
	}

	return labels, nil
}

func (p *Postgres) FindAllDevices() ([]*storage.Device, error) {
	query := "SELECT id, description FROM devices ORDER BY created_at DESC"

	rows, err := p.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("cannot get all unique labels: %v", err)
	}
	defer rows.Close()

	devices := []*storage.Device{}
	for rows.Next() {
		device := &storage.Device{}
		err := rows.Scan(&device.Id, &device.Description)
		if err != nil {
			return nil, fmt.Errorf("cannot read response from database: %v", err)
		}
		devices = append(devices, device)
	}

	return devices, nil
}
