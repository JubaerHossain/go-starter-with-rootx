package routes

import (
	"net/http"

	"github.com/JubaerHossain/rootx/pkg/core/app"
	"github.com/JubaerHossain/rootx/pkg/core/health"
	"github.com/JubaerHossain/rootx/pkg/core/middleware"
	"github.com/JubaerHossain/rootx/pkg/core/monitor"
	"github.com/JubaerHossain/rootx/pkg/utils"
)

// WebRouter registers routes for web endpoints
func Router(application *app.App) http.Handler {
	router := http.NewServeMux()

	// Register health check endpoint
	router.Handle("/health", middleware.LoggingMiddleware(http.HandlerFunc(health.HealthCheckHandler())))

	// Register monitoring endpoint
	router.Handle("/metrics", monitor.MetricsHandler())

	// Default route
	router.Handle("/", middleware.LimiterMiddleware(middleware.LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.WriteJSONResponse(w, http.StatusOK, map[string]interface{}{"message": "Welcome to the API"})
	}))))

	// Add Prometheus middleware to monitor all requests
	return middleware.PrometheusMiddleware(router, monitor.RequestsTotal(), monitor.RequestDuration())
}
