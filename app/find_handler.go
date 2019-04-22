package app

import (
	"encoding/json"
	"net/http"
)

// FindHandler struct
type FindHandler struct {
	App *App
}

// NewFindHandler ctor
func NewFindHandler(a *App) *FindHandler {
	m := &FindHandler{
		App: a,
	}
	return m
}

// ServeHTTP method
func (s *FindHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	couponName := query.Get("name")

	if couponName == "" {
		Write(w, http.StatusBadRequest, "Missed coupon name param")
		return
	}

	s.App.Logger.Info("Fulfilling find request")

	coupon, err := s.App.Usecase.FindCoupon(couponName)

	if err != nil || coupon.DiscountType == "" {
		WriteBytes(w, http.StatusNotFound, []byte("Coupon not found"))
		return
	}

	couponJSON, err := json.Marshal(coupon)

	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Failed to marshal json", err)
	}

	WriteSuccessWithJSON(w, http.StatusOK, couponJSON, "Success")
	return
}
