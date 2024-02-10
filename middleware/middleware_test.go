package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestLoggingMiddleware(t *testing.T) {
	// Create a test handler
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate some processing time
		time.Sleep(50 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	})

	// Create a request to test the middleware
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// Create the middleware handler
	loggingMiddleware := LoggingMiddleware(testHandler)

	// Serve the request through the middleware
	loggingMiddleware.ServeHTTP(rr, req)


	
	// Check if the status code is set correctly
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}
}
