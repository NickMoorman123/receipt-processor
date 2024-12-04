package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/NickMoorman123/receipt-processor/handlers"
	"github.com/NickMoorman123/receipt-processor/store"
)

type Args struct {
	// postgres connection string, of the form,
	// e.g "postgres://user:password@localhost:5432/database?sslmode=disable"
	conn string
	// port for the server of the form, e.g ":8080"
	port string
}

func Run(args Args) error {
	router := mux.NewRouter().
		PathPrefix("/api/v1/").
		Subrouter()

	st := store.NewPostgresReceiptStore(args.conn)
	hnd := handlers.NewReceiptHandler(st)
	RegisterAllRoutes(router, hnd)

	log.Println("Starting server at port: ", args.port)
	return http.ListenAndServe(args.port, router)
}

func RegisterAllRoutes(router *mux.Router, hnd handlers.IReceiptHandler) {
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	router.HandleFunc("/receipts/{id}/points", hnd.Get).Methods(http.MethodGet)
	router.HandleFunc("/receipts/process", hnd.Process).Methods(http.MethodPost)
}