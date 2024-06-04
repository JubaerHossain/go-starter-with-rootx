package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/JubaerHossain/rootx/docs"
	"github.com/JubaerHossain/rootx/domain/infrastructure/transport/routes/api"
	"github.com/JubaerHossain/rootx/domain/infrastructure/transport/routes/web"
	"github.com/JubaerHossain/rootx/pkg/core/app"
	"github.com/JubaerHossain/rootx/pkg/core/health"
	"github.com/JubaerHossain/rootx/pkg/core/middleware"
	"github.com/JubaerHossain/rootx/pkg/core/monitor"
	"github.com/JubaerHossain/rootx/pkg/utils"
	httpSwagger "github.com/swaggo/http-swagger"
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
	mux.Handle("/web", web.WebRouter(application))

	// Register API routes
	mux.Handle("/api/", http.StripPrefix("/api", api.APIRouter(application)))

	// Register Swagger routes
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	// Register health check endpoint
	mux.Handle("/health", middleware.LoggingMiddleware(http.HandlerFunc(health.HealthCheckHandler())))

	// Register monitoring endpoint
	mux.Handle("/metrics", monitor.MetricsHandler())

	// Add Prometheus middleware to monitor all requests

	// Default route
	mux.Handle("/", middleware.LimiterMiddleware(middleware.LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.WriteJSONResponse(w, http.StatusOK, map[string]interface{}{"message": "Welcome to the API"})
	}))))

	return middleware.PrometheusMiddleware(mux, monitor.RequestsTotal(), monitor.RequestDuration())
}
