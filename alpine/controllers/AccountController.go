package controllers

import (
	"fmt"
	"net/http"
	"../interfaces"
	"encoding/json"
	"../models"
	"github.com/go-chi/chi"
)

type AccountController struct {
	AccountService interfaces.IAccountService
	
}

func (controller *AccountController) GetAccount(w http.ResponseWriter, r *http.Request) {

	accountID := chi.URLParam(r, "accountid")

	account, err := controller.AccountService.GetAccount(accountID)

	if err != nil {
		http.Error(w, err.Error(), 404)
    	return
	}

	response := account

	json.NewEncoder(w).Encode(response)
}

func (controller *AccountController) PostAccount(w http.ResponseWriter, r *http.Request) {

	var accountModel models.AccountModel
	json.NewDecoder(r.Body).Decode(&accountModel)

	account, err := controller.AccountService.PostAccount(accountModel)

	if err != nil {
		http.Error(w, err.Error(), 404)
    	return
	}

	response := account

	json.NewEncoder(w).Encode(response)
}

func (controller *AccountController) PutAccount(w http.ResponseWriter, r *http.Request) {

	var accountModel models.AccountModel
	json.NewDecoder(r.Body).Decode(&accountModel)
	accountID := chi.URLParam(r, "accountid")

	account, err := controller.AccountService.PutAccount(accountID, accountModel)

	if err != nil {
		http.Error(w, err.Error(), 404)
    	return
	}

	response := account

	json.NewEncoder(w).Encode(response)
}

func (controller *AccountController) DeleteAccount(w http.ResponseWriter, r *http.Request) {

	accountID := chi.URLParam(r, "accountid")

	fmt.Fprint(w, "DELETE "+accountID)

}
