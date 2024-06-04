package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/JubaerHossain/rootx/pkg/core/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// Connect to the database
	pool, err := connectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer pool.Close()

	// Run migration creation
	migrationName := "migration_name"
	if err := createMigrationFile(migrationName); err != nil {
		log.Fatalf("Failed to create migration: %v", err)
	}

	// Apply migrations
	if err := applyMigrations(pool); err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	// Run seeders
	if err := runSeeders(pool); err != nil {
		log.Fatalf("Failed to run seeders: %v", err)
	}

	fmt.Println("Migration creation, migration application, and seeding completed successfully")
}

func connectDB() (*pgxpool.Pool, error) {
	cfg := config.GlobalConfig
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database URL: %w", err)
	}

	return pgxpool.NewWithConfig(context.Background(), config)
}

func createMigrationFile(name string) error {
	fmt.Println("Creating migration file...")
	fmt.Printf("Migration name: %s\n", name)

	timestamp := time.Now().Format("20060102150405")
	filename := filepath.Join("migrations", fmt.Sprintf("%s-%s.sql", timestamp, name))
	content := fmt.Sprintf("-- Migration %s\n\n", name) +
		fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n", name) +
		"    id SERIAL PRIMARY KEY,\n" +
		"    name VARCHAR(100) NOT NULL,\n" +
		"    description TEXT NOT NULL,\n" +
		"    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,\n" +
		"    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP\n" +
		");\n"

	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to create migration file: %w", err)
	}

	fmt.Printf("Migration file created: %s\n", filename)
	return nil
}

func applyMigrations(pool *pgxpool.Pool) error {
	fmt.Println("Applying migrations...")
	return executeScriptsInDirectory(pool, "migrations")
}

func runSeeders(pool *pgxpool.Pool) error {
	fmt.Println("Running seeders...")
	return executeScriptsInDirectory(pool, "seeds")
}

func executeScriptsInDirectory(pool *pgxpool.Pool, directory string) error {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filePath := filepath.Join(directory, entry.Name())
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", filePath, err)
		}

		_, err = pool.Exec(context.Background(), string(content))
		if err != nil {
			return fmt.Errorf("failed to execute file %s: %w", filePath, err)
		}
	}

	log.Printf("%s files executed successfully", directory)
	return nil
}
