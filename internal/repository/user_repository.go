package repository

type User struct {
	ID        int `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`	
}

type UserRepository interface {
	GetUserByID(id int) (*User, error)
	CreateUser(user *User) error
	DeleteUser(id int) error
}
