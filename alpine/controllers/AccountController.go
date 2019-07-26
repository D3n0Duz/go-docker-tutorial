package controllers

import (
	"net/http"
	"../interfaces"
	"encoding/json"
	"../models"
	"github.com/go-chi/chi"
)

type AccountController struct {
	AccountService interfaces.IAccountService
	ValidatorService interfaces.IValidatorService
}

func (controller *AccountController) GetAccount(w http.ResponseWriter, r *http.Request) {

	if !controller.ValidatorService.VerifyToken(r, w) {
		return // No valid token found
	}

	accountID := chi.URLParam(r, "accountid")

	account, err := controller.AccountService.GetAccount(accountID)

	if err.Code() != 0{
		http.Error(w, err.Error(), err.Code())
    	return
	}

	json.NewEncoder(w).Encode(account)
}

func (controller *AccountController) GetAccountType(w http.ResponseWriter, r *http.Request) {

	if !controller.ValidatorService.VerifyToken(r, w) {
		return // No valid token found
	}

	accountID := chi.URLParam(r, "accountid")

	account, err := controller.AccountService.GetAccountType(accountID)

	if err.Code() != 0{
		http.Error(w, err.Error(), err.Code())
    	return
	}

	json.NewEncoder(w).Encode(account)
}

func (controller *AccountController) GetAccountState(w http.ResponseWriter, r *http.Request) {

	if !controller.ValidatorService.VerifyToken(r, w) {
		return // No valid token found
	}

	accountID := chi.URLParam(r, "accountid")

	account, err := controller.AccountService.GetAccountState(accountID)

	if err.Code() != 0{
		http.Error(w, err.Error(), err.Code())
    	return
	}

	json.NewEncoder(w).Encode(account)
}

func (controller *AccountController) PostAccount(w http.ResponseWriter, r *http.Request) {

	if !controller.ValidatorService.VerifyToken(r, w) {
		return // No valid token found
	}

	var accountModel models.AccountModel
	json.NewDecoder(r.Body).Decode(&accountModel)

	account, err := controller.AccountService.PostAccount(accountModel)

	if err.Code() != 0{
		http.Error(w, err.Error(), err.Code())
    	return
	}

	json.NewEncoder(w).Encode(account)
}

func (controller *AccountController) PutAccount(w http.ResponseWriter, r *http.Request) {
	if !controller.ValidatorService.VerifyToken(r, w) {
		return // No valid token found
	}

	var accountModel models.AccountModel
	json.NewDecoder(r.Body).Decode(&accountModel)
	accountID := chi.URLParam(r, "accountid")

	account, err := controller.AccountService.PutAccount(accountID, accountModel)

	if err.Code() != 0{
		http.Error(w, err.Error(), err.Code())
    	return
	}

	json.NewEncoder(w).Encode(account)
}

func (controller *AccountController) PutAccountState(w http.ResponseWriter, r *http.Request) {
	if !controller.ValidatorService.VerifyToken(r, w) {
		return // No valid token found
	}

	var accountStateModel models.AccountStateModel
	json.NewDecoder(r.Body).Decode(&accountStateModel)
	accountID := chi.URLParam(r, "accountid")

	account, err := controller.AccountService.PutAccountState(accountID, accountStateModel)

	if err.Code() != 0{
		http.Error(w, err.Error(), err.Code())
    	return
	}

	json.NewEncoder(w).Encode(account)
}

func (controller *AccountController) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	if !controller.ValidatorService.VerifyToken(r, w) {
		return // No valid token found
	}
	
	accountID := chi.URLParam(r, "accountid")
	
	err := controller.AccountService.DeleteAccount(accountID)

	if err.Code() != 0{
		http.Error(w, err.Error(), err.Code())
    	return
	}

	
	json.NewEncoder(w).Encode("Deleted account : "+accountID)
}