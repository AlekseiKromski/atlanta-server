package storage

type User struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	First_name  string `json:"first_name"`
	Second_name string `json:"second_name"`
	Password    string `json:"password,omitempty"`
	Image       string `json:"image"`
	Role        string `json:"role"`
}

type StorageUser interface {
	CreateUser(user *User) error
	GetUserById(id string) (*User, error)
	GetUserByUsername(username string) (*User, error)
}
