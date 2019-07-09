package interfaces

import (
	"net/http"
	"../models"
)

type IValidatorService interface {
	VerifyToken(r *http.Request, w http.ResponseWriter) bool 
	IsValid(accoundModelFromDB models.AccountModel, accountModel models.AccountModel) models.Errors
	IsFieldValid(accountModel models.AccountModel) bool
}
