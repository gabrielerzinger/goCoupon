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
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSaveHandler(t *testing.T) {
	appTest := helpers.GetApp(t)
	handler := app.NewSaveHandler(appTest)

	tables := map[string]struct {
		request *http.Request
		asserts func(response *httptest.ResponseRecorder)
	}{
		"test_empty_body": {
			request: func() *http.Request {
				request, err := http.NewRequest("POST", "/coupon", nil)
				assert.NoError(t, err)
				return request
			}(),
			asserts: func(response *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, response.Code)
			},
		},
		"test_bad_request": {
			request: func() *http.Request {
				body := "bad bodyy"
				bts, _ := json.Marshal(body)

				request, err := http.NewRequest("POST", "/coupon", bytes.NewReader(bts))
				assert.NoError(t, err)
				return request
			}(),
			asserts: func(response *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, response.Code)
			},
		},
		"test_save_succes": {
			request: func() *http.Request {
				body := &models.CouponRequest{
					Name:         uuid.New().String(),
					DiscountType: "VALUE",
					Amount:       123,
					CartPrice:    500,
				}
				bts, _ := json.Marshal(body)

				request, err := http.NewRequest("POST", "/coupon", bytes.NewReader(bts))
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
