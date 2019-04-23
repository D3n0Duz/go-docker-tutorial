package main

import (
	"fmt"
	"net/http"
	"sync"

	"./controllers"

	"github.com/go-chi/chi"
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {

	var controllersss controllers.UserController
	r := chi.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world!")
	})

	r.Get("/users/{userid}", controllersss.GetUser)
	r.Post("/users/{userid}", controllersss.PostUser)
	r.Put("/users/{userid}", controllersss.PutUser)
	r.Delete("/users/{userid}", controllersss.DeleteUser)

	return r
}

var (
	m          *router
	routerOnce sync.Once
)

func ChiRouter() IChiRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}
