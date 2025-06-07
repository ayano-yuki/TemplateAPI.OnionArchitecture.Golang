package health

import (
	"encoding/json"
	"net/http"

	"API/usecase"
)

type Handler struct {
	healthUC *usecase.HealthUsecase
}

func NewHandler(healthUC *usecase.HealthUsecase) *Handler {
	return &Handler{healthUC: healthUC}
}

// CheckAPIConnection はAPI接続確認用ハンドラです。
func (h *Handler) CheckAPIConnection(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK!"))
}

// GetDBTexts はDBからtext一覧を取得するハンドラです。
func (h *Handler) GetDBTexts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	texts, err := h.healthUC.GetTestTexts(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(struct {
		Texts []string `json:"texts"`
	}{
		Texts: texts,
	})
}

// InsertDBText はリクエストのtextをDBに挿入するハンドラです。
func (h *Handler) InsertDBText(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req struct {
		Text string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.healthUC.InsertTestText(ctx, req.Text); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
