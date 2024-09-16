package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	

	var jsonMap map[string]string = map[string]string{
		"status": "available",
		"env": app.config.env,
		"version": version,
	}

	jsonObject, err := app.writeJSON(w, jsonMap, http.StatusOK, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Write(jsonObject)

}




