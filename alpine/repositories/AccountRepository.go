package repositories

import (
	"context"
	"encoding/json"
	"fmt"

	"../interfaces"
	"../models"
)

type AccountRepository struct {
	AccountDatabase interfaces.IAccountDatabase
}

func (repository *AccountRepository) GetAccount(accountid string) (models.AccountModel, error) {

	client := repository.AccountDatabase.GetClientConnection()

	ctx := context.Background()

	doc, err := client.Collection("accounts").Doc(accountid).Get(ctx)

	if err != nil {
		fmt.Println(doc.Data())
		return models.AccountModel{}, nil
	}

	data := models.AccountModel{}

	jsonString, _ := json.Marshal(doc.Data())
	json.Unmarshal(jsonString, &data)

	return data, nil
}

func (repository *AccountRepository) PostAccount(accountid string) (models.AccountModel, error) {
	return models.AccountModel{}, nil
}

func (repository *AccountRepository) PutAccount(accountid string) (models.AccountModel, error) {
	return models.AccountModel{}, nil
}

func (repository *AccountRepository) DeleteAccount(accountid string) (models.AccountModel, error) {
	return models.AccountModel{}, nil
}
