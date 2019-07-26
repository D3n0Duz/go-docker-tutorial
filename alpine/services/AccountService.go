package services

import (
	"../interfaces"
	"../models"
	"github.com/rs/xid"
)


type AccountService struct {
	AccountRepository interfaces.IAccountRepository
	AccountTypeRepository interfaces.IAccountTypeRepository
	AccountStateRepository interfaces.IAccountStateRepository
	ValidatorService interfaces.IValidatorService
}

const DefaultAccountTypeName = "Basic"
const DefaultAccountStateName = "Inactive"

func (service *AccountService) GetAccount(accountid string) (models.AccountModel, models.Errors){
	data, err := service.AccountRepository.GetAccount(accountid)

	if err.Code() != 0{
		return models.AccountModel{}, err
	}
	return data, models.Errors{}
}

func (service *AccountService) GetAccountType(accountid string) (models.AccountTypeModel, models.Errors){
	data, err := service.GetAccount(accountid)

	if err.Code() != 0{
		return models.AccountTypeModel{}, err
	}

	accountType, errType := service.AccountTypeRepository.GetAccountType(data.AccountTypeId)

	if errType.Code() != 0{
		return models.AccountTypeModel{}, errType
	}

	return accountType, models.Errors{}
}

func (service *AccountService) GetAccountState(accountid string) (models.AccountStateModel, models.Errors){
	data, err := service.GetAccount(accountid)

	if err.Code() != 0{
		return models.AccountStateModel{}, err
	}

	accountState, errState := service.AccountStateRepository.GetAccountState(data.AccountStateId)

	if errState.Code() != 0{
		return models.AccountStateModel{}, errState
	}

	return accountState, models.Errors{}
}

func (service *AccountService) PostAccount(accountModel models.AccountModel) (models.AccountModel, models.Errors){

	accountModel.AccountId = xid.New().String()
	if accountModel.AccountTypeId == ""{
		accountModel.AccountTypeId = xid.New().String()
	}
	if accountModel.AccountStateId == ""{
		accountModel.AccountStateId = xid.New().String()
	}
		
	if !service.ValidatorService.IsFieldValid(accountModel){
		return models.AccountModel{}, models.Errors{"Invalid accountModel", 400}
	}
	
	data, err := service.AccountRepository.AddAccount(accountModel.AccountId, accountModel)

	if err.Code() != 0{
		return models.AccountModel{}, err
	}

	accountModelType :=models.AccountTypeModel{DefaultAccountTypeName, accountModel.AccountTypeId}
	_, errType := service.AccountTypeRepository.AddAccountType(accountModel.AccountTypeId, accountModelType)
	if errType.Code() != 0{
		return models.AccountModel{}, errType
	}

	accountModelState :=models.AccountStateModel{DefaultAccountStateName, accountModel.AccountStateId}
	_, errState := service.AccountStateRepository.AddAccountState(accountModel.AccountStateId, accountModelState)
	if errState.Code() != 0{
		return models.AccountModel{}, errState
	}

	return data, models.Errors{}
}

func (service *AccountService) PutAccount(accountid string, accountModel models.AccountModel) (models.AccountModel, models.Errors){
	
	accountModelDb, err := service.GetAccount(accountid)
	if err.Code() != 0{
		return models.AccountModel{}, err
	}

	errIsValid := service.ValidatorService.IsValid(accountModelDb, accountModel)
	if errIsValid.Code() != 0{
		return models.AccountModel{}, errIsValid
	}

	data, err := service.AccountRepository.UpdateAccount(accountid, accountModel)

	if err.Code() != 0{
		return models.AccountModel{}, err
	}
	return data, models.Errors{}
}	

func (service *AccountService) PutAccountState(accountid string, accountStateModel models.AccountStateModel) (models.AccountStateModel, models.Errors){
	
	accountModelDb, err := service.GetAccount(accountid)
	if err.Code() != 0{
		return models.AccountStateModel{}, err
	}

	if !service.ValidatorService.IsFieldValidState(accountStateModel){
		return models.AccountStateModel{}, models.Errors{"Invalid accountStateModel", 400}
	}

	data, err := service.AccountStateRepository.UpdateAccountState(accountModelDb.AccountStateId, accountStateModel)

	if err.Code() != 0{
		return models.AccountStateModel{}, err
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

