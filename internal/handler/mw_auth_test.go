package handler

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/patyukin/mbs-api-gateway/internal/handler/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestAuthMiddleware тестирует middleware Auth
func TestAuthMiddleware(t *testing.T) {
	mockUseCase := &mocks.UseCase{}
	handler := New(mockUseCase)

	// Тестовый nextHandler
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mockUseCase.On("GetJWTToken").Return([]byte("secret"))

	tests := []struct {
		name               string
		token              string
		expectedStatusCode int
	}{
		{"No Token", "", http.StatusUnauthorized},
		{"Invalid Token Format", "Bearer invalidtoken", http.StatusUnauthorized},
		{"Valid Token", "Bearer " + createTestToken("secret", "12345"), http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/", nil)
			req.Header.Set("Authorization", tt.token)

			rr := httptest.NewRecorder()
			authMiddleware := handler.Auth(nextHandler)

			authMiddleware.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatusCode)
			}
		})
	}
}

// createTestToken создает тестовый JWT токен
func createTestToken(secret, userID string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userID,
	})
	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString
}
