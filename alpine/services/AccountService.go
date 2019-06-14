package services

import (
	"../interfaces"
	"../models"
	"github.com/rs/xid"
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

func (service *AccountService) GetAccount(accountid string) (models.AccountModel, models.Errors){
	data, err := service.AccountRepository.GetAccount(accountid)

	if err.Code() != 0{
		return models.AccountModel{}, err
	}
	return data, models.Errors{}
}

func (service *AccountService) PostAccount(accountModel models.AccountModel) (models.AccountModel, models.Errors){

	id := xid.New()
	accountModel.AccountId = id.String()
	if !isFieldValid(accountModel){
		return models.AccountModel{}, models.Errors{"Invalid accountModel", 400}
	}
	
	data, err := service.AccountRepository.AddAccount(accountModel.AccountId, accountModel)

	if err.Code() != 0{
		return models.AccountModel{}, err
	}
	return data, models.Errors{}
}

func (service *AccountService) PutAccount(accountid string, accountModel models.AccountModel) (models.AccountModel, models.Errors){
	
	errIsValid := service.isValid(accountid, accountModel)
	if errIsValid.Code() != 0{
		return models.AccountModel{}, errIsValid
	}

	data, err := service.AccountRepository.UpdateAccount(accountid, accountModel)

	if err.Code() != 0{
		return models.AccountModel{}, err
	}
	return data, models.Errors{}
}	
func (service *AccountService) DeleteAccount(accountid string) models.Errors{
	_, err := service.GetAccount(accountid)
	if err.Code() != 0{
		return err
	}
	return service.AccountRepository.DeleteAccount(accountid)
}

func isFieldValid (accountModel models.AccountModel) bool{
	emailMatch, _ := regexp.MatchString(emailRegex, accountModel.Email)
	accountStateIDMatch, _ := regexp.MatchString(stateIdRegex, strconv.Itoa(accountModel.AccountStateID))
	accountTypeIDMatch, _ := regexp.MatchString(typeIdRegex, strconv.Itoa(accountModel.AccountTypeID))
	return emailMatch && accountStateIDMatch && accountTypeIDMatch
}


func (service *AccountService) isValid(accountid string, accountModel models.AccountModel) models.Errors{
	data, err := service.GetAccount(accountid)
	if err.Code() != 0{
		return err
	}

	if data.Email != accountModel.Email{
		errMessage := "Email does not match"
		fmt.Println(errMessage)
		return models.Errors{errMessage, 400}
	}

	return models.Errors{}
}