package repositories

import (
	"fmt"
	"simple-crud/models"
	"simple-crud/services"
)

func CreateMovie(movie models.Movie) (int64, error) {
	conn, err := services.OpenDbConnection()
	if err != nil {
		fmt.Println("Falha em abrir db")
		return 0, err
	}
	defer conn.Close()
	query := `INSERT INTO movies (title, genre, year) VALUES (?,?,?)`
	result, err := conn.Exec(query, movie.Title, movie.Genre, movie.Year)
	if err != nil {
		fmt.Println("Impossivel cadastrar movie")
		return 0, err
	}
	return result.LastInsertId()
}

func GetMovieById(id int) (models.Movie, error) {
	conn, err := services.OpenDbConnection()
	if err != nil {
		fmt.Println("Falha em conectar ao db")
		return models.Movie{}, err
	}
	defer conn.Close()
	var movie models.Movie

	err = conn.QueryRow(`SELECT * FROM movies WHERE id=?`, id).Scan(
		&movie.Id, &movie.Title, &movie.Genre, &movie.Year)
	if err != nil {
		fmt.Println("Movie not found")
		return movie, err
	}
	return movie, err
}

func GetAll() ([]models.Movie, error) {
	conn, err := services.OpenDbConnection()
	if err != nil {
		fmt.Println("Impossivel acessar o db")
		return nil, err
	}
	defer conn.Close()
	var movies []models.Movie

	result, err := conn.Query(`SELECT * FROM movies`)
	for result.Next() {
		var movie models.Movie
		err = result.Scan(&movie.Id, &movie.Title, &movie.Genre, &movie.Year)
		if err != nil {
			continue
		}
		movies = append(movies, movie)
	}
	return movies, err
}

func UpdateMovie(movie models.Movie, id int) (int64, error) {
	conn, err := services.OpenDbConnection()
	if err != nil {
		fmt.Println("Impossive acessar o db")
		return 0, err
	}

	defer conn.Close()

	query := `UPDATE movies SET title=?, genre=?, year=? WHERE id=?`
	result, err := conn.Exec(query, movie.Title, movie.Genre, movie.Year, id)
	if err != nil {
		fmt.Println("Impossivel atualizar o arquivo")
		return 0, err
	}

	return result.RowsAffected()
}

func DeleteMovie(id int) (int64, error) {
	conn, err := services.OpenDbConnection()
	if err != nil {
		fmt.Println("Impossivel acessar o db")
		return 0, err
	}
	defer conn.Close()

	result, err := conn.Exec(`DELETE from movies WHERE id=?`, id)
	if err != nil {
		fmt.Println("Impossivel deletar o arquivo")
		return 0, err
	}

	return result.RowsAffected()
}
