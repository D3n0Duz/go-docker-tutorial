package interfaces

import (
	"../models"
)

type IAccountRepository interface {
	GetAccount(accountid string) (models.AccountModel, error)
	CreateOrUpdateAccount(accountid string, accountModel models.AccountModel) (models.AccountModel, error)
	DeleteAccount(accountid string) error
}
