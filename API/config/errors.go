package config

import (
	"encoding/json"
	"net/http"
)

type AppError struct {
	Code       string
	Message    string
	HTTPStatus int
}

// 共通エラー定義
var (
	ErrInternalServer = AppError{Code: "E1000", Message: "Internal server error", HTTPStatus: http.StatusInternalServerError}
	ErrInvalidParam   = AppError{Code: "E1001", Message: "Invalid parameter", HTTPStatus: http.StatusBadRequest}

	// DB関連
	ErrDBQuery  = AppError{Code: "E2001", Message: "Database query failed", HTTPStatus: http.StatusInternalServerError}
	ErrDBInsert = AppError{Code: "E2002", Message: "Database insert failed", HTTPStatus: http.StatusInternalServerError}
)

func (e AppError) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.HTTPStatus)
	json.NewEncoder(w).Encode(map[string]string{
		"code":    e.Code,
		"message": e.Message,
	})
}
