package main

import (
	"sync"

	"./controllers"
	"./repositories"
	"./services"
	"./interfaces"
)

type IServiceContainer interface {
	InjectAccountController() controllers.AccountController
}

type kernel struct{
	AccountDatabase interfaces.IAccountDatabase
	ValidatorService interfaces.IValidatorService
}

func (k *kernel) InjectAccountController() controllers.AccountController {
	AccountRepository := &repositories.AccountRepository{k.AccountDatabase}
	AccountTypeRepository := &repositories.AccountTypeRepository{k.AccountDatabase}
	AccountStateRepository := &repositories.AccountStateRepository{k.AccountDatabase}
	AccountService := &services.AccountService{AccountRepository, AccountTypeRepository, AccountStateRepository, k.ValidatorService}
	AccountController := controllers.AccountController{AccountService, k.ValidatorService}

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
