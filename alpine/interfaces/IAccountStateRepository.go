package interfaces

import (
	"../models"
)

type IAccountStateRepository interface {
	GetAccountState(accountid string) (models.AccountStateModel, models.Errors)
	AddAccountState(accountid string, accountStateModel models.AccountStateModel) (models.AccountStateModel, models.Errors)
	UpdateAccountState(accountid string, accountStateModel models.AccountStateModel) (models.AccountStateModel, models.Errors)
}
