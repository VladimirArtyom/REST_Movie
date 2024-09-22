package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	
	var jsonEnvelope envelope = envelope{
		"status" : "available",
		"system_info": map[string]string {
			"env": app.config.env,
			"version": version,
		},
	}

	jsonObject, err := app.writeJSON(w, jsonEnvelope, http.StatusOK, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	
	w.Write(jsonObject)

}




