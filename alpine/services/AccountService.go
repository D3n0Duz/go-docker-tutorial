package services

import (
	"../interfaces"
	"../models"
)

type AccountService struct {
	AccountRepository interfaces.IAccountRepository
}

func (service *AccountService) GetAccount(accountid string) (models.AccountModel, error){
	data, err := service.AccountRepository.GetAccount(accountid)

	if err != nil {
		return models.AccountModel{}, err
	}
	return data, nil
}

func (service *AccountService) PostAccount(accountid string) (models.AccountModel, error){
	return models.AccountModel{}, nil
}

func (service *AccountService) PutAccount(accountid string) (models.AccountModel, error){
	return models.AccountModel{}, nil
}	

func (service *AccountService) DeleteAccount(accountid string) (models.AccountModel, error){
	return models.AccountModel{}, nil
}