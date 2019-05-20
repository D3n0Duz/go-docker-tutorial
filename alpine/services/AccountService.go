package services

import (
	"../interfaces"
	"../models"
	"github.com/rs/xid"
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

func (service *AccountService) PostAccount(accountModel models.AccountModel) (models.AccountModel, error){

	id := xid.New()
	accountModel.AccountId = id.String()
	data, err := service.AccountRepository.CreateOrUpdateAccount(accountModel.AccountId, accountModel)

	if err != nil {
		return models.AccountModel{}, err
	}
	return data, nil
}

func (service *AccountService) PutAccount(accountid string, accountModel models.AccountModel) (models.AccountModel, error){
	data, err := service.AccountRepository.CreateOrUpdateAccount(accountid, accountModel)

	if err != nil {
		return models.AccountModel{}, err
	}
	return data, nil
}	
func (service *AccountService) DeleteAccount(accountid string) error{
	return service.AccountRepository.DeleteAccount(accountid)
}