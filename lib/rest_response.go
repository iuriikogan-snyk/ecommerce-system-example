package lib

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Error *Error      `json:"error"`
	Data  interface{} `json:"data"`
}

func failResponseWriter(w http.ResponseWriter, err error, errStatusCode int) {
	w.Header().Set("Content-Type", "application/json")

	var resp Response
	w.WriteHeader(errStatusCode)
	resp.Error = &Error{Code: errStatusCode, Message: err.Error()}
	resp.Data = nil

	responseBytes, _ := json.Marshal(resp)
	w.Write(responseBytes)
}

func successResponseWriter(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")

	var resp Response
	w.WriteHeader(statusCode)
	resp.Error = nil
	resp.Data = data

	responseBytes, _ := json.Marshal(resp)
	w.Write(responseBytes)
}

func WriteResponse(w http.ResponseWriter, err error, statusCode int, data any) {
	switch err.(type) {
	case *ErrNotFound, ErrNotFound, *ErrBadRequest, ErrBadRequest:
		failResponseWriter(w, err, statusCode)
		return
	case nil:
		successResponseWriter(w, data, statusCode)
		return
	default:
		failResponseWriter(w, err, http.StatusInternalServerError)
		return
	}
}
