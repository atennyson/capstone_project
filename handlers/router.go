package handlers

import (
	"github.com/gorilla/mux"
)

func ConfigureRouter(handler GameHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/games", handler.GetAllGames).Methods("GET")
	r.HandleFunc("/games/sorted", handler.GetSortedGames).Methods("GET")
	r.HandleFunc("/games/unplayed", handler.GetUnplayedGames).Methods("GET")
	r.HandleFunc("/games/started/unfinished", handler.GetUnfinishedGames).Methods("GET")
	r.HandleFunc("/games/finished", handler.GetFinishedGames).Methods("GET")
	r.HandleFunc("/games/{title}", handler.GetSpecific).Methods("GET")
	r.HandleFunc("/games/newgame", handler.AddNewGame).Methods("POST")
	r.HandleFunc("/games/{title}", handler.UpdateGam).Methods("PUT")
	r.HandleFunc("/games/{title}", handler.DeleteGam).Methods("DELETE")

	return r
}
