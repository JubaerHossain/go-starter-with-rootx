package dataSeed

import (
	"context"
	"fmt"
	"time"

	userEntity "github.com/JubaerHossain/rootx/domain/entity"
	"github.com/JubaerHossain/rootx/pkg/core/entity"
	"github.com/JubaerHossain/rootx/pkg/core/logger"
	utilQuery "github.com/JubaerHossain/rootx/pkg/query"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var roles = []entity.Role{entity.AdminRole, entity.ManagerRole, entity.UserRole}

// SeedUsers generates and inserts dummy user data into the database.
func SeedUsers(pool *pgxpool.Pool, numUsers int) error {
	// Begin transaction
	tx, err := pool.Begin(context.Background())
	if err != nil {
		return err
	}

	// Defer a function to handle transaction rollback or commit
	defer func() {
		if r := recover(); r != nil {
			// Recover from panic and rollback the transaction
			tx.Rollback(context.Background())
		} else {
			// Commit the transaction if no error occurred
			if err := tx.Commit(context.Background()); err != nil {
				logger.Error("Error committing transaction:", zap.Error(err))
			}
		}
	}()

	// Delete existing users and update logs
	_, err = tx.Exec(context.Background(), "DELETE FROM users")
	if err != nil {
		return err
	}

	// Hash the default password
	password, err := utilQuery.HashPassword("password")
	if err != nil {
		return err
	}

	// Loop through roles and create users
	for i, role := range roles {
		var user userEntity.User
		user.Name = string(role)
		user.Phone = fmt.Sprintf("0170000000%d", i+1)
		user.Password = password
		user.Role = role
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()
		user.Status = entity.Active

		fmt.Println(user)
		// Construct the insert query with placeholders
		query := `INSERT INTO users (name, phone, password, role, created_at, updated_at, status) VALUES ($1, $2, $3, $4, $5, $6, $7)`

		// Execute the query with user data as arguments
		_, err := tx.Exec(context.Background(), query, user.Name, user.Phone, user.Password, user.Role, user.CreatedAt, user.UpdatedAt, user.Status)
		if err != nil {
			return err
		}
	}

	logger.Info("Seeded users successfully")
	return nil
}
