package main

import (
	"fmt"
	"simple-crud/models"
	"simple-crud/repositories"
)

func main() {
	deleteMovie()
}
func createMovie() {
	movie := models.Movie{
		Title: "Lord of the rings",
		Genre: "Fantasy",
		Year:  "2010",
	}
	repositories.CreateMovie(movie)
}
func getMovieById() {
	movie, _ := repositories.GetMovieById(1)
	fmt.Println(movie)
}
func getMovies() {
	movies, _ := repositories.GetAll()
	for _, movie := range movies {
		fmt.Println(movie)
	}
}
func updateMovie() {
	movie := models.Movie{
		Title: "Matrix II",
		Genre: "Ficçao",
		Year:  "2003"}

	repositories.UpdateMovie(movie, 2)
}
func deleteMovie() {
	repositories.DeleteMovie(3)
}
