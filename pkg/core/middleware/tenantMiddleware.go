package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

type dbKeyType string

const dbKey dbKeyType = "database"

func SubdomainMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract subdomain from the request URL
		subdomain := extractSubdomain(r.Host)

		// Use the subdomain to determine the tenant name
		tenantName := getTenantName(subdomain)

		// Get the corresponding database for the tenant
		db := getTenantDatabase(tenantName)

		// Set the database context in the request context
		ctx := context.WithValue(r.Context(), dbKey, db)

		// Call the next handler with the updated request context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Function to extract subdomain from host
func extractSubdomain(host string) string {
	parts := strings.Split(host, ".")
	if len(parts) > 2 {
		return parts[0]
	}
	return ""
}

// Function to get tenant name from subdomain
func getTenantName(subdomain string) string {
	// Use the subdomain to lookup the tenant name from a map or lookup table
	fmt.Println(subdomain)
	// Return the tenant name
	return ""
}

// Function to get database for a given tenant name
func getTenantDatabase(tenantName string) *gorm.DB {
	// Use the tenant name to lookup the corresponding database from a map or lookup table
	fmt.Println("db")
	fmt.Println(tenantName)
	// Return the database
	return nil
}
