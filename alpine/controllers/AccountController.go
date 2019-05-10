package controllers

import (
	"fmt"
	"net/http"
	"../interfaces"
	"encoding/json"
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

	accountID := chi.URLParam(r, "accountid")
	fmt.Fprint(w, "POST "+accountID)

	
}

func (controller *AccountController) PutAccount(w http.ResponseWriter, r *http.Request) {

	accountID := chi.URLParam(r, "accountid")
	fmt.Fprint(w, "PUT "+accountID)

	
}

func (controller *AccountController) DeleteAccount(w http.ResponseWriter, r *http.Request) {

	accountID := chi.URLParam(r, "accountid")

	fmt.Fprint(w, "DELETE "+accountID)

}
