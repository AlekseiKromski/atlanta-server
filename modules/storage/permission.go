package storage

type Role struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Endpoint struct {
	Urn string `json:"urn"`
}

type StoragePermissions interface {
	GetPermissions() (map[string][]*Endpoint, error)
}
