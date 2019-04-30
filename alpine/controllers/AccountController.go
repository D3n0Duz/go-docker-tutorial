package controllers

import (
	"fmt"
	"net/http"

	//"github.com/irahardianto/service-pattern-go/helpers"
	//"github.com/irahardianto/service-pattern-go/interfaces"

	"github.com/go-chi/chi"
)

type AccountController struct {
}

func (controller *AccountController) GetAccount(w http.ResponseWriter, r *http.Request) {

	accountID := chi.URLParam(r, "accountid")

	fmt.Fprint(w, "GET "+accountID)

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
