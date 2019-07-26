package repositories

import (
	"context"
	"encoding/json"
	"fmt"

	"../interfaces"
	"../models"
	"github.com/fatih/structs"
)

const collectionAccountTypes string = "accountTypes"

type AccountTypeRepository struct {
	AccountDatabase interfaces.IAccountDatabase
}

func (repository *AccountTypeRepository) GetAccountType(accountTypeid string) (models.AccountTypeModel, models.Errors) {

	client := repository.AccountDatabase.GetClientConnection()

	ctx := context.Background()

	doc, err := client.Collection(collectionAccountTypes).Doc(accountTypeid).Get(ctx)

	if err != nil {
		fmt.Println(err)
		return models.AccountTypeModel{}, models.Errors{"failed to retrieved document: " + accountTypeid, 400}
	}

	data := models.AccountTypeModel{}

	jsonString, _ := json.Marshal(doc.Data())
	json.Unmarshal(jsonString, &data)

	return data, models.Errors{}
}

func (repository *AccountTypeRepository) AddAccountType(accountid string, accountTypeModel models.AccountTypeModel) (models.AccountTypeModel, models.Errors) {
	
	client := repository.AccountDatabase.GetClientConnection()

	ctx := context.Background()

	accountToAdd := structs.Map(accountTypeModel)

	_, err := client.Collection(collectionAccountTypes).Doc(accountid).Set(ctx, accountToAdd)

	if err != nil {
		fmt.Println(err)
		return models.AccountTypeModel{}, models.Errors{err.Error(), 400}
	}

	return accountTypeModel, models.Errors{}
}

func (repository *AccountTypeRepository) UpdateAccount(accountid string, accountTypeModel models.AccountTypeModel) (models.AccountTypeModel, models.Errors) {
	// TODO LATER

	// client := repository.AccountDatabase.GetClientConnection()

	// ctx := context.Background()

	// accountToAdd := structs.Map(accountTypeModel)

	// _, err := client.Collection(collectionAccountTypes).Doc(accountid).Set(ctx, accountToAdd)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return models.AccountTypeModel{}, models.Errors{err.Error(), 400}
	// }

	// return accountTypeModel, models.Errors{}
	return models.AccountTypeModel{}, models.Errors{}
}