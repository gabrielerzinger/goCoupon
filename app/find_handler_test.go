package app_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gabrielerzinger/goCoupon/app"
	"github.com/gabrielerzinger/goCoupon/models"
	helpers "github.com/gabrielerzinger/goCoupon/testing"
	"github.com/stretchr/testify/assert"
)

func TestFindHandler(t *testing.T) {
	appTest := helpers.GetApp(t)
	handler := app.NewFindHandler(appTest)
	saveHandler := app.NewSaveHandler(appTest)

	body := &models.CouponRequest{
		Name:         "ARABELA",
		DiscountType: "VALUE",
		Amount:       123,
		CartPrice:    500,
	}
	bts, _ := json.Marshal(body)

	requestSave, err := http.NewRequest("POST", "/coupon", bytes.NewReader(bts))
	assert.NoError(t, err)

	responseSave := httptest.NewRecorder()
	saveHandler.ServeHTTP(responseSave, requestSave)

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
				request, err := http.NewRequest("GET", "/coupon?name=ARABELA", nil)
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
