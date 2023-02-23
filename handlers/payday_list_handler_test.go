package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetNextPaydays(t *testing.T) {
	testCases := []struct {
		name       string
		path       string
		statusCode int
	}{
		{
			name:       "invalid URL",
			path:       "/invalid",
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "invalid pay day",
			path:       "/till-salary/pay-day/32/list-dates",
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "valid pay day",
			path:       "/till-salary/pay-day/15/list-dates",
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tc.path, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			recorder := httptest.NewRecorder()
			handler := http.HandlerFunc(GetNextPaydays)
			handler.ServeHTTP(recorder, req)

			if recorder.Code != tc.statusCode {
				t.Errorf("got status code %d, want %d", recorder.Code, tc.statusCode)
			}
		})
	}
}
