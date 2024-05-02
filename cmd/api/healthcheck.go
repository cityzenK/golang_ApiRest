package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandle(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"status":     "avaliable",
		"enviroment": app.config.env,
		"version":    version,
	}
	js := `{"status": "avaliable", "enviroment": %q, "version": %q}`
	js = fmt.Sprintf(js, app.config.env, version)

	w.Header().Set("Content-type", "application/json")

	w.Write([]byte(js))
}
