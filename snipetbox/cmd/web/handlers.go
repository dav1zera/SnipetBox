package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
  "log"
)

//ListAndServer = async
//Response e Request

/*

  curl -i -X POST          http://localhost:4000/snippet/create
  curl -i -X GET           http://localhost:4000/snippet/create

*/

//http://localhost:4000/snippet?id=123

func home(responseWriter http.ResponseWriter, request *http.Request) {
  if request.URL.Path != "/" {
    http.NotFound(responseWriter, request) 
    return
     
  }
  
    files := []string{
      "./ui/html/home.page.tmpl.html",
      "./ui/html/base.layout.tmpl.html",
      "./ui/html/footer.partial.tmpl.html",
      
    }

    ts, err := template.ParseFiles(files...)
    if err != nil {
      log.Println(err.Error())
      http.Error(responseWriter, "Internal Error", 500)
      return
    }

    err = ts.Execute(responseWriter, nil) 
    if err != nil {
      log.Println(err.Error())
      http.Error(responseWriter, "Internal Error", 500) 
      return
      }
      
}

func showSnippet (responseWriter http.ResponseWriter, request *http.Request)  {
  id, err := strconv.Atoi(request.URL.Query().Get("id"))

  if err != nil || id < 1  {
    http.NotFound(responseWriter, request)
    return
  }
  fmt.Fprintf(responseWriter, "Exibir o snippet de ID: %d", id)
}

func createSnippet (responseWriter http.ResponseWriter, request *http.Request)  {
  
   if request.Method != "POST" {
    responseWriter.Header().Set("Allow", "POST")
    responseWriter.WriteHeader(405)
     http.Error(responseWriter, "Não permitido", http.StatusMethodNotAllowed)
    
    return 
   }
  responseWriter.Write([]byte ("Criando novo snippet"))
} 