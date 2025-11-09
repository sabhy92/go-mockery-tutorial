package main

import (
	"fmt"
	"log"
	"github.com/sabhy92/go-mockery-tutorial/internal/service"
	"github.com/sabhy92/go-mockery-tutorial/internal/repository"
)

type SimpleUserRepository struct {
	users map[int]*repository.User
	nextID int
}

func NewSimpleUserRepository() *SimpleUserRepository {
	return &SimpleUserRepository{
		users: make(map[int]*repository.User),
		nextID: 1,
	}
}
func (r *SimpleUserRepository) GetUserByID(id int) (*repository.User, error) {
	if user, ok := r.users[id]; ok {
		return user, nil
	}
	return nil, fmt.Errorf("user not found")
}

func (r *SimpleUserRepository) CreateUser(user *repository.User) error {
	user.ID = r.nextID
	r.users[r.nextID] = user
	r.nextID++
	return nil
}

func (r *SimpleUserRepository) DeleteUser(id int) error {
	delete(r.users, id)
	return nil
}

func main() {
	repo := NewSimpleUserRepository()
	service := service.NewUserService(repo)

	err := service.RegisterUser("Alice Smith", "alice@example.com")
	if err != nil {
		log.Fatalf("Failed to register user: %v", err)
	}
	fmt.Println("User registered successfully")

	user, err := service.GetUser(1)
	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}
	fmt.Println("User:", user)

	err = service.RemoveUser(1)
	if err != nil {
		log.Fatalf("Failed to remove user: %v", err)
	}
	fmt.Println("User removed successfully")
}