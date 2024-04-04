package postgres

import (
	"alekseikromski.com/atlanta/modules/storage"
	"fmt"
	"time"
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

func (p *Postgres) CreatePermission(roleId, endpointId string) error {
	query := "INSERT INTO roles_endpoints (roleuuid, endpointuuid) VALUES ($1, $2)"
	if _, err := p.db.Exec(query, roleId, endpointId); err != nil {
		return fmt.Errorf("cannot save permission: %v", err)
	}

	return nil
}

func (p *Postgres) GetEndpointIdsByRoleId(roleId string) ([]string, error) {
	rows, err := p.db.Query("SELECT endpointuuid FROM roles_endpoints WHERE roleuuid = $1 AND deleted_at IS NULL", roleId)
	if err != nil {
		return nil, fmt.Errorf("cannot get endpoint ids: %v", err)
	}
	defer rows.Close()

	endpointids := []string{}
	for rows.Next() {
		endpointId := ""
		err := rows.Scan(&endpointId)
		if err != nil {
			return nil, fmt.Errorf("cannot read response from database: %v", err)
		}

		endpointids = append(endpointids, endpointId)
	}

	return endpointids, nil
}

func (p *Postgres) DeletePermission(roleId, endpointId string) error {
	query := "UPDATE roles_endpoints SET deleted_at = $1, updated_at = $2 WHERE endpointuuid = $3 AND roleuuid = $4"
	now := time.Now().UTC().Format(time.RFC3339)
	if _, err := p.db.Exec(query, now, now, endpointId, roleId); err != nil {
		return fmt.Errorf("cannot delete permission: %v", err)
	}

	return nil
}
