package web

import (
	"net/http"

	"github.com/JubaerHossain/rootx/pkg/core/app"
)

// WebRouter registers routes for web endpoints
func WebRouter(application *app.App) http.Handler {
	router := http.NewServeMux()

	// Default route

	return router
}
