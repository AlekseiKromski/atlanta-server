package postgres

import (
	"alekseikromski.com/atlanta/modules/storage"
	"fmt"
	"time"
)

func (p *Postgres) GetDevices() ([]*storage.Device, error) {
	query := "SELECT id, description, status, created_at, updated_at, deleted_at FROM devices WHERE deleted_at IS NULL ORDER BY created_at DESC"

	rows, err := p.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("cannot get all unique labels: %v", err)
	}
	defer rows.Close()

	devices := []*storage.Device{}
	for rows.Next() {
		device := &storage.Device{}
		err := rows.Scan(&device.Id, &device.Description, &device.Status, &device.CreatedAt, &device.UpdatedAt, &device.DeletedAt)
		if err != nil {
			return nil, fmt.Errorf("cannot read response from database: %v", err)
		}
		devices = append(devices, device)
	}

	return devices, nil
}

func (p *Postgres) CreateDevice(description string, status bool) error {
	query := "INSERT INTO devices (description, status) VALUES ($1, $2)"

	_, err := p.db.Exec(query, description, status)
	if err != nil {
		return fmt.Errorf("cannot create datapoint: %v", err)
	}

	return nil
}

func (p *Postgres) UpdateDevice(id, description string, status bool) error {
	query := "UPDATE devices SET description = $1, status = $2, updated_at = $3 WHERE id = $4 AND deleted_at IS NULL"

	now := time.Now().UTC().Format(time.RFC3339)
	_, err := p.db.Exec(query, description, status, now, id)
	if err != nil {
		return fmt.Errorf("cannot update datapoint: %v", err)
	}

	return nil
}

func (p *Postgres) DeleteDevice(id string) error {
	query := "UPDATE devices SET deleted_at = $1, updated_at = $2 WHERE id = $3"

	now := time.Now().UTC().Format(time.RFC3339)
	_, err := p.db.Exec(query, now, now, id)
	if err != nil {
		return fmt.Errorf("cannot update datapoint: %v", err)
	}

	return nil
}
