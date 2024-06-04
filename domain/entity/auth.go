package entity

import "github.com/JubaerHossain/rootx/pkg/core/entity"

type LoginUser struct {
	Phone    string `json:"phone" validate:"required,min=11,max=15"`
	Password string `json:"password" validate:"required,min=6,max=20"`
}

type LoginUserResponse struct {
	ID       uint          `json:"id"`
	Username string        `json:"username"`
	Phone    string        `json:"phone"`
	Status   entity.Status `json:"status"`
	Token    string        `json:"token"`
}
type AuthUser struct {
	ID       uint          `json:"id"`
	Username string        `json:"username"`
	Phone    string        `json:"phone"`
	Role     entity.Role   `json:"role"`
	Status   entity.Status `json:"status"`
}
