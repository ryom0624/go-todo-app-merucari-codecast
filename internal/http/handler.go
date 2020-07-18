package http

import (
	"encoding/json"
	"go-todoapp/internal/db"
	"go-todoapp/internal/todo"
	"net/http"

	"github.com/google/uuid"
)

var _ http.Handler = (*createHandler)(nil)

type createHandler struct {
	db db.DB
}

func (h *createHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var t todo.TODO
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	t.ID = uuid.New().String()
	if err := h.db.PutTODO(r.Context(), &t); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&t); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
