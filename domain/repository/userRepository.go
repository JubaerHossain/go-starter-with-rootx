package repository

import (
	"net/http"

	"github.com/JubaerHossain/restaurant-golang/domain/entity"
)

// UserRepository defines methods for user data access
type UserRepository interface {
	GetAllUsers(r *http.Request) (*entity.ResponsePagination, error)
	GetUserByID(userID uint) (*entity.User, error)
	GetUser(userID uint) (*entity.ResponseUser, error)
	GetUserDetails(userID uint) (*entity.ResponseUser, error)
	CreateUser(user *entity.ValidateUser, r *http.Request) (error)
	UpdateUser(oldUser *entity.User, user *entity.UpdateUser, req *http.Request) (*entity.User, error)
	DeleteUser(user *entity.User, req *http.Request) error
	ChangePassword(oldUser *entity.User, user *entity.UserPasswordChange, r *http.Request) error
	TerminateUser(oldUser *entity.User, user *entity.TerminateUser, r *http.Request) error
	Login(loginUser *entity.LoginUser) (*entity.LoginUserResponse, error)
}
