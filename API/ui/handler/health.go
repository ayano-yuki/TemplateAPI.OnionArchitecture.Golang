package handler

import (
	"encoding/json"
	"net/http"
	"templateapi/onionarchitecture/golang/usecase"
)

type HealthHandler struct {
	Usecase *usecase.HealthUsecase
}

func NewHealthHandler(u *usecase.HealthUsecase) *HealthHandler {
	return &HealthHandler{Usecase: u}
}

func (h *HealthHandler) HealthAPI(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

func (h *HealthHandler) HealthDB1(w http.ResponseWriter, r *http.Request) {
	data, err := h.Usecase.GetTestTexts()
	if err != nil {
		http.Error(w, "DB fetch error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (h *HealthHandler) HealthDB2(w http.ResponseWriter, r *http.Request) {
	if err := h.Usecase.InsertText(); err != nil {
		http.Error(w, "DB insert error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(`{"inserted":"文字列1"}`))
}
