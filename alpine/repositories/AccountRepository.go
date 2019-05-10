package repositories

import (
	"context"
	"encoding/json"
	"fmt"

	"../interfaces"
	"../models"
	"github.com/fatih/structs"
)

type AccountRepository struct {
	AccountDatabase interfaces.IAccountDatabase
}

func (repository *AccountRepository) GetAccount(accountid string) (models.AccountModel, error) {

	client := repository.AccountDatabase.GetClientConnection()

	ctx := context.Background()

	doc, err := client.Collection("accounts").Doc(accountid).Get(ctx)

	if err != nil {
		fmt.Println(err)
		return models.AccountModel{}, fmt.Errorf("failed to retrieved docId : %s", accountid)
	}

	data := models.AccountModel{}

	jsonString, _ := json.Marshal(doc.Data())
	json.Unmarshal(jsonString, &data)

	return data, nil
}

func (repository *AccountRepository) PostAccount(account string) (models.AccountModel, error) {

	// client := repository.AccountDatabase.GetClientConnection()

	// ctx := context.Background()

	// accountModel := models.AccountModel{}
	// fmt.Println(account)
	// jsonString, _ := json.Marshal(account)
	// json.Unmarshal(jsonString, &accountModel)

	// accountToAdd := structs.Map(accountModel)
	// fmt.Println(accountToAdd)

	// _, _, err := client.Collection("accounts").Add(ctx, accountToAdd)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return models.AccountModel{}, err
	// }

	return models.AccountModel{}, nil
}

func (repository *AccountRepository) PutAccount(accountid string) (models.AccountModel, error) {
	return models.AccountModel{}, nil
}

func (repository *AccountRepository) DeleteAccount(accountid string) (models.AccountModel, error) {
	return models.AccountModel{}, nil
}
