// File: internal/user/infrastructure/persistence/user_repository_impl.go

package persistence

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/JubaerHossain/restaurant-golang/domain/entity"
	"github.com/JubaerHossain/restaurant-golang/domain/repository"
	"github.com/JubaerHossain/restaurant-golang/pkg/core/app"
	"github.com/JubaerHossain/restaurant-golang/pkg/core/auth"
	"github.com/JubaerHossain/restaurant-golang/pkg/core/cache"
	"github.com/JubaerHossain/restaurant-golang/pkg/core/config"
	utilQuery "github.com/JubaerHossain/restaurant-golang/pkg/query"
)

type UserRepositoryImpl struct {
	app *app.App
}

// NewUserRepository returns a new instance of UserRepositoryImpl
func NewUserRepository(app *app.App) repository.UserRepository {
	return &UserRepositoryImpl{
		app: app,
	}
}

func CacheClear(req *http.Request, cache cache.CacheService) error {
	ctx := req.Context()
	if _, err := cache.ClearPattern(ctx, "get_all_users_*"); err != nil {
		return err
	}
	if _, err := cache.ClearPattern(ctx, "get_payroll_users_*"); err != nil {
		return err
	}
	if _, err := cache.ClearPattern(ctx, "get_all_users__wise_sell_report*"); err != nil {
		return err
	}
	return nil
}

// GetAllUsers returns all users from the database
func (r *UserRepositoryImpl) GetAllUsers(req *http.Request) (*entity.ResponsePagination, error) {
	// Implement logic to get all users
	ctx := req.Context()
	cacheKey := fmt.Sprintf("get_all_users_%s", req.URL.Query().Encode()) // Encode query parameters
	if cachedData, errCache := r.app.Cache.Get(ctx, cacheKey); errCache == nil && cachedData != "" {
		users := &entity.ResponsePagination{}
		if err := json.Unmarshal([]byte(cachedData), users); err != nil {
			return &entity.ResponsePagination{}, err
		}
		return users, nil
	}

	users := []*entity.ResponseUser{}
	query := "SELECT * FROM users" // Example SQL query

	// Get database connection from pool
	conn, err := r.app.DB.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	// Perform the query
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and parse the results
	for rows.Next() {
		var user entity.ResponseUser
		err := rows.Scan(&user.ID, &user.Username) // Example scan, update according to your database schema
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	response := entity.ResponsePagination{
		Data: users,
		// Other pagination details can be set here
	}

	// Cache the response
	jsonData, err := json.Marshal(response)
	if err != nil {
		return &entity.ResponsePagination{}, err
	}
	if err := r.app.Cache.Set(ctx, cacheKey, string(jsonData), time.Duration(config.GlobalConfig.RedisExp)*time.Second); err != nil {
		return &entity.ResponsePagination{}, err
	}
	return &response, nil
}

// GetUserByID returns a user by ID from the database
func (r *UserRepositoryImpl) GetUserByID(userID uint) (*entity.User, error) {
	// Implement logic to get user by ID
	user := &entity.User{}
	if err := r.app.DB.QueryRow(context.Background(), "SELECT id, username, phone FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Username, &user.Phone); err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

// GetUser returns a user by ID from the database
func (r *UserRepositoryImpl) GetUser(userID uint) (*entity.ResponseUser, error) {
	// Implement logic to get user by ID
	resUser := &entity.ResponseUser{}
	query := "SELECT id, username, phone FROM users WHERE id = $1"
	if err := r.app.DB.QueryRow(context.Background(), query, userID).Scan(&resUser.ID, &resUser.Username, &resUser.Phone); err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return resUser, nil
}

func (r *UserRepositoryImpl) GetUserDetails(userID uint) (*entity.ResponseUser, error) {
	// Implement logic to get user details by ID
	resUser := &entity.ResponseUser{}
	err := r.app.DB.QueryRow(context.Background(), `
		SELECT u.id, u.username, u.phone, u.role, u.status
		FROM users u
		WHERE u.id = $1
	`, userID).Scan(&resUser.ID, &resUser.Username, &resUser.Phone, &resUser.Role, &resUser.Status)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return resUser, nil
}

func (r *UserRepositoryImpl) CreateUser(user *entity.ValidateUser, req *http.Request) error {
	// Begin a transaction
	tx, err := r.app.DB.Begin(context.Background())
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			// Recover from panic and rollback the transaction
			tx.Rollback(context.Background())
		} else if err := tx.Commit(context.Background()); err != nil {
			// Commit the transaction if no error occurred, otherwise rollback
			tx.Rollback(context.Background())
		}
	}()

	// Hash the password securely
	hashedPassword, err := utilQuery.HashPassword(user.Password)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}
	user.Password = hashedPassword

	// Create the user within the transaction
	_, err = tx.Exec(context.Background(), `
		INSERT INTO users (username, phone, password) VALUES ($1, $2, $3)
	`, user.Username, user.Phone, user.Password)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	// Clear cache
	if err := CacheClear(req, r.app.Cache); err != nil {
		tx.Rollback(context.Background())
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) UpdateUser(oldUser *entity.User, user *entity.UpdateUser, req *http.Request) (*entity.User, error) {
	tx, err := r.app.DB.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback(context.Background())
		} else if err := tx.Commit(context.Background()); err != nil {
			tx.Rollback(context.Background())
		}
	}()

	query := `
		UPDATE users
		SET username = $1, phone = $2, role = $3, status = $4
		WHERE id = $5
		RETURNING id, username, phone, role, status
	`
	row := tx.QueryRow(context.Background(), query, user.Username, user.Phone, user.Role, user.Status, oldUser.ID)
	updateUser := &entity.User{}
	err = row.Scan(&updateUser.ID, &updateUser.Username, &updateUser.Phone, &updateUser.Role, &updateUser.Status)
	if err != nil {
		tx.Rollback(context.Background())
		return nil, err
	}

	// Clear cache
	if err := CacheClear(req, r.app.Cache); err != nil {
		tx.Rollback(context.Background())
		return nil, err
	}

	return updateUser, nil
}

func (r *UserRepositoryImpl) DeleteUser(user *entity.User, req *http.Request) error {
	tx, err := r.app.DB.Begin(context.Background())
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback(context.Background())
		} else if err := tx.Commit(context.Background()); err != nil {
			tx.Rollback(context.Background())
		}
	}()

	query := "DELETE FROM users WHERE id = $1"
	if _, err := tx.Exec(context.Background(), query, user.ID); err != nil {
		tx.Rollback(context.Background())
		return err
	}

	// Clear cache
	if err := CacheClear(req, r.app.Cache); err != nil {
		tx.Rollback(context.Background())
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) ChangePassword(oldUser *entity.User, user *entity.UserPasswordChange, req *http.Request) error {
	tx, err := r.app.DB.Begin(context.Background())
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback(context.Background())
		} else if err := tx.Commit(context.Background()); err != nil {
			tx.Rollback(context.Background())
		}
	}()

	if err := utilQuery.ComparePassword(oldUser.Password, user.OldPassword); err != nil {
		tx.Rollback(context.Background())
		return err
	}

	newPassword, err := utilQuery.HashPassword(user.NewPassword)
	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	query := "UPDATE users SET password = $1 WHERE id = $2"
	if _, err := tx.Exec(context.Background(), query, newPassword, oldUser.ID); err != nil {
		tx.Rollback(context.Background())
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) TerminateUser(oldUser *entity.User, user *entity.TerminateUser, req *http.Request) error {
	tx, err := r.app.DB.Begin(context.Background())
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback(context.Background())
		} else if err := tx.Commit(context.Background()); err != nil {
			tx.Rollback(context.Background())
		}
	}()

	// Implement logic to terminate user
	query := "DELETE FROM users WHERE id = $1"
	if _, err := tx.Exec(context.Background(), query, oldUser.ID); err != nil {
		tx.Rollback(context.Background())
		return err
	}

	// Clear cache
	if err := CacheClear(req, r.app.Cache); err != nil {
		tx.Rollback(context.Background())
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) Login(loginUser *entity.LoginUser) (*entity.LoginUserResponse, error) {
	user := &entity.User{}
	err := r.app.DB.QueryRow(context.Background(), `
		SELECT id, username, phone, status, password
		FROM users
		WHERE phone = $1
	`, loginUser.Phone).Scan(&user.ID, &user.Username, &user.Phone, &user.Status, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	if err := utilQuery.ComparePassword(user.Password, loginUser.Password); err != nil {
		return nil, fmt.Errorf("invalid password")
	}

	token, err := auth.CreateToken(user)
	if err != nil {
		return nil, err
	}

	responseTokenUser := &entity.LoginUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Phone:    user.Phone,
		Status:   user.Status,
		Token:    token,
	}
	return responseTokenUser, nil
}