package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetNextPaydays(t *testing.T) {
	testCases := []struct {
		name         string
		path         string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "invalid URL",
			path:         "/till-salary/pay-day/32/list-dates/invalid",
			expectedCode: http.StatusBadRequest,
			expectedBody: "Invalid URL\n",
		},
		{
			name:         "invalid pay day",
			path:         "/till-salary/pay-day/32/list-dates",
			expectedCode: http.StatusBadRequest,
			expectedBody: "Invalid pay day\n",
		},
		{
			name:         "invalid pay day, string",
			path:         "/till-salary/pay-day/ssss/list-dates",
			expectedCode: http.StatusBadRequest,
			expectedBody: "Invalid pay day\n",
		},
		{
			name:         "invalid pay day, empty",
			path:         "/till-salary/pay-day//list-dates",
			expectedCode: http.StatusBadRequest,
			expectedBody: "Invalid pay day\n",
		},
		{
			name:         "valid pay day",
			path:         "/till-salary/pay-day/15/list-dates",
			expectedCode: http.StatusOK,
			expectedBody: `{"dates":["2023-03-15","2023-04-14","2023-05-15","2023-06-15","2023-07-14","2023-08-15","2023-09-15","2023-10-13","2023-11-15","2023-12-15"]}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tc.path, nil)
			if err != nil {
				t.Errorf("could not create request: %v", err)
			}

			recorder := httptest.NewRecorder()
			handler := http.HandlerFunc(GetNextPaydays)

			handler.ServeHTTP(recorder, req)

			if recorder.Code != tc.expectedCode {
				t.Errorf("got status code %d, want %d", recorder.Code, tc.expectedCode)
			}

			if recorder.Body.String() != tc.expectedBody {
				t.Errorf("Expected response body %q, but got %q", tc.expectedBody, recorder.Body.String())
			}
		})
	}
}
