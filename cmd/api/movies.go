package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/VladimirArtyom/REST_Movie_API/internal/data"
)


func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "creating a movie")

}


func (app *application) showMovieHandler(w http.ResponseWriter, r* http.Request) {
	
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	var movie data.Movie = data.Movie{

		ID: id,
		CreatedAt: time.Now(),
		Title: "Yes sir",
		Year: 201,
		Runtime: 100,
		Genres: []string{"drama", "comedy", "jombie"},
		Version: 10,
	}

	
	jsonObject, err := app.writeJSON(w, movie, http.StatusOK, nil)

	if err != nil {
		app.logger.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonObject)

}

