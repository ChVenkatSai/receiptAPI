package api

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/ChVenkatSai/receiptAPI/pkg/models"
	"github.com/ChVenkatSai/receiptAPI/internal/service"
)

type Handler struct {
    Service *service.Service
}

func NewHandler(s *service.Service) *Handler {
    return &Handler{Service: s}
}

//Processes Receipt
func (h *Handler) ProcessReceipt(w http.ResponseWriter, r *http.Request) {
    var receipt models.Receipt
    if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    id := h.Service.ProcessReceipt(receipt)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"id": id})
}

//Gets Points
func (h *Handler) GetPoints(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    points, err := h.Service.GetPoints(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(map[string]int{"points": points})
}
