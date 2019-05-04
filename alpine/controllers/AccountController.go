package controllers

import (
	"fmt"
	"net/http"
	"../interfaces"
	"encoding/json"
	"github.com/go-chi/chi"
)

type AccountController struct {
	AccountService interfaces.IAccountService
	
}

func (controller *AccountController) GetAccount(w http.ResponseWriter, r *http.Request) {

	accountID := chi.URLParam(r, "accountid")

	fmt.Fprint(w, "GET "+accountID)

	account, err := controller.AccountService.GetAccount(accountID)

	if err != nil {
		//Handle error
	}

	response := account

	json.NewEncoder(w).Encode(response)
	/*
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
		}
	*/

	///////////////////////////////////

	/*scores, err := controller.PlayerService.GetScores(player1Name, player2Name)
	if err != nil {
		//Handle error
	}

	response := controller.PlayerHelper.BuildScoresVM(scores)

	json.NewEncoder(res).Encode(response)*/
}

func (controller *AccountController) PostAccount(w http.ResponseWriter, r *http.Request) {

	accountID := chi.URLParam(r, "accountid")
	fmt.Fprint(w, "POST "+accountID)

	/*scores, err := controller.PlayerService.GetScores(player1Name, player2Name)
	if err != nil {
		//Handle error
	}

	response := controller.PlayerHelper.BuildScoresVM(scores)

	json.NewEncoder(res).Encode(response)
	*/
}

func (controller *AccountController) PutAccount(w http.ResponseWriter, r *http.Request) {

	accountID := chi.URLParam(r, "accountid")
	fmt.Fprint(w, "PUT "+accountID)

	/*scores, err := controller.PlayerService.GetScores(player1Name, player2Name)
	if err != nil {
		//Handle error
	}

	response := controller.PlayerHelper.BuildScoresVM(scores)

	json.NewEncoder(res).Encode(response)
	*/
}

func (controller *AccountController) DeleteAccount(w http.ResponseWriter, r *http.Request) {

	accountID := chi.URLParam(r, "accountid")

	fmt.Fprint(w, "DELETE "+accountID)

	/*scores, err := controller.PlayerService.GetScores(player1Name, player2Name)
	if err != nil {
		//Handle error
	}

	response := controller.PlayerHelper.BuildScoresVM(scores)

	json.NewEncoder(res).Encode(response)
	*/
}
