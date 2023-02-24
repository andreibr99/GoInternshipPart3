package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHowMuchUntilPayday(t *testing.T) {
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
			path:       "/till-salary/how-much?pay_day=32",
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "invalid pay day, string",
			path:       "/till-salary/how-much?pay_day=sss",
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "invalid pay day, empty",
			path:       "/till-salary/how-much?pay_day=",
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "valid pay day",
			path:       "/till-salary/how-much?pay_day=25",
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tc.path, nil)
			if err != nil {
				t.Errorf("could not create request: %v", err)
			}

			recorder := httptest.NewRecorder()
			handler := http.HandlerFunc(HowMuchUntilPayday)
			handler.ServeHTTP(recorder, req)

			if recorder.Code != tc.statusCode {
				t.Errorf("got status code %d, want %d", recorder.Code, tc.statusCode)
			}
		})
	}
}
