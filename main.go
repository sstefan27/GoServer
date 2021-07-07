package main

import (
	"go/problem2/rest"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()
	router.Route("/v1", func(r chi.Router) {
		r.Get("/getmyvehicle", rest.GetVehicle)

	})
	http.ListenAndServe(":8081", router)
}
