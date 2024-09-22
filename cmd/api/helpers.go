package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type envelope map[string]interface{}

func (app *application) readIDParam(r *http.Request) (int64, error){

	var params httprouter.Params = httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt( params.ByName("id"), 10, 64)

	if err != nil {
		return 0, errors.New("Invalid id parameter")
	}

	return id, nil
}

func (app * application) writeHeaders(w http.ResponseWriter,
																		  status int, headers http.Header){

	for key, value := range headers {
	
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
} 

func (app *application) writeJSON(w http.ResponseWriter, data envelope,
																	status int, headers http.Header) ([]byte, error){

	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return nil, err
	}

	jsonData = append(jsonData, '\n')

	app.writeHeaders(w, status, headers)

	return jsonData, nil

} 
