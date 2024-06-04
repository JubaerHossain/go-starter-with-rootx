package apiHandler

import (
	"fmt"
	"net/http"

	"github.com/JubaerHossain/restaurant-golang/domain/application"
	"github.com/JubaerHossain/restaurant-golang/domain/entity"
	"github.com/JubaerHossain/restaurant-golang/pkg/core/app"
	utilQuery "github.com/JubaerHossain/restaurant-golang/pkg/query"
	"github.com/JubaerHossain/restaurant-golang/pkg/utils"
)

// Handler handles API requests
type Handler struct {
	App *application.App
}

// NewHandler creates a new instance of Handler
func NewHandler(app *app.App) *Handler {
	return &Handler{
		App: application.AppInterface(app),
	}
}

// GetUsers handles requests to fetch users

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	// Implement GetUsers handler
	users, err := h.App.GetUsers(r)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
		// logger.Error("Failed to fetch users", zap.Error(err))
		utils.WriteJSONError(w, http.StatusInternalServerError, "Failed to fetch users")
		return
	}
	// user, err := auth.User(r)
	// if err != nil {
	// 	utils.WriteJSONError(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// fmt.Println("user")
	// fmt.Println(user)

	// Write response
	utils.JsonResponse(w, http.StatusOK, map[string]interface{}{
		"results": users,
	})
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Implement CreateUser handler
	var newUser entity.ValidateUser

	pareErr := utilQuery.BodyParse(&newUser, w, r, true) // Parse request body and validate it
	if pareErr != nil {
		return
	}

	// Call the CreateUser function to create the user
	err := h.App.CreateUser(&newUser, r)
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Write response
	utils.WriteJSONResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "User created successfully",
	})
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := h.App.GetUser(r)
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// Write response
	utils.WriteJSONResponse(w, http.StatusOK, map[string]interface{}{
		"message": "User fetched successfully",
		"results": user,
	})

}

// GetUserDetails  get user details
func (h *Handler) GetUserDetails(w http.ResponseWriter, r *http.Request) {
	user, err := h.App.GetUserDetails(r)
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// Write response
	utils.WriteJSONResponse(w, http.StatusOK, map[string]interface{}{
		"message": "User fetched successfully",
		"results": user,
	})

}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Implement UpdateUser handler
	var updateUser entity.UpdateUser
	pareErr := utilQuery.BodyParse(&updateUser, w, r, true) // Parse request body and validate it
	if pareErr != nil {
		return
	}

	// Call the CreateUser function to create the user
	_, err := h.App.UpdateUser(r, &updateUser)
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Write response
	utils.WriteJSONResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "User updated successfully",
	})
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Implement DeleteUser handler
	err := h.App.DeleteUser(r)
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// Write response
	utils.WriteJSONResponse(w, http.StatusOK, map[string]interface{}{
		"message": "User deleted successfully",
	})
}

func (h *Handler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	// Implement ChangePassword handler
	var updateUser entity.UserPasswordChange
	pareErr := utilQuery.BodyParse(&updateUser, w, r, true) // Parse request body and validate it
	if pareErr != nil {
		return
	}

	// Call the CreateUser function to create the user
	err := h.App.ChangePassword(r, &updateUser)
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Write response
	utils.WriteJSONResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "Password changed successfully",
	})
}
func (h *Handler) TerminateUser(w http.ResponseWriter, r *http.Request) {
	// Implement TerminateUser handler
	var updateUser entity.TerminateUser
	pareErr := utilQuery.BodyParse(&updateUser, w, r, true) // Parse request body and validate it
	if pareErr != nil {
		return
	}
	// Call the CreateUser function to create the user
	err := h.App.TerminateUser(r, &updateUser)
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Write response
	utils.WriteJSONResponse(w, http.StatusCreated, map[string]interface{}{
		"message": "User terminated successfully",
	})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	// Implement Login handler
	var loginUser entity.LoginUser
	pareErr := utilQuery.BodyParse(&loginUser, w, r, true) // Parse request body and validate it
	if pareErr != nil {
		return
	}

	// Call the CreateUser function to create the user
	user, err := h.App.Login(&loginUser)
	if err != nil {
		utils.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Write response
	utils.ReturnResponse(w, http.StatusOK, "Login successful", user)
}
