package auth

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	userEntity "github.com/JubaerHossain/rootx/domain/entity"
	"github.com/JubaerHossain/rootx/pkg/core/config"
	"github.com/JubaerHossain/rootx/pkg/core/entity"
	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(payload interface{}) (string, error) {
	// Create the JWT claims
	jwtTime := config.GlobalConfig.JwtExpiration
	if jwtTime == "" {
		jwtTime = "24"
	}
	expiration, err := strconv.Atoi(jwtTime)
	if err != nil {
		expiration = 24
	}

	claims := jwt.MapClaims{
		"user": payload,
		"exp":  time.Now().Add(time.Hour * time.Duration(expiration)).Unix(), // Token expiration time
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	secretKey := config.GlobalConfig.JwtSecretKey
	if secretKey == "" {
		secretKey = "your-secret"
	}
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT token: %v", err)
	}

	return tokenString, nil
}

// VerifyToken verifies the JWT token
func VerifyToken(tokenString string) (bool, *userEntity.AuthUser, error) {
	// Remove the "Bearer " prefix from the token string
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key
		secretKey := config.GlobalConfig.JwtSecretKey
		if secretKey == "" {
			secretKey = "your-secret"
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return false, nil, fmt.Errorf("failed to parse token: %v", err)
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extract user information from claims
		if userClaim, exists := claims["user"]; exists {
			// Extract user information from claims
			if userData, ok := userClaim.(map[string]interface{}); ok {
				// Deserialize user data from the map
				user := userEntity.AuthUser{
					ID:       uint(int(userData["id"].(float64))),
					Name: userData["name"].(string),
					Phone:    userData["phone"].(string),
					Role:     entity.Role(userData["role"].(string)),
					Status:   entity.Status(userData["status"].(string)),
				}
				return true, &user, nil
			}
		}
	}

	return false, nil, nil
}

func User(r *http.Request) (*userEntity.AuthUser, error) {
	user, ok := r.Context().Value(entity.AuthUser).(*userEntity.AuthUser)
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}
