package postgres

import (
	"alekseikromski.com/atlanta/modules/storage"
	"fmt"
)

func (p *Postgres) GetPermissions() (map[string][]*storage.Endpoint, error) {
	rows, err := p.db.Query("SELECT roles.id, endpoints.urn FROM roles_endpoints INNER JOIN roles ON roles.ID = roles_endpoints.roleuuid INNER JOIN endpoints ON endpoints.ID = roles_endpoints.endpointuuid")
	if err != nil {
		return nil, fmt.Errorf("cannot get roles / endpoint permissions: %v", err)
	}
	defer rows.Close()

	permissions := map[string][]*storage.Endpoint{}
	for rows.Next() {
		role_id := ""
		urn := ""
		err := rows.Scan(&role_id, &urn)
		if err != nil {
			return nil, fmt.Errorf("cannot read response from database: %v", err)
		}

		permissions[role_id] = append(permissions[role_id], &storage.Endpoint{
			Urn: urn,
		})
	}

	return permissions, nil
}
