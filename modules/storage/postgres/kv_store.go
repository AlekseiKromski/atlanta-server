package postgres

import (
	"alekseikromski.com/atlanta/modules/storage"
	"fmt"
	"time"
)

func (p *Postgres) GetValues(key, user_id string) ([]*storage.KVStore, error) {
	rows, err := p.db.Query("SELECT id, key, value, useruuid, created_at, updated_at, deleted_at FROM store WHERE useruuid = $1 AND deleted_at IS NULL AND key = $2", user_id, key)
	if err != nil {
		return nil, fmt.Errorf("cannot get value: %v", err)
	}
	defer rows.Close()

	records := []*storage.KVStore{}
	for rows.Next() {
		record := &storage.KVStore{}
		err := rows.Scan(&record.Id, &record.Key, &record.Value, &record.UserId, &record.CreateAt, &record.UpdatedAt, &record.DeletedAt)
		if err != nil {
			return nil, fmt.Errorf("cannot read response from database: %v", err)
		}

		records = append(records, record)
	}

	return records, nil
}

func (p *Postgres) UpsertValue(key, value, user_id string) error {
	records, err := p.GetValues(key, user_id)
	if err != nil {
		return fmt.Errorf("cannot get value: %v", err)
	}

	if len(records) == 0 {
		// Create
		query := "INSERT INTO store (key, value, useruuid) VALUES ($1, $2, $3)"
		if _, err := p.db.Exec(query, key, value, user_id); err != nil {
			return fmt.Errorf("cannot save kv record: %v", err)
		}

		return nil
	}

	// Update
	query := "UPDATE store SET value = $1, updated_at = $2 WHERE useruuid = $3 AND key = $4"
	if _, err := p.db.Exec(query, value, time.Now().UTC().Format(time.RFC3339), user_id, key); err != nil {
		return fmt.Errorf("cannot update kv record: %v", err)
	}

	return nil
}
