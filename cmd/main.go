package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"81HW/internal/config"
	"81HW/internal/handler"
	"81HW/internal/repository"
	"81HW/pkg/database"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config yuklashda xato: %v", err)
	}

	db, err := database.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Ma'lumotlar bazasiga ulanishda xato: %v", err)
	}
	defer db.Close()

	repo := repository.NewPostgresRepository(db)
	handler := handler.NewBookHandler(repo)

	r := mux.NewRouter()
	r.HandleFunc("/books", handler.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", handler.GetBook).Methods("GET")
	r.HandleFunc("/books", handler.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", handler.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", handler.DeleteBook).Methods("DELETE")

	log.Printf("Server %s portida ishga tushdi", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}