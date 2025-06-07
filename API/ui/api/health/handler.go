package health

import (
	"encoding/json"
	"net/http"

	"API/config"
	"API/usecase"
)

type Handler struct {
	healthUC *usecase.HealthUsecase
}

func NewHandler(healthUC *usecase.HealthUsecase) *Handler {
	return &Handler{healthUC: healthUC}
}

func (h *Handler) CheckAPIConnection(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"status": "ok",
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		config.ErrInternalServer.Write(w)
		return
	}
}

func (h *Handler) GetDBTexts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	texts, err := h.healthUC.GetTestTexts(ctx)
	if err != nil {
		config.ErrInternalServer.Write(w)
		return
	}

	resp := DBQueryResponse{
		Texts: texts,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		config.ErrInternalServer.Write(w)
		return
	}
}

func (h *Handler) InsertDBText(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req DBCommandRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		config.ErrInvalidParam.Write(w)
		return
	}

	if err := h.healthUC.InsertTestText(ctx, req.Text); err != nil {
		config.ErrInternalServer.Write(w)
		return
	}

	w.WriteHeader(http.StatusCreated)
	resp := DBCommandResponse{
		Result: "inserted",
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		config.ErrInternalServer.Write(w)
		return
	}
}
