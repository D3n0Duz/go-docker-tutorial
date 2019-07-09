package interfaces

import (
	"../models"
)

type IAccountService interface {
	GetAccount(accountid string) (models.AccountModel, models.Errors)
	PostAccount(accountModel models.AccountModel) (models.AccountModel, models.Errors)
	PutAccount(accountid string, accountModel models.AccountModel) (models.AccountModel, models.Errors)
	DeleteAccount(accountid string)  models.Errors

	GetAccountType(accountid string) (models.AccountTypeModel, models.Errors)
	GetAccountState(accountid string) (models.AccountStateModel, models.Errors)
}
