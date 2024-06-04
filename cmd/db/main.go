package main

import (
	"bufio"
	"context"
	"fmt"
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
			fmt.Println("\x1b[31mError creating migration\x1b[0m")
			showCommandOptions()
			getUserInput("Enter the command number: ")
		}
	case "2":
		migrationName := getUserInput("Enter migration name: ")
		if err := createMigrationFileWithSeeder(migrationName); err != nil {
			fmt.Println("\x1b[31mError creating migration with seeder\x1b[0m")
			showCommandOptions()
			getUserInput("Enter the command number: ")
		}
	case "3":
		if err := applyMigrations(pool); err != nil {
			fmt.Println("\x1b[31mError applying migrations\x1b[0m")
			showCommandOptions()
			getUserInput("Enter the command number: ")
		}
	case "4":
		if err := runSeeders(pool); err != nil {
			fmt.Println("\x1b[31mError running seeders . Please migrate before \x1b[0m")
			showCommandOptions()
			getUserInput("Enter the command number: ")
		}
	case "0":
		fmt.Println("Exiting...")
		os.Exit(0)

	default:
		fmt.Println("Invalid command")
	}

	showCommandOptions()
	// Prompt user to select a command
	getUserInput("Enter the command number: ")
}

func showCommandOptions() {
	fmt.Println(`
   ___  ____  ____  _______  __
  / _ \/ __ \/ __ \/_  __/ |/_/
 / , _/ /_/ / /_/ / / / _>  <  
/_/|_|\____/\____/ /_/ /_/|_|  
								 
`)
	fmt.Println("\x1b[35mSelect a command:\x1b[0m")
	fmt.Println("\x1b[32m1. Create Migration\x1b[0m")
	fmt.Println("\x1b[37m2. Create Migration with Seeder\x1b[0m")
	fmt.Println("\x1b[33m3. Apply Migrations\x1b[0m")
	fmt.Println("\x1b[34m4. Run Seeders\x1b[0m")
	fmt.Println("\x1b[31m0. Exit\x1b[0m")
}

func getUserInput(prompt string) string {
	// ANSI escape code for green color
	green := "\033[32m"
	// ANSI escape code to reset color
	reset := "\033[0m"

	// Print prompt in green color
	fmt.Print(green + prompt + reset)

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

	timestamp := time.Now().Format("2006_01_02_150405")
	filename := filepath.Join("migrations", fmt.Sprintf("%s_%s.sql", timestamp, name))
	content := fmt.Sprintf("-- Migration %s\n\n", name) +
		fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n", name) +
		"    id SERIAL PRIMARY KEY,\n" +
		"    name VARCHAR(100) NOT NULL,\n" +
		"    description TEXT NOT NULL,\n" +
		"    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,\n" +
		"    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP\n" +
		");\n\n" +
		fmt.Sprintf("CREATE INDEX ON %s (name);\n", name) // Modify column_name with the actual column name

	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to create migration file: %w", err)
	}

	fmt.Printf("Migration file created: %s\n", filename)
	return nil
}

func createSeedFile(tableName string) error {
	fmt.Println("Creating seed file...")
	fmt.Printf("Seed table name: %s\n", tableName)

	timestamp := time.Now().Format("2006_01_02_150405")
	filename := filepath.Join("seeds", fmt.Sprintf("%s_%s_seeder.sql", timestamp, tableName))
	content := fmt.Sprintf("-- Seeder for table %s\n\n", tableName) +
		fmt.Sprintf("INSERT INTO %s (name, description, created_at, updated_at) VALUES\n", tableName) +
		"    ('Value1', 'Description 1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),\n" +
		"    ('Value2', 'Description 2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);\n"

	if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to create seed file: %w", err)
	}

	fmt.Printf("Seed file created: %s\n", filename)
	return nil
}

func createMigrationFileWithSeeder(name string) error {
	if err := createMigrationFile(name); err != nil {
		return err
	}

	tableName := strings.ToLower(name)
	if err := createSeedFile(tableName); err != nil {
		return err
	}

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
		content, err := os.ReadFile(filePath)
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
