package controllers

import (
	"fmt"
	"net/http"

	//"github.com/irahardianto/service-pattern-go/helpers"
	//"github.com/irahardianto/service-pattern-go/interfaces"

	"github.com/go-chi/chi"
)

type UserController struct {
}

func (controller *UserController) GetUser(w http.ResponseWriter, r *http.Request) {

	userID := chi.URLParam(r, "userid")

	fmt.Fprint(w, "GET "+userID)

	/*scores, err := controller.PlayerService.GetScores(player1Name, player2Name)
	if err != nil {
		//Handle error
	}

	response := controller.PlayerHelper.BuildScoresVM(scores)

	json.NewEncoder(res).Encode(response)*/
}

func (controller *UserController) PostUser(w http.ResponseWriter, r *http.Request) {

	userID := chi.URLParam(r, "userid")
	fmt.Fprint(w, "POST "+userID)

	/*scores, err := controller.PlayerService.GetScores(player1Name, player2Name)
	if err != nil {
		//Handle error
	}

	response := controller.PlayerHelper.BuildScoresVM(scores)

	json.NewEncoder(res).Encode(response)
	*/
}

func (controller *UserController) PutUser(w http.ResponseWriter, r *http.Request) {

	userID := chi.URLParam(r, "userid")
	fmt.Fprint(w, "PUT "+userID)

	/*scores, err := controller.PlayerService.GetScores(player1Name, player2Name)
	if err != nil {
		//Handle error
	}

	response := controller.PlayerHelper.BuildScoresVM(scores)

	json.NewEncoder(res).Encode(response)
	*/
}

func (controller *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {

	userID := chi.URLParam(r, "userid")

	fmt.Fprint(w, "DELETE "+userID)

	/*scores, err := controller.PlayerService.GetScores(player1Name, player2Name)
	if err != nil {
		//Handle error
	}

	response := controller.PlayerHelper.BuildScoresVM(scores)

	json.NewEncoder(res).Encode(response)
	*/
}
