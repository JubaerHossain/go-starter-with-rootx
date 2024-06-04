package seed

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewSeed seeds the database
func NewSeed(pool *pgxpool.Pool) error {
	ctx := context.Background()

	// Example of seeding logic
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	// Assuming you have some SQL statements for seeding
	seedStatements := []string{
		`INSERT INTO users (id, name) VALUES (1, 'John Doe');`,
		// Add more seed statements here
	}

	for _, stmt := range seedStatements {
		if _, err := conn.Exec(ctx, stmt); err != nil {
			return err
		}
	}

	log.Println("database seeding completed")
	return nil
}
