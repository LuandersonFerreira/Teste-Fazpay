package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"main.go/configs"
	"main.go/handlers"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.CreateUser)
	r.Put("/", handlers.UpdateUser)
	r.Delete("/", handlers.DeleteUser)
	r.Get("/list", handlers.ListUser)
	r.Get("/", handlers.GetUser)

	err = http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
	fmt.Println(err)
}
