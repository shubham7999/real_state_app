package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleError(t *testing.T) {
	testCases := []struct {
		Name           string
		StatusCode     int
		Message        string
		ExpectedStatus int
		ExpectedBody   map[string]string
	}{
		{
			Name:           "HandlesErrorCorrectly",
			StatusCode:     http.StatusBadRequest,
			Message:        "Invalid input",
			ExpectedStatus: http.StatusBadRequest,
			ExpectedBody:   map[string]string{"error": "Invalid input"},
		},
		{
			Name:           "HandlesDifferentStatusCode",
			StatusCode:     http.StatusInternalServerError,
			Message:        "Internal Server Error",
			ExpectedStatus: http.StatusInternalServerError,
			ExpectedBody:   map[string]string{"error": "Internal Server Error"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Create a response recorder to capture the response
			rr := httptest.NewRecorder()

			// Call the HandleError function with the test case parameters
			HandleError(rr, tc.StatusCode, tc.Message)

			// Check the status code
			if rr.Code != tc.ExpectedStatus {
				t.Errorf("Expected status code %d, got %d", tc.ExpectedStatus, rr.Code)
			}

			// Check the response body
			var responseBody map[string]string
			err := json.Unmarshal(rr.Body.Bytes(), &responseBody)
			if err != nil {
				t.Errorf("Error decoding response body: %v", err)
			}

			// Check that the response body matches the expected body
			if responseBody["error"] != tc.ExpectedBody["error"] {
				t.Errorf("Expected error message '%s', got '%s'", tc.ExpectedBody["error"], responseBody["error"])
			}
		})
	}
}
