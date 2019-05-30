package services

import (
	"../interfaces"
	"../models"
	"github.com/rs/xid"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)
const emailRegex = "^[a-zA-Z0-9_+&*-]+(?:\\.[a-zA-Z0-9_+&*-]+)*@(?:[a-zA-Z0-9-]+\\.)+[a-zA-Z]{2,7}$"
const dateRegex = "^[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}Z$"
const typeIdRegex = "^[1-3]{1}$"
const stateIdRegex = "^[1-3]{1}$"

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
	if !isFieldValid(accountModel){
		return models.AccountModel{}, errors.New("Invalid accountModel")
	}
	
	data, err := service.AccountRepository.AddAccount(accountModel.AccountId, accountModel)

	if err != nil {
		return models.AccountModel{}, err
	}
	return data, nil
}

func (service *AccountService) PutAccount(accountid string, accountModel models.AccountModel) (models.AccountModel, error){
	
	errIsValid := service.isValid(accountid, accountModel)
	if errIsValid != nil{
		return models.AccountModel{}, errors.New("Invalid AccountModel provided")
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

func isFieldValid (accountModel models.AccountModel) bool{
	emailMatch, _ := regexp.MatchString(emailRegex, accountModel.Email)
	accountStateIDMatch, _ := regexp.MatchString(stateIdRegex, strconv.Itoa(accountModel.AccountStateID))
	accountTypeIDMatch, _ := regexp.MatchString(typeIdRegex, strconv.Itoa(accountModel.AccountTypeID))
	return emailMatch && accountStateIDMatch && accountTypeIDMatch
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