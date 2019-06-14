package interfaces

import (
	"../models"
)

type IAccountRepository interface {
	GetAccount(accountid string) (models.AccountModel, models.Errors)
	AddAccount(accountid string, accountModel models.AccountModel) (models.AccountModel, models.Errors)
	UpdateAccount(accountid string, accountModel models.AccountModel) (models.AccountModel, models.Errors)
	DeleteAccount(accountid string) models.Errors
}
