package app_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gabrielerzinger/goCoupon/app"
	helpers "github.com/gabrielerzinger/goCoupon/testing"
	"github.com/stretchr/testify/assert"
)

func TestFindHandler(t *testing.T) {
	appTest := helpers.GetApp(t)
	handler := app.NewFindHandler(appTest)

	tables := map[string]struct {
		request *http.Request
		asserts func(response *httptest.ResponseRecorder)
	}{
		"test_get_coupon_404": {
			request: func() *http.Request {
				request, err := http.NewRequest("GET", "/coupon?name=ARABELLA", nil)
				assert.NoError(t, err)
				return request
			}(),
			asserts: func(response *httptest.ResponseRecorder) {
				assert.Contains(t, response.Body.String(), "Coupon not found")
				assert.Equal(t, http.StatusNotFound, response.Code)

			},
		},
		"test_bad_request": {
			request: func() *http.Request {
				request, err := http.NewRequest("GET", "/coupon?name=", nil)
				assert.NoError(t, err)
				return request
			}(),
			asserts: func(response *httptest.ResponseRecorder) {
				assert.Contains(t, response.Body.String(), "Missed coupon name param")
				assert.Equal(t, http.StatusBadRequest, response.Code)
			},
		},
		"test_found_coupon": {
			request: func() *http.Request {
				request, err := http.NewRequest("GET", "/coupon?name=OFF12", nil)
				assert.NoError(t, err)
				return request
			}(),
			asserts: func(response *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, response.Code)
			},
		},
	}

	for name, table := range tables {
		t.Run(name, func(t *testing.T) {
			response := httptest.NewRecorder()
			handler.ServeHTTP(response, table.request)
			table.asserts(response)
		})
	}
}
