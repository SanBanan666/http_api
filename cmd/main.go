package main

import (
	"github.com/go-chi/chi/v5"
	"http_api/api"
	"http_api/service"
	"net/http"
)

func main() {
	service.StartWorker()

	r := chi.NewRouter()
	api.RegisterRoutes(r)

	http.ListenAndServe(":8080", r)
}
