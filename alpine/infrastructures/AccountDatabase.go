package infrastructures

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	firestore "cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type AccountDatabase struct {
	client firestore.Client
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

	// Uncomment will close the connection at the end of the method. We want to avoid this for now...
	//defer client.Close()

	nosql.client = *client
	return err
}

func (nosql *AccountDatabase) GetClientConnection() firestore.Client {
	return nosql.client
}
