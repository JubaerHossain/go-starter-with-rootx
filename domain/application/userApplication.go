// File: internal/user/application/user_service.go

package application

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/JubaerHossain/restaurant-golang/domain/entity"
	"github.com/JubaerHossain/restaurant-golang/domain/infrastructure/persistence"
	"github.com/JubaerHossain/restaurant-golang/domain/repository"
	"github.com/JubaerHossain/restaurant-golang/pkg/core/app"
)

type App struct {
	app  *app.App
	repo repository.UserRepository
}

func AppInterface(app *app.App) *App {
	repo := persistence.NewUserRepository(app)
	return &App{
		app:  app,
		repo: repo,
	}
}

func (c *App) GetUsers(r *http.Request) (*entity.ResponsePagination, error) {
	// Call repository to get all users
	// queryValues := r.URL.Query()
	users, userErr := c.repo.GetAllUsers(r)
	if userErr != nil {
		return nil, userErr
	}
	return users, nil
}

// CreateUser creates a new user
func (c *App) CreateUser(user *entity.ValidateUser, r *http.Request) error {
	err2 := c.repo.CreateUser(user, r)
	if err2 != nil {
		return err2
	}
	return nil
}

// GetUserByID retrieves a user by ID
func (c *App) GetUserByID(r *http.Request) (*entity.User, error) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID")
	}
	user, userErr := c.repo.GetUserByID(uint(id))
	if userErr != nil {
		return nil, userErr
	}
	return user, nil
}
func (c *App) GetUser(r *http.Request) (*entity.ResponseUser, error) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID")
	}
	user, userErr := c.repo.GetUser(uint(id))
	if userErr != nil {
		return nil, userErr
	}
	return user, nil
}

// GetUserDetails retrieves a user by ID
func (c *App) GetUserDetails(r *http.Request) (*entity.ResponseUser, error) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID")
	}
	user, userErr := c.repo.GetUserDetails(uint(id))
	if userErr != nil {
		return nil, userErr
	}
	return user, nil
}

// UpdateUser updates an existing user
func (c *App) UpdateUser(r *http.Request, user *entity.UpdateUser) (*entity.User, error) {
	// Call repository to update user
	oldUser, err := c.GetUserByID(r)
	if err != nil {
		return nil, err
	}

	updateUser, err2 := c.repo.UpdateUser(oldUser, user, r)
	if err2 != nil {
		return nil, err2
	}
	return updateUser, nil
}

// DeleteUser deletes a user by ID
func (c *App) DeleteUser(r *http.Request) error {
	// Call repository to delete user
	user, err := c.GetUserByID(r)
	if err != nil {
		return err
	}

	err2 := c.repo.DeleteUser(user, r)
	if err2 != nil {
		return err2
	}

	return nil
}

// ChangePassword changes the password of a user
func (c *App) ChangePassword(r *http.Request, user *entity.UserPasswordChange) error {
	// Call repository to change password
	oldUser, err := c.GetUserByID(r)
	if err != nil {
		return err
	}
	userErr := c.repo.ChangePassword(oldUser, user, r)
	if userErr != nil {
		return userErr
	}
	return nil
}

// TerminateUser changes the password of a user
func (c *App) TerminateUser(r *http.Request, user *entity.TerminateUser) error {
	// Call repository to change password
	oldUser, err := c.GetUserByID(r)
	if err != nil {
		return err
	}
	userErr := c.repo.TerminateUser(oldUser, user, r)
	if userErr != nil {
		return userErr
	}
	return nil
}

// Login authenticates a user
func (c *App) Login(loginUser *entity.LoginUser) (*entity.LoginUserResponse, error) {
	user, userErr := c.repo.Login(loginUser)
	if userErr != nil {
		return nil, userErr
	}
	return user, nil
}
