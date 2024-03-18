package main

import (
	"fmt"
	"net/http"

	"github.com/devbersi/finance-go/configs"
	"github.com/devbersi/finance-go/handlers"
	"github.com/go-chi/chi/v5"
)

func main(){
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	 r := chi.NewRouter()
	 r.Post("/", handlers.Create)
	 r.Delete("/{id}", handlers.Delete)
	 r.Get("/{id}", handlers.Get)
	 r.Get("/", handlers.List)
	 r.Put("/{id}", handlers.Update)

	http.ListenAndServe(fmt.Sprintf(": %s", configs.GetServerPort()), r)
}