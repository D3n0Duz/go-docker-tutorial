package services

import (
	"net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"regexp"
	"../models"
)

type ValidatorService struct{
}

const jwtKey = "YOUR_JWT_KEY" // TODO : get key from Vault
const emailRegex = "^[a-zA-Z0-9_+&*-]+(?:\\.[a-zA-Z0-9_+&*-]+)*@(?:[a-zA-Z0-9-]+\\.)+[a-zA-Z]{2,7}$"
const dateRegex = "^[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}Z$"
const typeIdRegex = "^[a-zA-Z0-9]{20}$"
const stateIdRegex = "^[a-zA-Z0-9]{20}$"

func (service *ValidatorService) VerifyToken(r *http.Request, w http.ResponseWriter) bool {

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

func (service *ValidatorService) IsFieldValid(accountModel models.AccountModel) bool{
	emailMatch, _ := regexp.MatchString(emailRegex, accountModel.Email)
	accountStateIdMatch, _ := regexp.MatchString(stateIdRegex, accountModel.AccountStateId)
	accountTypeIdMatch, _ := regexp.MatchString(typeIdRegex, accountModel.AccountTypeId)
	return emailMatch && accountStateIdMatch && accountTypeIdMatch
}


func (service *ValidatorService) IsValid(accoundModelFromDB models.AccountModel, accountModel models.AccountModel) models.Errors{

	if accoundModelFromDB.Email != accountModel.Email{
		errMessage := "Email does not match"
		fmt.Println(errMessage)
		return models.Errors{errMessage, 400}
	}

	return models.Errors{}
}