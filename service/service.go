package service

import (
	"errors"
	"log"

	"github.com/atennyson/capstone_project/entities"
	"github.com/atennyson/capstone_project/repo"
)

type Repo interface {
	AddGame(g entities.Game) error
	ViewAll() (repo.DataBase, error)
	ViewSorted() (repo.DataBase, error)
	ViewUnFinished() (repo.DataBase, error)
	ViewUnPlayed() (repo.DataBase, error)
	ViewFinished() (repo.DataBase, error)
	FindByTitle(title string) (entities.Game, error)
	DeleteGame(title string) error
	UpdateGame(title string, game entities.Game) error
}

type Service struct {
	Repo Repo
}

func CreateService(r Repo) Service {
	return Service{
		Repo: r,
	}
}

func (s Service) AddGame(g entities.Game) error {
	if s.IterateData(g.Title) {
		return errors.New("game already exists")
	}
	err := s.Repo.AddGame(g)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) ViewAll() (repo.DataBase, error) {
	db, err := s.Repo.ViewAll()
	if err != nil {
		return db, err
	}

	return db, nil
}

func (s Service) ViewSorted() (repo.DataBase, error) {
	db, err := s.Repo.ViewSorted()
	if err != nil {
		return db, err
	}

	return db, nil
}

func (s Service) ViewUnFinished() (repo.DataBase, error) {
	db, err := s.Repo.ViewUnFinished()
	if err != nil {
		return db, err
	}

	return db, nil
}

func (s Service) ViewUnPlayed() (repo.DataBase, error) {
	db, err := s.Repo.ViewUnPlayed()
	if err != nil {
		return db, err
	}

	return db, nil
}

func (s Service) ViewFinished() (repo.DataBase, error) {
	db, err := s.Repo.ViewFinished()
	if err != nil {
		return db, err
	}

	return db, nil
}

func (s Service) FindByTitle(title string) (entities.Game, error) {
	if !s.IterateData(title) {
		return entities.Game{}, errors.New("game not found")
	}

	game, err := s.Repo.FindByTitle(title)
	if err != nil {
		return entities.Game{}, err
	}

	return game, nil
}

func (s Service) DeleteGame(title string) error {
	if !s.IterateData(title) {
		return errors.New("game not found")
	}

	err := s.Repo.DeleteGame(title)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) UpdateGame(title string, game entities.Game) error {
	if !s.IterateData(title) {
		return errors.New("game not found")
	}

	err := s.Repo.UpdateGame(title, game)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) IterateData(title string) bool {
	games, err := s.Repo.ViewAll()
	if err != nil {
		log.Print(err)
	}

	for _, game := range games.Games {
		if game.Title == title {
			return true
		}
	}

	return false
}
