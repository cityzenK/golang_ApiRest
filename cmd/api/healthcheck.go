package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) healthcheckHandle(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"status":     "avaliable",
		"enviroment": app.config.env,
		"version":    version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
