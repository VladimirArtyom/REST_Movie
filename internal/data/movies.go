package data

import "time"

type Movie struct {

	ID int64 `json:"id"` // Unique integer ID for the movie 
	Title string `json:"title"`  // The title of the movie
	Year int32 `json:"year"` // The year the movie was released
	Runtime int32 `json:"runtime"` // The runtime of the movie in minutes
	Genres []string  `json:"genres"`
	Version int32 `json:"version"` 
	CreatedAt time.Time `json:"created_at"` // Timestamp for when the movie was added to database.

}


