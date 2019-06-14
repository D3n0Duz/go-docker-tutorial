package repositories

import (
	"context"
	"encoding/json"
	"fmt"

	"../interfaces"
	"../models"
	"github.com/fatih/structs"
)

const collection string = "accounts"

type AccountRepository struct {
	AccountDatabase interfaces.IAccountDatabase
}

func (repository *AccountRepository) GetAccount(accountid string) (models.AccountModel, models.Errors) {

	client := repository.AccountDatabase.GetClientConnection()

	ctx := context.Background()

	doc, err := client.Collection(collection).Doc(accountid).Get(ctx)

	if err != nil {
		return models.AccountModel{}, models.Errors{"failed to retrieved document: " + accountid, 400}
	}

	data := models.AccountModel{}

	jsonString, _ := json.Marshal(doc.Data())
	json.Unmarshal(jsonString, &data)

	return data, models.Errors{}
}

func (repository *AccountRepository) AddAccount(accountid string, accountModel models.AccountModel) (models.AccountModel, models.Errors) {
	
	client := repository.AccountDatabase.GetClientConnection()

	ctx := context.Background()

	accountToAdd := structs.Map(accountModel)

	_, err := client.Collection(collection).Doc(accountid).Set(ctx, accountToAdd)

	if err != nil {
		fmt.Println(err)
		return models.AccountModel{}, models.Errors{err.Error(), 400}
	}

	return accountModel, models.Errors{}
}

func (repository *AccountRepository) UpdateAccount(accountid string, accountModel models.AccountModel) (models.AccountModel, models.Errors) {
	
	client := repository.AccountDatabase.GetClientConnection()

	ctx := context.Background()

	accountToAdd := structs.Map(accountModel)

	_, err := client.Collection(collection).Doc(accountid).Set(ctx, accountToAdd)

	if err != nil {
		fmt.Println(err)
		return models.AccountModel{}, models.Errors{err.Error(), 400}
	}

	return accountModel, models.Errors{}
}

func (repository *AccountRepository) DeleteAccount(accountid string) models.Errors {
	
	client := repository.AccountDatabase.GetClientConnection()

	ctx := context.Background()

	_, err := client.Collection(collection).Doc(accountid).Delete(ctx)
	if err != nil {
		fmt.Println(err)
		return models.Errors{err.Error(), 400}
	}

	return models.Errors{}
}
