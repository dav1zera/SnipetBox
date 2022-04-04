package main

import (
  "net/http"
	"strconv"
  "fmt"
	
)

//ListAndServer = async
//Response e Request

/* 

  curl -i -X POST          http://localhost:4000/snippet/create
  curl -i -X GET           http://localhost:4000/snippet/create

*/

func home(responseWriter http.ResponseWriter, request *http.Request) {
  if request.URL.Path != "/" {
    http.NotFound(responseWriter, request) 
    return
     
  }
  
  
  responseWriter.Write([]byte("Welcome"))
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