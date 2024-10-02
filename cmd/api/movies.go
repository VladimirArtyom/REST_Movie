package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/VladimirArtyom/REST_Movie_API/internal/data"
	"github.com/VladimirArtyom/REST_Movie_API/internal/validator"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

	var tempOutput struct {
		Title   string       `json:"title"`
		Year    int32          `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	err := app.readJSON(w, r, &tempOutput)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}


	var v *validator.Validator = validator.New()
	var mov *data.Movie = &data.Movie{
		Title: tempOutput.Title,
		Year: tempOutput.Year,
		Runtime: tempOutput.Runtime,
		Genres: tempOutput.Genres,
	}

	data.ValidateMovie(v, mov)

	if !v.IsValid()  {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}


	fmt.Fprintf(w, "%+v\n", tempOutput)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var movie data.Movie = data.Movie{

		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Yes sir",
		Year:      201,
		Runtime:   100,
		Genres:    []string{"drama", "comedy", "jombie"},
		Version:   10,
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

