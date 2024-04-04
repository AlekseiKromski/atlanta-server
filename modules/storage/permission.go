package storage

type StoragePermissions interface {
	GetPermissions() (map[string][]*Endpoint, error)
	CreatePermission(roleId, endpointId string) error
	GetEndpointIdsByRoleId(roleId string) ([]string, error)
	DeletePermission(roleId, endpointId string) error
}
