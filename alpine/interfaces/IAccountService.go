package interfaces

import (
	"../models"
)

type IAccountService interface {
	GetAccount(accountid string) (models.AccountModel, error)
	PostAccount(accountModel models.AccountModel) (models.AccountModel, error)
	PutAccount(accountid string, accountModel models.AccountModel) (models.AccountModel, error)
	DeleteAccount(accountid string)  error
}
