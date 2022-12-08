package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	handlers "github.com/atennyson/capstone_project/handlers"
	"github.com/atennyson/capstone_project/repo"
	"github.com/atennyson/capstone_project/service"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	var err error

	host := os.Getenv("HOST")
	po := os.Getenv("DBPORT")
	user := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	port, _ := strconv.Atoi(po)

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	DB, err = sql.Open("postgres", sqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
}

func main() {

	repository := repo.NewRepo(DB)
	serv := service.CreateService(repository)
	handle := handlers.NewGameHandler(serv)
	router := handlers.ConfigureRouter(handle)

	server := &http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	log.Fatal(server.ListenAndServe())
}
