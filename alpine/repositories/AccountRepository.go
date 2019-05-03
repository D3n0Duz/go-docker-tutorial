package repositories

import (
	"../models"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"context"
	"log"
	
	"google.golang.org/api/iterator"
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

	iter := client.Collection("accounts").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)

		}

		fmt.Println(doc.Data())
		var data models.AccountModel
		//json.Unmarshal([]byte(doc.Data()[0]), &data)
		//data, _ := json.Marshal(doc.Data())

		return data, nil
	}
	return models.AccountModel{}, err
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