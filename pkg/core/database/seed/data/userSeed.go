package dataSeed

import (
	"fmt"
	"time"

	userEntity "github.com/JubaerHossain/rootx/domain/entity"
	"github.com/JubaerHossain/rootx/pkg/core/entity"
	"github.com/JubaerHossain/rootx/pkg/core/logger"
	utilQuery "github.com/JubaerHossain/rootx/pkg/query"
	"gorm.io/gorm"
)

var roles = []entity.Role{entity.AdminRole, entity.ManagerRole, entity.UserRole}

// SeedUsers generates and inserts dummy user data into the database.
func SeedUsers(db *gorm.DB, numUsers int) error {
	// Begin transaction
	tx := db.Begin()

	// Defer a function to handle transaction rollback or commit
	defer func() {
		if r := recover(); r != nil {
			// Recover from panic and rollback the transaction
			tx.Rollback()
		} else {
			// Commit the transaction if no error occurred
			tx.Commit()
		}
	}()

	// Delete existing users and update logs
	if err := tx.Exec("DELETE FROM users").Error; err != nil {
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
		// Create user in the database
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
	}

	logger.Info("Seeded users successfully")
	return nil
}
