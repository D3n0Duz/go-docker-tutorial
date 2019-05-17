package main

import (
	"fmt"
	"net/http"
	"sync"
	"github.com/go-chi/chi"
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {

	accountController := ServiceContainer().InjectAccountController()
	r := chi.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world!")
	})

	r.Get("/accounts/{accountid}", accountController.GetAccount)
	r.Post("/accounts", accountController.PostAccount)
	r.Put("/accounts/{accountid}", accountController.PutAccount)
	r.Delete("/accounts/{accountid}", accountController.DeleteAccount)

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
