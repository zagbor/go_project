package services

import (
	"errors"
	"go-microservice/models"
	"sync"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

// UserService handles business logic for users.
type UserService struct {
	users  map[int]models.User
	mu     sync.RWMutex
	nextID int
}

// NewUserService creates a new instance of UserService.
func NewUserService() *UserService {
	return &UserService{
		users:  make(map[int]models.User),
		nextID: 1,
	}
}

// Create adds a new user to the storage.
func (s *UserService) Create(user models.User) models.User {
	s.mu.Lock()
	defer s.mu.Unlock()

	user.ID = s.nextID
	s.nextID++
	s.users[user.ID] = user
	return user
}

// Get retrieves a user by ID.
func (s *UserService) Get(id int) (models.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.users[id]
	if !exists {
		return models.User{}, ErrUserNotFound
	}
	return user, nil
}

// GetAll retrieves all users.
func (s *UserService) GetAll() []models.User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	users := make([]models.User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}

// Update modifies an existing user.
func (s *UserService) Update(id int, updatedUser models.User) (models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.users[id]
	if !exists {
		return models.User{}, ErrUserNotFound
	}

	updatedUser.ID = id // Ensure ID remains consistent
	s.users[id] = updatedUser
	return updatedUser, nil
}

// Delete removes a user by ID.
func (s *UserService) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; !exists {
		return ErrUserNotFound
	}

	delete(s.users, id)
	return nil
}
