package main

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

func main() {
	// Connect to the database
	pool, err := connectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer pool.Close()

	// Display the command options
	showCommandOptions()

	// Prompt user to select a command
	command := getUserInput("Enter the command number: ")

	// Perform action based on the selected command
	switch command {
	case "1":
		migrationName := getUserInput("Enter migration name: ")
		if err := createMigrationFile(migrationName); err != nil {
			log.Fatalf("Failed to create migration: %v", err)
		}
	case "2":
		if err := applyMigrations(pool); err != nil {
			log.Fatalf("Failed to apply migrations: %v", err)
		}
	case "3":
		if err := runSeeders(pool); err != nil {
			log.Fatalf("Failed to run seeders: %v", err)
		}
	default:
		fmt.Println("Invalid command")
	}

	fmt.Println("Task completed successfully")
}

func showCommandOptions() {
	fmt.Println("\x1b[35mSelect a command:\x1b[0m")
	fmt.Println("\x1b[32m1. Create Migration\x1b[0m")
	fmt.Println("\x1b[33m2. Apply Migrations\x1b[0m")
	fmt.Println("\x1b[34m3. Run Seeders\x1b[0m")
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func connectDB() (*pgxpool.Pool, error) {
	// Load configuration from environment file
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Retrieve database configuration
	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetInt("DB_PORT")
	dbName := viper.GetString("DB_NAME")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	fmt.Println("Database URL:", dsn)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database URL: %w", err)
	}

	config.MaxConnIdleTime = 10 * time.Minute
	config.MaxConnLifetime = 60 * time.Minute // Set to 1 hour
	config.MaxConns = 5000                    // Adjust based on your environment
	config.MinConns = 100

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
