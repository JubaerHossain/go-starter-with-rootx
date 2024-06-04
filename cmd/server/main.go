package main

import (
	"fmt"
	"log"
	"net/http"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/JubaerHossain/restaurant-golang/docs"
	"github.com/JubaerHossain/restaurant-golang/domain/infrastructure/transport/routes"
	"github.com/JubaerHossain/restaurant-golang/domain/infrastructure/transport/routes/api"
	"github.com/JubaerHossain/restaurant-golang/pkg/core/app"
)

// @title           Golang Starter API
// @version         1.0
// @description     This is a starter API for Golang projects
// @host            localhost:3021
// @BasePath        /api
func main() {
	// Initialize the application
	application, err := app.StartApp()
	if err != nil {
		log.Fatalf("❌ failed to start application: %v", err)

	}
	// Initialize HTTP server
	httpServer := initHTTPServer(application)
	application.HttpServer = httpServer
	fmt.Printf("🌐 API base URL: http://localhost:%d\n", application.HttpPort) // Added globe emoji
	// Start the server
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("❌ Could not start server: %v\n", err) // Added cross mark emoji
	}
}

func initHTTPServer(application *app.App) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", application.HttpPort),
		Handler: setupRoutes(application),
	}
}

func setupRoutes(application *app.App) http.Handler {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register web routes
	mux.Handle("/", routes.WebRouter(application))

	// Register API routes
	mux.Handle("/api/", http.StripPrefix("/api", api.APIRouter(application)))

	// Register Swagger routes
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	return mux
}
