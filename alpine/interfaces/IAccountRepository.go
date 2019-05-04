package interfaces

import (
	"../models"
)

type IAccountRepository interface {
	GetAccount(accountid string) (models.AccountModel, error)
	PostAccount(accountid string) (models.AccountModel, error)
	PutAccount(accountid string) (models.AccountModel, error)
	DeleteAccount(accountid string) (models.AccountModel, error)
}
