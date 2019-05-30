package interfaces

import (
	"../models"
)

type IAccountRepository interface {
	GetAccount(accountid string) (models.AccountModel, error)
	AddAccount(accountid string, accountModel models.AccountModel) (models.AccountModel, error)
	UpdateAccount(accountid string, accountModel models.AccountModel) (models.AccountModel, error)
	DeleteAccount(accountid string) error
}
