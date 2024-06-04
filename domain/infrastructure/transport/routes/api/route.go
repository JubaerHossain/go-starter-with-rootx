package api

import (
	"net/http"

	apiHandler "github.com/JubaerHossain/rootx/domain/infrastructure/transport/http/api"
	"github.com/JubaerHossain/rootx/pkg/core/app"
)

// APIRouter registers routes for API endpoints
func APIRouter(application *app.App) http.Handler {
	router := http.NewServeMux()

	// Register user routes
	apiHandler := apiHandler.NewHandler(application)

	// Register user routes
	router.Handle("/users", http.HandlerFunc(apiHandler.GetUsers))

	return router
}
