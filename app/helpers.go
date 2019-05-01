package app

import (
	"fmt"
	"net/http"
)

//Write to response with status code
func Write(w http.ResponseWriter, status int, text string) {
	WriteBytes(w, status, []byte(text))
}

//WriteBytes to the response and with the status code
func WriteBytes(w http.ResponseWriter, status int, text []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(text)
}

// WriteError to the response with message
func WriteError(w http.ResponseWriter, status int, errorMsg string, err error) {
	errMsg := fmt.Sprintf(`{"success":false, "message":"%s", "reason": "%s"}`, errorMsg, err.Error())
	Write(w, status, errMsg)
}

// WriteSuccessWithJSON sends response with statusOK to request
func WriteSuccessWithJSON(w http.ResponseWriter, status int, res []byte, msg string) {
	retMsg := fmt.Sprintf(`{"success" : true, "response":%s}`, res)
	Write(w, status, retMsg)
}
