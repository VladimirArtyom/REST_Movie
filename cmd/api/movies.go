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
		app.notFoundResponse(w, r)
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

	var movieEnvelope envelope = envelope{
		"movie": movie,
	}

	
	jsonObject, err := app.writeJSON(w, movieEnvelope, http.StatusOK, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	w.Write(jsonObject)

}

