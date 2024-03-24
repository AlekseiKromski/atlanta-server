package postgres

import (
	"alekseikromski.com/atlanta/modules/storage"
	"fmt"
)

func (p *Postgres) CreateUser(user *storage.User) error {
	query := "INSERT INTO users (username, first_name, second_name, image, email, password) VALUES ($1, $2, $3, $4, $5, $6)"
	if _, err := p.db.Exec(query, user.Username, user.First_name, user.Second_name, user.Image, user.Email, user.Password); err != nil {
		return fmt.Errorf("cannot save datapoint: %v", err)
	}

	return nil
}

func (p *Postgres) GetUserById(id string) (*storage.User, error) {
	rows, err := p.db.Query("SELECT id, username, first_name, second_name, image, email, password, role FROM users WHERE id = $1 LIMIT 1", id)
	if err != nil {
		return nil, fmt.Errorf("cannot send request to check migrations tables: %v", err)
	}
	defer rows.Close()

	user := &storage.User{}
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.First_name, &user.Second_name, &user.Image, &user.Email, &user.Password, &user.Role)
		if err != nil {
			return nil, fmt.Errorf("cannot read response from database: %v", err)
		}
	}

	return user, nil
}

func (p *Postgres) GetUserByUsername(username string) (*storage.User, error) {
	rows, err := p.db.Query("SELECT id, username, first_name, second_name, image, email, password, role FROM users WHERE username = $1 LIMIT 1", username)
	if err != nil {
		return nil, fmt.Errorf("cannot send request to check migrations tables: %v", err)
	}
	defer rows.Close()

	user := &storage.User{}
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.First_name, &user.Second_name, &user.Image, &user.Email, &user.Password, &user.Role)
		if err != nil {
			return nil, fmt.Errorf("cannot read response from database: %v", err)
		}
	}

	return user, nil
}
