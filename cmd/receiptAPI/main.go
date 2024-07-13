package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/ChVenkatSai/receiptAPI/internal/api"
    "github.com/ChVenkatSai/receiptAPI/internal/service"
    "github.com/ChVenkatSai/receiptAPI/internal/storage"
)

func main() {
    r := mux.NewRouter()
    store := storage.NewInMemoryStorage()
    svc := service.NewService(store)
    h := api.NewHandler(svc)

    r.HandleFunc("/receipts/process", h.ProcessReceipt).Methods("POST")
    r.HandleFunc("/receipts/{id}/points", h.GetPoints).Methods("GET")

    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
