package repo

import (
	"database/sql"

	"github.com/atennyson/capstone_project/entities"
)

type DataBase struct {
	Games []entities.Game
}

type Repo struct {
	DB *sql.DB
}

func NewRepo(d *sql.DB) Repo {
	return Repo{
		DB: d,
	}
}

func (r Repo) AddGame(g entities.Game) error {
	_, err := r.DB.Exec("INSERT INTO games (title, developer, started, finished) VALUES ($1, $2, $3, $4)", g.Title, g.Developer, g.Started, g.Finished)
	if err != nil {
		return err
	}

	return nil
}

func (r Repo) ViewAll() (DataBase, error) {
	rows, err := r.DB.Query("SELECT * FROM games")
	if err != nil {
		return DataBase{}, err
	}
	defer rows.Close()

	games := DataBase{}
	for rows.Next() {
		game := entities.Game{}
		err := rows.Scan(&game.ID, &game.Title, &game.Developer, &game.Started, &game.Finished)
		if err != nil {
			return DataBase{}, err
		}

		games.Games = append(games.Games, game)
	}

	return games, nil
}

func (r Repo) ViewSorted() (DataBase, error) {
	rows, err := r.DB.Query("SELECT * FROM games ORDER BY title ASC")
	if err != nil {
		return DataBase{}, err
	}
	defer rows.Close()

	games := DataBase{}
	for rows.Next() {
		game := entities.Game{}
		err := rows.Scan(&game.ID, &game.Title, &game.Developer, &game.Started, &game.Finished)
		if err != nil {
			return DataBase{}, err
		}

		games.Games = append(games.Games, game)
	}

	return games, nil
}

func (r Repo) ViewUnPlayed() (DataBase, error) {
	rows, err := r.DB.Query("SELECT * FROM games WHERE started=false")
	if err != nil {
		return DataBase{}, err
	}
	defer rows.Close()

	games := DataBase{}
	for rows.Next() {
		game := entities.Game{}
		err := rows.Scan(&game.ID, &game.Title, &game.Developer, &game.Started, &game.Finished)
		if err != nil {
			return DataBase{}, err
		}

		games.Games = append(games.Games, game)
	}

	return games, nil
}

func (r Repo) ViewUnFinished() (DataBase, error) {
	rows, err := r.DB.Query("SELECT * FROM games WHERE started=true AND finished=false")
	if err != nil {
		return DataBase{}, err
	}
	defer rows.Close()

	games := DataBase{}
	for rows.Next() {
		game := entities.Game{}
		err := rows.Scan(&game.ID, &game.Title, &game.Developer, &game.Started, &game.Finished)
		if err != nil {
			return DataBase{}, err
		}

		games.Games = append(games.Games, game)
	}

	return games, nil
}

func (r Repo) ViewFinished() (DataBase, error) {
	rows, err := r.DB.Query("SELECT * FROM games WHERE finished=true")
	if err != nil {
		return DataBase{}, err
	}
	defer rows.Close()

	games := DataBase{}
	for rows.Next() {
		game := entities.Game{}
		err := rows.Scan(&game.ID, &game.Title, &game.Developer, &game.Started, &game.Finished)
		if err != nil {
			return DataBase{}, err
		}

		games.Games = append(games.Games, game)
	}

	return games, nil
}

func (r Repo) FindByTitle(title string) (entities.Game, error) {
	rows, err := r.DB.Query("SELECT * FROM games WHERE title=$1", title)
	if err != nil {
		return entities.Game{}, err
	}
	defer rows.Close()
	game := entities.Game{}
	for rows.Next() {
		err := rows.Scan(&game.ID, &game.Title, &game.Developer, &game.Started, &game.Finished)
		if err != nil {
			return entities.Game{}, err
		}
	}

	return game, nil
}

func (r Repo) DeleteGame(title string) error {
	_, err := r.DB.Exec("DELETE FROM games WHERE title=$1", title)
	if err != nil {
		return err
	}

	return nil
}

func (r Repo) UpdateGame(title string, game entities.Game) error {
	_, err := r.DB.Exec("UPDATE games SET title=$2, developer=$3, started=$4, finished=$5 WHERE title=$1", title, game.Title, game.Developer, game.Started, game.Finished)
	if err != nil {
		return err
	}

	return nil
}
