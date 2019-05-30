package services

import (
	"../interfaces"
	"../models"
	"github.com/rs/xid"
	"errors"
	"fmt"
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
	data, err := service.AccountRepository.AddAccount(accountModel.AccountId, accountModel)

	if err != nil {
		return models.AccountModel{}, err
	}
	return data, nil
}

func (service *AccountService) PutAccount(accountid string, accountModel models.AccountModel) (models.AccountModel, error){
	
	errIsValid := service.isValid(accountid, accountModel)
	if errIsValid != nil{
		return models.AccountModel{}, errors.New("Invalid accountId or email adress")
	}

	data, err := service.AccountRepository.UpdateAccount(accountid, accountModel)

	if err != nil {
		return models.AccountModel{}, err
	}
	return data, nil
}	
func (service *AccountService) DeleteAccount(accountid string) error{
	return service.AccountRepository.DeleteAccount(accountid)
}

func (service *AccountService) isValid(accountid string, accountModel models.AccountModel) error{
	data, err := service.GetAccount(accountid)
	if err != nil{
		return err
	}

	if data.Email != accountModel.Email{
		errMessage := "Email does not match"
		fmt.Println(errMessage)
		return errors.New(errMessage)
	}
	return nil
}