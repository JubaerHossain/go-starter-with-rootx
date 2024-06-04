package entity

import (
	"time"

	"github.com/JubaerHossain/rootx/pkg/core/entity"
)

// User represents the user entity
type User struct {
	ID        uint          `json:"id" gorm:"primaryKey;autoIncrement;not null"` // Primary key
	Username  string        `json:"username" validate:"required,min=3,max=50" gorm:"index"`
	Phone     string        `json:"phone" validate:"required,min=11,max=15" gorm:"index;unique"`
	Password  string        `json:"password" validate:"required,min=6,max=20" gorm:"index;not null"`
	Role      entity.Role   `json:"role" gorm:"index;default:chef" validate:"required,oneof=admin manager waiter chef"`
	CreatedAt time.Time     `json:"created_at" gorm:"index;autoCreateTime"`
	UpdatedAt time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	Status    entity.Status `json:"status" gorm:"index;default:pending" validate:"required,oneof=active inactive deleted pending"`
}

type ValidateUser struct {
	ID        uint          `json:"id" gorm:"primaryKey;autoIncrement;not null"` // Primary key
	Username  string        `json:"username" validate:"required,min=3,max=50" gorm:"index"`
	Phone     string        `json:"phone" validate:"required,min=11,max=15" gorm:"index;unique"`
	Password  string        `json:"password" validate:"required,min=6,max=20" gorm:"index;not null"`
	Role      entity.Role   `json:"role" gorm:"index;default:chef" validate:"required,oneof=admin manager waiter chef"`
	CreatedAt time.Time     `json:"created_at" gorm:"index;autoCreateTime"`
	UpdatedAt time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	Status    entity.Status `json:"status" gorm:"index;default:pending" validate:"required,oneof=active inactive deleted pending"`
}

// updateUser represents the user update request
type UpdateUser struct {
	Username  string        `json:"username" validate:"omitempty,min=3,max=50"`
	Phone     string        `json:"phone" validate:"omitempty,min=11,max=15"`
	UpdatedAt time.Time     `json:"updated_at" gorm:"autoUpdateTime" validate:"omitempty"`
	Role      entity.Role   `json:"role" validate:"omitempty,oneof=admin manager waiter chef"`
	Status    entity.Status `json:"status" gorm:"default:pending" validate:"omitempty,oneof=active inactive deleted pending"`
}

// responseUser represents the user response
type ResponseUser struct {
	ID        uint          `json:"id"`
	Username  string        `json:"username"`
	Phone     string        `json:"phone"`
	Role      entity.Role   `json:"role"`
	CreatedAt time.Time     `json:"created_at"`
	Status    entity.Status `json:"status"`
}

type TerminateUser struct {
	ID     uint          `json:"id"`
	Status entity.Status `json:"status"`
}

type UserPasswordChange struct {
	ID          uint   `json:"id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ResponsePagination struct {
	Data       []*ResponseUser   `json:"data"`
	Pagination entity.Pagination `json:"pagination"`
}
