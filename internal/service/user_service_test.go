package service

import (
    "errors"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "github.com/sabhy92/go-mockery-tutorial/internal/repository"
    "github.com/sabhy92/go-mockery-tutorial/internal/mocks"
)

func TestUserService_GetUser(t *testing.T) {
	mockRepo := mocks.NewMockUserRepository(t)

	expectedUser := &repository.User{
		ID: 1,
		Name: "John Doe",
		Email: "john.doe@example.com",
	}

	mockRepo.On("GetUserByID", 1).Return(expectedUser, nil)

	service := NewUserService(mockRepo)

	user, err := service.GetUser(1)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john.doe@example.com", user.Email)

	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUser_InvalidID(t *testing.T) {
	mockRepo := mocks.NewMockUserRepository(t)

	service := NewUserService(mockRepo)

	user, err := service.GetUser(0)

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "invalid user ID", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUser_RepositoryError(t *testing.T) {
	mockRepo := mocks.NewMockUserRepository(t)

	mockRepo.On("GetUserByID", 1).Return(nil, errors.New("database error"))

	service := NewUserService(mockRepo)

	user, err := service.GetUser(1)

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Contains(t, err.Error(), "failed to get user: database error")

	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUser_UserNotFound(t *testing.T) {
	mockRepo := mocks.NewMockUserRepository(t)

	mockRepo.On("GetUserByID", 999).Return(nil, nil)

	service := NewUserService(mockRepo)

	user, err := service.GetUser(999)

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "user not found", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestUserService_RegisterUser_Success(t *testing.T) {
	mockRepo := mocks.NewMockUserRepository(t)

	mockRepo.On("CreateUser", mock.AnythingOfType("*repository.User")).Return(nil)


	service := NewUserService(mockRepo)

	err := service.RegisterUser("Jane Doe", "jane@example.com")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
	
}

func TestUserService_RegisterUser_EmptyName(t *testing.T) {
	mockRepo := mocks.NewMockUserRepository(t)

	service := NewUserService(mockRepo)

	err := service.RegisterUser("", "jane@example.com")

	assert.Error(t, err)
	assert.Equal(t, "name cannot be empty", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestUserService_RegisterUser_EmptyEmail(t *testing.T) {
	mockRepo := mocks.NewMockUserRepository(t)

	service := NewUserService(mockRepo)

	err := service.RegisterUser("Jane Doe", "")

	assert.Error(t, err)
	assert.Equal(t, "email cannot be empty", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestUserService_RemoveUser_Success(t *testing.T) {
    // Arrange
    mockRepo := mocks.NewMockUserRepository(t)
    mockRepo.On("DeleteUser", 1).Return(nil)

    service := NewUserService(mockRepo)

    // Act
    err := service.RemoveUser(1)

    // Assert
    assert.NoError(t, err)
    mockRepo.AssertExpectations(t)
}

func TestUserService_RemoveUser_InvalidID(t *testing.T) {
    // Arrange
    mockRepo := mocks.NewMockUserRepository(t)
    service := NewUserService(mockRepo)

    // Act
    err := service.RemoveUser(-1)

    // Assert
    assert.Error(t, err)
    assert.Equal(t, "invalid user ID", err.Error())
    
    mockRepo.AssertExpectations(t)
}
