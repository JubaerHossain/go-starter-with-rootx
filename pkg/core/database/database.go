package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/JubaerHossain/restaurant-golang/pkg/core/config"
	"github.com/JubaerHossain/restaurant-golang/pkg/core/database/seed"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxDatabaseService struct {
	pool *pgxpool.Pool
}

// NewPgxDatabaseService initializes a new database service using pgxpool
func NewPgxDatabaseService() (*PgxDatabaseService, error) {
	cfg := config.GlobalConfig
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}

	config.MaxConnIdleTime = 10 * time.Minute
	config.MaxConnLifetime = 900 * time.Minute
	config.MaxConns = 2000
	config.MinConns = 500

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("failed to create database connection pool: %w", err)
	}

	// Test the database connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for i := 0; i < 3; i++ { // Retry logic
		err := pool.Ping(ctx)
		if err == nil {
			break
		}
		log.Printf("failed to ping database: %v (attempt %d)", err, i+1)
		time.Sleep(2 * time.Second) // Wait before retrying
	}
	if err != nil {
		return nil, fmt.Errorf("failed to ping database after multiple attempts: %w", err)
	}

	log.Println("connected to database")

	return &PgxDatabaseService{pool: pool}, nil
}

func (db *PgxDatabaseService) GetPool() *pgxpool.Pool {
	return db.pool
}

func (db *PgxDatabaseService) Close() {
	db.pool.Close()
	log.Println("database connection closed")
}

func (db *PgxDatabaseService) Migrate() error {
	// Add your migration logic here
	log.Println("database migration completed")
	return nil
}

func (db *PgxDatabaseService) Seed() error {
	if err := seed.NewSeed(db.pool); err != nil {
		return fmt.Errorf("failed to seed database: %w", err)
	}
	log.Println("database seeding completed")
	return nil
}
