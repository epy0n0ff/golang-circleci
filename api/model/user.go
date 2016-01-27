package model

// User is ...
type User struct {
	ID   int `json:"id"`
	Name string `json:"name"`
}

// NewUser is ...
func NewUser(id int, name string) User {
	return User{id, name}
}
