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

	var controllersss controllers.AccountController
	r := chi.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world!")
	})

	r.Get("/accounts/{accountid}", controllersss.GetAccount)
	r.Post("/accounts/{accountid}", controllersss.PostAccount)
	r.Put("/accounts/{accountid}", controllersss.PutAccount)
	r.Delete("/accounts/{accountid}", controllersss.DeleteAccount)

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
