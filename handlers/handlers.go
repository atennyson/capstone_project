package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/atennyson/capstone_project/entities"
	"github.com/atennyson/capstone_project/repo"
	"github.com/gorilla/mux"
)

type Service interface {
	AddGame(g entities.Game) error
	ViewAll() (repo.DataBase, error)
	ViewSorted() (repo.DataBase, error)
	ViewUnFinished() (repo.DataBase, error)
	ViewUnPlayed() (repo.DataBase, error)
	FindByTitle(title string) (entities.Game, error)
	DeleteGame(title string) error
	UpdateGame(title string, g entities.Game) error
}

type GameHandler struct {
	Serv Service
}

func NewGameHandler(s Service) GameHandler {
	return GameHandler{
		Serv: s,
	}
}

func (gme GameHandler) AddNewGame(w http.ResponseWriter, r *http.Request) {
	gm := entities.Game{}

	err := json.NewDecoder(r.Body).Decode(&gm)
	if err != nil {
		w.WriteHeader(400)
	}

	err = gme.Serv.AddGame(gm)
	if err != nil {
		switch err.Error() {
		case "game already exists":
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
}

func (gme GameHandler) GetAllGames(w http.ResponseWriter, r *http.Request) {
	gmeDB, err := gme.Serv.ViewAll()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	for _, game := range gmeDB.Games {
		fmt.Fprintf(w, "ID: %d Title: %s Developer: %s Started?: %t Finished?: %t\n", game.ID, game.Title, game.Developer, game.Started, game.Finished)
	}

	w.WriteHeader(200)
}

func (gme GameHandler) GetSortedGames(w http.ResponseWriter, r *http.Request) {
	gmeDB, err := gme.Serv.ViewAll()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	for _, game := range gmeDB.Games {
		fmt.Fprintf(w, "ID: %d Title: %s Developer: %s Started?: %t Finished?: %t\n", game.ID, game.Title, game.Developer, game.Started, game.Finished)
	}

	w.WriteHeader(200)
}

func (gme GameHandler) GetUnfinishedGames(w http.ResponseWriter, r *http.Request) {
	gmeDB, err := gme.Serv.ViewAll()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	for _, game := range gmeDB.Games {
		fmt.Fprintf(w, "ID: %d Title: %s Developer: %s Started?: %t Finished?: %t\n", game.ID, game.Title, game.Developer, game.Started, game.Finished)
	}

	w.WriteHeader(200)
}

func (gme GameHandler) GetUnplayedGames(w http.ResponseWriter, r *http.Request) {
	gmeDB, err := gme.Serv.ViewAll()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	for _, game := range gmeDB.Games {
		fmt.Fprintf(w, "ID: %d Title: %s Developer: %s Started?: %t Finished?: %t\n", game.ID, game.Title, game.Developer, game.Started, game.Finished)
	}

	w.WriteHeader(200)
}

func (gme GameHandler) GetFinishedGames(w http.ResponseWriter, r *http.Request) {
	gmeDB, err := gme.Serv.ViewAll()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	for _, game := range gmeDB.Games {
		fmt.Fprintf(w, "ID: %d Title: %s Developer: %s Started?: %t Finished?: %t\n", game.ID, game.Title, game.Developer, game.Started, game.Finished)
	}

	w.WriteHeader(200)
}

func (gme GameHandler) GetSpecific(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	game, err := gme.Serv.FindByTitle(title)
	if err != nil {
		switch err.Error() {
		case "game not found":
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}

	fmt.Fprintf(w, "ID: %d Title: %s Developer: %s Started?: %t Finished?: %t\n", game.ID, game.Title, game.Developer, game.Started, game.Finished)
	w.WriteHeader(200)
}

func (gme GameHandler) DeleteGam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	err := gme.Serv.DeleteGame(title)
	if err != nil {
		switch err.Error() {
		case "game not found":
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}

func (gme GameHandler) UpdateGam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	game := entities.Game{}

	err := json.NewDecoder(r.Body).Decode(&game)
	if err != nil {
		w.WriteHeader(400)
	}

	err = gme.Serv.UpdateGame(title, game)
	if err != nil {
		w.WriteHeader(400)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
}
