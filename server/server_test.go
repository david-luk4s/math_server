package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type TestWeather struct {
	Name          string
	Server        *httptest.Server
	Response      *Weather
	ExpectedError error
}

func TestGetWeather(t *testing.T) {
	cases := []TestWeather{
		{
			Name: "request-get-weather",
			Server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"city": "Paraiba", "forecast": "20c"}`))
			})),
			Response:      &Weather{City: "Paraiba", Forecast: "20c"},
			ExpectedError: nil,
		},
	}

	for _, ct := range cases {
		resp, err := GetWeather(ct.Server.URL)

		t.Run(ct.Name, func(t *testing.T) {
			if !reflect.DeepEqual(resp, ct.Response) {
				t.Errorf("FAILED: expected %v, got %v", ct.Response, resp)
			}

			if !errors.Is(ct.ExpectedError, err) {
				t.Errorf("Expected Error FAILED: expected %v, got %d", ct.ExpectedError, err)
			}
		})
	}
}
