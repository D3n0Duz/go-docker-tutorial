package main

import (
	"fmt"
	"net/http"
	"sync"
	"github.com/go-chi/chi"
	"./infrastructures"
	"./services"
)

type IChiRouter interface {
	InitRouter() *chi.Mux
}

type router struct{}

func (router *router) InitRouter() *chi.Mux {

	AccountDatabase := &infrastructures.AccountDatabase{}
	ValidatorService := &services.ValidatorService{}
	AccountDatabase.InitDatabase()
	ServiceContainer := kernel{AccountDatabase, ValidatorService}

	accountController := ServiceContainer.InjectAccountController()
	r := chi.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world!")
	})

	r.Get("/accounts/{accountid}", accountController.GetAccount)
	r.Post("/accounts", accountController.PostAccount)
	r.Put("/accounts/{accountid}", accountController.PutAccount)
	r.Delete("/accounts/{accountid}", accountController.DeleteAccount)
	
	r.Get("/accounts/{accountid}/types", accountController.GetAccountType)
	r.Get("/accounts/{accountid}/states", accountController.GetAccountState)
	r.Put("/accounts/{accountid}/states", accountController.PutAccountState)
	//r.Put("/accounts/{accountid}/types/{accountTypeId}", accountTypeController.PutAccount) TODO later

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
