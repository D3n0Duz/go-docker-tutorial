package infrastructures

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	auth "firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type AccountDatabase struct {
	client *auth.Client
}

func (nosql *AccountDatabase) InitDatabase() error {
	var err error

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

	return err
}

func (nosql *AccountDatabase) GetClientConnection() *auth.Client {
	nosql.client.Collection("accounts").Doc(accountid).Get(ctx)
	return nosql.client
}
