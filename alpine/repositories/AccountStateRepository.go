package repositories

import (
	"context"
	"encoding/json"
	"fmt"

	"../interfaces"
	"../models"
	"github.com/fatih/structs"
)

const collectionAccountState string = "accountStates"

type AccountStateRepository struct {
	AccountDatabase interfaces.IAccountDatabase
}

func (repository *AccountStateRepository) GetAccountState(accountStateid string) (models.AccountStateModel, models.Errors) {

	client := repository.AccountDatabase.GetClientConnection()

	ctx := context.Background()

	doc, err := client.Collection(collectionAccountState).Doc(accountStateid).Get(ctx)

	if err != nil {
		fmt.Println(err)
		return models.AccountStateModel{}, models.Errors{"failed to retrieved document: " + accountStateid, 400}
	}

	data := models.AccountStateModel{}

	jsonString, _ := json.Marshal(doc.Data())
	json.Unmarshal(jsonString, &data)

	return data, models.Errors{}
}

func (repository *AccountStateRepository) AddAccountState(accountid string, accountStateModel models.AccountStateModel) (models.AccountStateModel, models.Errors) {
	
	client := repository.AccountDatabase.GetClientConnection()

	ctx := context.Background()

	accountToAdd := structs.Map(accountStateModel)

	_, err := client.Collection(collectionAccountState).Doc(accountid).Set(ctx, accountToAdd)

	if err != nil {
		fmt.Println(err)
		return models.AccountStateModel{}, models.Errors{err.Error(), 400}
	}

	return accountStateModel, models.Errors{}
}

func (repository *AccountStateRepository) UpdateAccountState(accountStateId string, accountStateModel models.AccountStateModel) (models.AccountStateModel, models.Errors) {

	client := repository.AccountDatabase.GetClientConnection()

	ctx := context.Background()

	accountToAdd := structs.Map(accountStateModel)

	_, err := client.Collection(collectionAccountState).Doc(accountStateId).Set(ctx, accountToAdd)

	if err != nil {
		return models.AccountStateModel{}, models.Errors{err.Error(), 400}
	}

	return accountStateModel, models.Errors{}
}