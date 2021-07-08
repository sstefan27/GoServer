package main

import (
	"go/problem2/db"
	"go/problem2/rest"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	db.InitDatabase("postgres://postgres:password@localhost:5432/postgres?sslmode=disable", 1)
	router := chi.NewRouter()
	router.Route("/v1", func(r chi.Router) {
		r.Get("/getmyvehicle", rest.GetVehicle)
		r.Post("/postmyvehicle", rest.PostVehicle)
		r.Post("/person", rest.PostPerson)
		r.Get("/person", rest.GetPerson)
		r.Post("/canDrive", rest.PostCanDrive)

	})
	http.ListenAndServe(":8081", router)
}
