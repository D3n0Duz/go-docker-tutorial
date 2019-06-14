package controllers

import (
	"fmt"
	"strings"
	"net/http"
	"../interfaces"
	"encoding/json"
	"../models"
	"github.com/go-chi/chi"
	"github.com/dgrijalva/jwt-go"
)

type AccountController struct {
	AccountService interfaces.IAccountService

}

const jwtKey = "YOUR_JWT_KEY" // TODO : get key from Vault

func (controller *AccountController) GetAccount(w http.ResponseWriter, r *http.Request) {

	if !VerifyToken(r, w) {
		return // No valid token found
	}

	accountID := chi.URLParam(r, "accountid")

	account, err := controller.AccountService.GetAccount(accountID)

	if err.Code() != 0{
		http.Error(w, err.Error(), err.Code())
    	return
	}

	response := account

	json.NewEncoder(w).Encode(response)
}

func (controller *AccountController) PostAccount(w http.ResponseWriter, r *http.Request) {

	if !VerifyToken(r, w) {
		return // No valid token found
	}

	var accountModel models.AccountModel
	json.NewDecoder(r.Body).Decode(&accountModel)

	account, err := controller.AccountService.PostAccount(accountModel)

	if err.Code() != 0{
		http.Error(w, err.Error(), err.Code())
    	return
	}

	response := account

	json.NewEncoder(w).Encode(response)
}

func (controller *AccountController) PutAccount(w http.ResponseWriter, r *http.Request) {
	if !VerifyToken(r, w) {
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

	response := account

	json.NewEncoder(w).Encode(response)
}

func (controller *AccountController) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	if !VerifyToken(r, w) {
		return // No valid token found
	}
	
	accountID := chi.URLParam(r, "accountid")
	
	err := controller.AccountService.DeleteAccount(accountID)

	if err.Code() != 0{
		http.Error(w, err.Error(), err.Code())
    	return
	}

	fmt.Fprint(w, "Deleted account : "+accountID)
}

// TODO create abstract class that can do that validation then go to the rest API calls
func VerifyToken(r *http.Request, w http.ResponseWriter) bool {

	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		http.Error(w, "Permission denied", 401)
    	return false
	}

	reqToken = strings.TrimSpace(splitToken[1])

    token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
        return []byte(jwtKey), nil
    })
    if err == nil && token.Valid {
        fmt.Println("valid token found")
        return true
    } else {
        http.Error(w, "Permission denied", 401)
        return false
    }
}
