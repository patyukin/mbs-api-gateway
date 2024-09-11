package handler

import (
	"context"
	"github.com/patyukin/mbs-api-gateway/internal/handler/mocks"
	"github.com/patyukin/mbs-api-gateway/pkg/rate_limiter"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func TestHandler_RateLimitMiddleware(t *testing.T) {
	mockUseCase := &mocks.UseCase{}
	handler := New(mockUseCase)

	limiter := rate_limiter.New(context.Background(), 1, time.Second)

	rateLimitMiddleware := handler.RateLimitMiddleware(limiter, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// Тесты
	tests := []struct {
		name               string
		numRequests        int
		expectedStatusCode int
		expectedAllowed    int
		expectedBlocked    int
	}{
		{
			name:               "Concurrent requests under limit",
			numRequests:        1,
			expectedStatusCode: http.StatusOK,
			expectedAllowed:    1,
			expectedBlocked:    0,
		},
		{
			name:               "Concurrent requests over limit",
			numRequests:        5,
			expectedStatusCode: http.StatusTooManyRequests,
			expectedAllowed:    1,
			expectedBlocked:    4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var wg sync.WaitGroup
			allowedCount := 0
			blockedCount := 0

			results := make(chan int, tt.numRequests)

			for i := 0; i < tt.numRequests; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					rr := httptest.NewRecorder()
					req, _ := http.NewRequest("GET", "/", nil)

					rateLimitMiddleware.ServeHTTP(rr, req)
					results <- rr.Code
				}()
			}

			wg.Wait()
			close(results)

			for res := range results {
				if res == http.StatusOK {
					allowedCount++
				} else if res == http.StatusTooManyRequests {
					blockedCount++
				}
			}

			if allowedCount != tt.expectedAllowed {
				t.Errorf("expected %d allowed requests, got %d", tt.expectedAllowed, allowedCount)
			}

			if blockedCount != tt.expectedBlocked {
				t.Errorf("expected %d blocked requests, got %d", tt.expectedBlocked, blockedCount)
			}

			time.Sleep(1 * time.Second)
		})
	}
}
