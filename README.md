Simple CRUD app for studying how an interface is used for Database operations.
This project also demonstrates how to use mockery for generating mocks in Go.

## UserRepository Interface

The `UserRepository` interface defines the contract for user data operations. It is located in `internal/repository/user_repository.go`.

```go
type UserRepository interface {
    GetUserByID(id int) (*User, error)
    CreateUser(user *User) error
    DeleteUser(id int) error
}
```

### Methods

#### `GetUserByID(id int) (*User, error)`
Retrieves a user by their ID.
- **Parameters**: `id` - The unique identifier of the user
- **Returns**: A pointer to the `User` struct if found, or an error if the user doesn't exist or an error occurs

#### `CreateUser(user *User) error`
Creates a new user in the repository.
- **Parameters**: `user` - A pointer to a `User` struct containing the user's information
- **Returns**: An error if the creation fails, or `nil` on success

#### `DeleteUser(id int) error`
Deletes a user from the repository by their ID.
- **Parameters**: `id` - The unique identifier of the user to delete
- **Returns**: An error if the deletion fails, or `nil` on success

### Implementation

The interface is implemented by `SimpleUserRepository` in `internal/main.go`, which provides an in-memory implementation for demonstration purposes.

### Mock Generation

Mocks for this interface are generated using mockery and stored in `internal/mocks/mock_UserRepository.go`.
