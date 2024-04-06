package storage

type KVStore struct {
	Id        string       `json:"id"`
	Key       string       `json:"key"`
	Value     string       `json:"value"`
	UserId    string       `json:"user_id"`
	CreateAt  string       `json:"created_at"`
	UpdatedAt string       `json:"updated_at"`
	DeletedAt *string      `json:"deleted_at"`
	Endpoints *[]*Endpoint `json:"endpoints,omitempty"`
}

type StorageKvStore interface {
	GetValues(key, user_id string) ([]*KVStore, error)
	UpsertValue(key, value, user_id string) error
}
