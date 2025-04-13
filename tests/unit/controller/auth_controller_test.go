package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yashkadam007/brewkar/internal/controller"
	"github.com/yashkadam007/brewkar/internal/domain"

	"time"

	"github.com/google/uuid"
)

// Simple test implementation of AuthService
type TestAuthService struct{}

// Register implements the Register method for AuthService
func (s *TestAuthService) Register(email, password, displayName string) (*domain.User, string, error) {
	// For simplicity, always return a successful result
	// In a real test, we could add logic to return different results based on inputs
	userID := uuid.New()
	now := time.Now()
	user := &domain.User{
		ID:           userID,
		Email:        email,
		PasswordHash: "hashed-password", // Not testing the actual hashing
		DisplayName:  displayName,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	return user, "test-jwt-token", nil
}

// Login implements the Login method for AuthService
func (s *TestAuthService) Login(email, password string) (*domain.User, string, error) {
	// For simplicity, always return a successful result
	userID := uuid.New()
	now := time.Now()
	user := &domain.User{
		ID:           userID,
		Email:        email,
		PasswordHash: "hashed-password", // Not testing the actual hashing
		DisplayName:  "Test User",
		CreatedAt:    now,
		UpdatedAt:    now,
		LastLoginAt:  &now,
	}
	return user, "test-jwt-token", nil
}

// RefreshToken implements the RefreshToken method for AuthService
func (s *TestAuthService) RefreshToken(refreshToken string) (string, string, error) {
	// For simplicity, always return new tokens
	return "new-access-token", "new-refresh-token", nil
}

func setupTest() (*gin.Engine, *controller.AuthController) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	authService := &TestAuthService{}
	authController := controller.NewAuthController(authService)
	return router, authController
}

func TestRegisterEndpoint(t *testing.T) {
	router, authController := setupTest()
	router.POST("/register", authController.Register)

	tests := []struct {
		name           string
		requestBody    map[string]interface{}
		expectedStatus int
	}{
		{
			name: "Valid Registration",
			requestBody: map[string]interface{}{
				"email":       "test@example.com",
				"password":    "password123",
				"displayName": "Test User",
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "Invalid Email",
			requestBody: map[string]interface{}{
				"email":       "invalid-email",
				"password":    "password123",
				"displayName": "Test User",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Missing Password",
			requestBody: map[string]interface{}{
				"email":       "test@example.com",
				"displayName": "Test User",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Short Password",
			requestBody: map[string]interface{}{
				"email":       "test@example.com",
				"password":    "pass",
				"displayName": "Test User",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Missing DisplayName",
			requestBody: map[string]interface{}{
				"email":    "test@example.com",
				"password": "password123",
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestJSON, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(requestJSON))
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, tt.expectedStatus, resp.Code)

			if tt.expectedStatus == http.StatusCreated {
				var response map[string]interface{}
				json.Unmarshal(resp.Body.Bytes(), &response)
				assert.Equal(t, "success", response["status"])

				data, _ := response["data"].(map[string]interface{})
				assert.NotNil(t, data["user"])
				assert.NotNil(t, data["token"])

				user, _ := data["user"].(map[string]interface{})
				assert.Equal(t, tt.requestBody["email"], user["email"])
				assert.Equal(t, tt.requestBody["displayName"], user["displayName"])
			} else {
				var response map[string]interface{}
				json.Unmarshal(resp.Body.Bytes(), &response)
				assert.Equal(t, "error", response["status"])

				errorData, _ := response["error"].(map[string]interface{})
				assert.NotEmpty(t, errorData["code"])
				assert.NotEmpty(t, errorData["message"])
			}
		})
	}
}

func TestLoginEndpoint(t *testing.T) {
	router, authController := setupTest()
	router.POST("/login", authController.Login)

	tests := []struct {
		name           string
		requestBody    map[string]interface{}
		expectedStatus int
	}{
		{
			name: "Valid Login",
			requestBody: map[string]interface{}{
				"email":    "test@example.com",
				"password": "password123",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Invalid Email",
			requestBody: map[string]interface{}{
				"email":    "invalid-email",
				"password": "password123",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Missing Password",
			requestBody: map[string]interface{}{
				"email": "test@example.com",
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestJSON, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestJSON))
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, tt.expectedStatus, resp.Code)

			if tt.expectedStatus == http.StatusOK {
				var response map[string]interface{}
				json.Unmarshal(resp.Body.Bytes(), &response)
				assert.Equal(t, "success", response["status"])

				data, _ := response["data"].(map[string]interface{})
				assert.NotNil(t, data["user"])
				assert.NotNil(t, data["token"])

				user, _ := data["user"].(map[string]interface{})
				assert.Equal(t, tt.requestBody["email"], user["email"])
			} else {
				var response map[string]interface{}
				json.Unmarshal(resp.Body.Bytes(), &response)
				assert.Equal(t, "error", response["status"])

				errorData, _ := response["error"].(map[string]interface{})
				assert.NotEmpty(t, errorData["code"])
				assert.NotEmpty(t, errorData["message"])
			}
		})
	}
}

func TestRefreshEndpoint(t *testing.T) {
	router, authController := setupTest()
	router.POST("/refresh", authController.Refresh)

	tests := []struct {
		name           string
		requestBody    map[string]interface{}
		expectedStatus int
	}{
		{
			name: "Valid Refresh",
			requestBody: map[string]interface{}{
				"refreshToken": "valid-refresh-token",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Missing Refresh Token",
			requestBody:    map[string]interface{}{},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestJSON, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest(http.MethodPost, "/refresh", bytes.NewBuffer(requestJSON))
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, tt.expectedStatus, resp.Code)

			if tt.expectedStatus == http.StatusOK {
				var response map[string]interface{}
				json.Unmarshal(resp.Body.Bytes(), &response)
				assert.Equal(t, "success", response["status"])

				data, _ := response["data"].(map[string]interface{})
				assert.Equal(t, "new-access-token", data["token"])
				assert.Equal(t, "new-refresh-token", data["refreshToken"])
			} else {
				var response map[string]interface{}
				json.Unmarshal(resp.Body.Bytes(), &response)
				assert.Equal(t, "error", response["status"])

				errorData, _ := response["error"].(map[string]interface{})
				assert.NotEmpty(t, errorData["code"])
				assert.NotEmpty(t, errorData["message"])
			}
		})
	}
}
