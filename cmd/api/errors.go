package main

import (
	"fmt"
	"net/http"
)

// Cree log error
// Cree responseError
// Cree MethodNotAllowed
// Cree NotFound
// Cree ServerError

func (app *application) logError(r *http.Request, err error) {

	app.logger.Println(err)

}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request,
	status int, message interface{}) {

	env := envelope{"error": message}

	jsonObject, err := app.writeJSON(w, env, status, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
		return
	}

	w.Write(jsonObject)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {

	var message string = fmt.Sprintf("the %s method is not supported for this reouserce", r.Method)

	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {

	var message string = fmt.Sprintf("the requested resource %s was not found", r.URL)
	app.errorResponse(w, r, http.StatusNotFound, message)

}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	var message string = "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)

}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {

	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errs map[string]string) {

	app.errorResponse(w, r, http.StatusUnprocessableEntity, errs)
}
