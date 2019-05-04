package services

import (
	"../interfaces"
	"../models"
)

type AccountService strut {
	AccountRepository interfaces.IAccountService
}

func (service *AccountService) GetAccount(accountID string) {
	
	var result string
	
	// Validate that the loged in user can only get his own account
	



}