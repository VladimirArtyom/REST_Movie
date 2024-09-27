package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

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

func (app *application) readJSON(w http.ResponseWriter,
																 r *http.Request, target interface{}) error {

	var maxBytes int = 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	
	var jsonDecoder *json.Decoder = json.NewDecoder(r.Body)
	jsonDecoder.DisallowUnknownFields() // Making sure if there is an unknown fields, return errors when decoding
	err := jsonDecoder.Decode(&target)

	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshallError *json.InvalidUnmarshalError 

		switch {

		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")
		
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {

				return fmt.Errorf("body contains incorrect JSON type for field (%s)", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		case strings.HasPrefix(err.Error(), "json: unknown field"):
			var fieldName string = strings.TrimPrefix(err.Error(), "json: unknown field")
			return fmt.Errorf("body contains uknown key %s", fieldName)

		case err.Error() == "http: request body too large":
			return fmt.Errorf("body must not be larger than %d bytes", maxBytes)

		case errors.As(err, &invalidUnmarshallError):
			panic(err)

		default:
			fmt.Println("This error is not handled properly")
			return err
		}


	}
	err = jsonDecoder.Decode(&struct{}{})
	if err != io.EOF {
		// If the json is not empty, then there must be something 
		return errors.New("body must only contain a single JSON value")
	}
	return nil
		
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


