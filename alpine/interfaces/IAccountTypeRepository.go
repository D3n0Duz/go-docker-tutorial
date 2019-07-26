package interfaces

import (
	"../models"
)

type IAccountTypeRepository interface {
	GetAccountType(accountid string) (models.AccountTypeModel, models.Errors)
	AddAccountType(accountid string, accountTypeModel models.AccountTypeModel) (models.AccountTypeModel, models.Errors)
	UpdateAccount(accountid string, accountTypeModel models.AccountTypeModel) (models.AccountTypeModel, models.Errors)
}
