package router_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yashkadam007/brewkar/internal/controller"
	"github.com/yashkadam007/brewkar/internal/domain"
	"github.com/yashkadam007/brewkar/internal/router"
)

// Simple AuthService for testing
type TestAuthService struct{}

func (s *TestAuthService) Register(email, password, displayName string) (*domain.User, string, error) {
	return nil, "", nil
}

func (s *TestAuthService) Login(email, password string) (*domain.User, string, error) {
	return nil, "", nil
}

func (s *TestAuthService) RefreshToken(refreshToken string) (string, string, error) {
	return "", "", nil
}

func TestPingEndpoint(t *testing.T) {
	// Setup Gin in test mode
	gin.SetMode(gin.TestMode)

	// Create an auth controller using the test service
	authService := &TestAuthService{}
	authController := controller.NewAuthController(authService)

	// Setup the router with the auth controller
	r := router.SetupRouter(authController)

	// Create a test request to the ping endpoint
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	resp := httptest.NewRecorder()

	// Serve the request to our test router
	r.ServeHTTP(resp, req)

	// Assert status code
	assert.Equal(t, http.StatusOK, resp.Code)

	// Parse the response body
	var responseBody map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &responseBody)

	// Assert no error in parsing JSON
	assert.NoError(t, err)

	// Assert that the response matches the expected format
	assert.Equal(t, "ok", responseBody["status"])
	assert.Equal(t, "pong", responseBody["message"])
}

func TestAPIEndpointPaths(t *testing.T) {
	// Setup Gin in test mode
	gin.SetMode(gin.TestMode)

	// Create an auth controller using the test service
	authService := &TestAuthService{}
	authController := controller.NewAuthController(authService)

	// Setup the router with the auth controller
	r := router.SetupRouter(authController)

	// Table driven tests for API path existence
	tests := []struct {
		name   string
		path   string
		method string
	}{
		{
			name:   "Auth Register Endpoint",
			path:   "/v1/auth/register",
			method: http.MethodPost,
		},
		{
			name:   "Auth Login Endpoint",
			path:   "/v1/auth/login",
			method: http.MethodPost,
		},
		{
			name:   "Auth Refresh Endpoint",
			path:   "/v1/auth/refresh",
			method: http.MethodPost,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// We don't need to send actual data for this test
			// We're just checking if the route is registered
			req, _ := http.NewRequest(tt.method, tt.path, nil)
			resp := httptest.NewRecorder()

			// Serve the request
			r.ServeHTTP(resp, req)

			// If the path doesn't exist, Gin would return 404
			// If the path exists but the method isn't allowed, it would return 405
			// Either way, it should not be 404
			assert.NotEqual(t, http.StatusNotFound, resp.Code, "Route should exist")
		})
	}
}
