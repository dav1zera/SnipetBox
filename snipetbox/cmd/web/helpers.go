package main

import (
	"fmt"
	"net/http"
  "runtime/debug"
)

func (app *applicationAux) serverError(w http.ResponseWriter, err error) {
  trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

  app.errorsLog.Output(2, trace)
  http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *applicationAux) clientError(w http.ResponseWriter, status int) {
  http.Error(w, http.StatusText(status), status)
}

func (app *applicationAux) notFound (w http.ResponseWriter) {
  app.clientError(w, http.StatusNotFound)
}