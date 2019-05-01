package app

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gabrielerzinger/goCoupon/models"
)

// SaveHandler struct
type SaveHandler struct {
	App *App
}

// NewSaveHandler ctor
func NewSaveHandler(a *App) *SaveHandler {
	m := &SaveHandler{
		App: a,
	}
	return m
}

// ServeHTTP method
func (s *SaveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request models.CouponRequest

	if r.Body == nil {
		WriteError(w, http.StatusBadRequest, "request body shouldnt be empty", errors.New("Empty Body"))
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		WriteError(w, http.StatusBadRequest, "failed to decode request", err)
		return
	}

	err := s.App.Usecase.CreateCoupon(request.Name, request.DiscountType,
		request.Amount, request.CartPrice, time.Now())

	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Failed to save new coupon", err)
		return
	}

	s.App.Logger.Info("Saved new coupon")
	Write(w, http.StatusOK, "Saved new coupon succesfully")
}
