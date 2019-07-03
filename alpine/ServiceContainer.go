package main

import (
	"sync"

	"./controllers"
	"./infrastructures"
	"./repositories"
	"./services"
)

type IServiceContainer interface {
	InjectAccountController() controllers.AccountController
}

type kernel struct{}

func (k *kernel) InjectAccountController() controllers.AccountController {

	AccountDatabase := &infrastructures.AccountDatabase{}
	ValidatorService := &services.ValidatorService{}
	AccountDatabase.InitDatabase()
	AccountRepository := &repositories.AccountRepository{AccountDatabase}
	AccountService := &services.AccountService{AccountRepository}
	AccountController := controllers.AccountController{AccountService, ValidatorService}

	return AccountController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
