package repositories

import (
	"../models"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"context"
	"log"
	"encoding/json"
	"fmt"
)

type AccountRepository struct {
}

func (repository *AccountRepository) GetAccount(accountid string) (models.AccountModel, error){
	
	sa := option.WithCredentialsFile("./serviceAccountKey.json")
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

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

func (repository *AccountRepository) PostAccount(accountid string) (models.AccountModel, error){
	return models.AccountModel{}, nil
}

func (repository *AccountRepository) PutAccount(accountid string) (models.AccountModel, error){
	return models.AccountModel{}, nil
}

func (repository *AccountRepository) DeleteAccount(accountid string) (models.AccountModel, error){
	return models.AccountModel{}, nil
}