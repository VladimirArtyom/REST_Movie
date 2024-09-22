package data

import "time"

type Movie struct {

	ID int64 `json:"id"` // Unique integer ID for the movie 
	Title string `json:"title"`  // The title of the movie
	Year int32 `json:"year,omitempty"` // The year the movie was released
	Runtime Runtime `json:"runtime,omitempty"` // The runtime of the movie in minutes
	Genres []string  `json:"genres,omitempty"`
	Version int32 `json:"version"` 
	CreatedAt time.Time `json:"-"` // Timestamp for when the movie was added to database.

}

