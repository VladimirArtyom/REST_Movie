package data

import (
	"time"

	"github.com/VladimirArtyom/REST_Movie_API/internal/validator"
)

type Movie struct {

	ID int64 `json:"id"` // Unique integer ID for the movie 
	Title string `json:"title"`  // The title of the movie
	Year int32 `json:"year,omitempty"` // The year the movie was released
	Runtime Runtime `json:"runtime,omitempty"` // The runtime of the movie in minutes
	Genres []string  `json:"genres,omitempty"`
	Version int32 `json:"version"` 
	CreatedAt time.Time `json:"-"` // Timestamp for when the movie was added to database.

}

func ValidateMovie(v *validator.Validator, movie *Movie) {

	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must be lower than 500 length")

	v.Check(movie.Year > 1888, "year", "must be greater than 1888")
	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not have more than 5 genres")
	// Genres need to be unique
	v.Check(v.IsUnique(movie.Genres), "genres", "genres must be unique")


}
