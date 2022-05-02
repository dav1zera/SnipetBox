package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
 
)

func(app *applicationAux) home(responseWriter http.ResponseWriter, request *http.Request) {
  if request.URL.Path != "/" {
   app.notFound(responseWriter)
    return
     
  }
  
    files := []string{
      "./ui/html/home.page.tmpl.html",
      "./ui/html/base.layout.tmpl.html",
      "./ui/html/footer.partial.tmpl.html",
      
    }

    ts, err := template.ParseFiles(files...)
    if err != nil {
     app.serverError(responseWriter, err)
      return
    }

    err = ts.Execute(responseWriter, nil) 
    if err != nil {
     app.serverError(responseWriter, err)
      return
      }
      
}

func(app *applicationAux) showSnippet (responseWriter http.ResponseWriter, request *http.Request)  {
  id, err := strconv.Atoi(request.URL.Query().Get("id"))

  if err != nil || id < 1  {
    app.notFound(responseWriter)
    return
  }
  fmt.Fprintf(responseWriter, "Exibir o snippet de ID: %d", id)
}

func(app *applicationAux) createSnippet (responseWriter http.ResponseWriter, request *http.Request)  {
  
   if request.Method != "POST" {
    responseWriter.Header().Set("Allow", "POST")
    responseWriter.WriteHeader(405)
     http.Error(responseWriter, "Não permitido", http.StatusMethodNotAllowed)
    app.clientError(responseWriter, http.StatusMethodNotAllowed)
    return 
   }
  responseWriter.Write([]byte ("Criando novo snippet"))
} 