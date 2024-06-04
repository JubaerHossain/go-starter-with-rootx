package middleware

import (
	"context"
	"net/http"

	"github.com/JubaerHossain/restaurant-golang/pkg/core/auth"
	"github.com/JubaerHossain/restaurant-golang/pkg/core/entity"
	"github.com/JubaerHossain/restaurant-golang/pkg/utils"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request is authenticated
		token := r.Header.Get("Authorization")
		if token == "" {
			utils.WriteJSONError(w, http.StatusUnauthorized, "Unauthorized: missing token")
			return
		}

		// Verify the token
		isValid, user, err := auth.VerifyToken(token)
		if err != nil {
			utils.WriteJSONError(w, http.StatusUnauthorized, "Unauthorized: "+err.Error())
			return
		}
		if !isValid {
			utils.WriteJSONError(w, http.StatusUnauthorized, "Unauthorized: invalid token")
			return
		}

		// Add the authenticated user to the request context
		ctx := r.Context()
		ctx = context.WithValue(ctx, entity.AuthUser, user)
		r = r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
