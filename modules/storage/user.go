package storage

type User struct {
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	First_name  string  `json:"first_name"`
	Second_name string  `json:"second_name"`
	Password    *string `json:"password,omitempty"`
	Image       string  `json:"image"`
}

type StorageUser interface {
	CreateUser(user *User) error

	GetUser(id string) (*User, error)
}
