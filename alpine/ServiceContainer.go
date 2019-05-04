package main

import (
	"sync"

	"./controllers"
	"./repositories"
	"./services"
	
)

type IServiceContainer interface {
	InjectAccountController() controllers.AccountController
}

type kernel struct{}

func (k *kernel) InjectAccountController() controllers.AccountController {
	
	AccountRepository := &repositories.AccountRepository{}
	AccountService := &services.AccountService{AccountRepository}
	AccountController := controllers.AccountController{AccountService}


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