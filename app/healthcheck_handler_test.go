package app_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gabrielerzinger/goCoupon/app"
	helpers "github.com/gabrielerzinger/goCoupon/testing"
	"github.com/stretchr/testify/assert"
)

func TestHealthcheckHandler(t *testing.T) {
	appTest := helpers.GetApp(t)
	handler := app.NewHealthcheckHandler(appTest)

	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	res := httptest.NewRecorder()

	handler.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}
