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

func (repository *AccountRepository) GetAccount(accountid string) (models.AccountModel, error) {

	client := repository.AccountDatabase.GetClientConnection()

	ctx := context.Background()

	doc, err := client.Collection(collection).Doc(accountid).Get(ctx)

	if err != nil {
		fmt.Println(err)
		return models.AccountModel{}, fmt.Errorf("failed to retrieved docId : %s", accountid)
	}

	data := models.AccountModel{}

	jsonString, _ := json.Marshal(doc.Data())
	json.Unmarshal(jsonString, &data)

	return data, nil
}

func (repository *AccountRepository) CreateOrUpdateAccount(accountid string, accountModel models.AccountModel) (models.AccountModel, error) {
	
	client := repository.AccountDatabase.GetClientConnection()

	ctx := context.Background()

	accountToAdd := structs.Map(accountModel)

	_, err := client.Collection(collection).Doc(accountid).Set(ctx, accountToAdd)

	if err != nil {
		fmt.Println(err)
		return models.AccountModel{}, err
	}

	return accountModel, nil
}

func (repository *AccountRepository) DeleteAccount(accountid string) error {
	
	client := repository.AccountDatabase.GetClientConnection()

	ctx := context.Background()

	_, err := client.Collection(collection).Doc(accountid).Delete(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
