package main

import (
	"log"
	"net/http"
  
)

//ListAndServer = async
//Response e Request
/* curl -i -X POST         http://localhost:4000/snippet/create
curl -i -X GET         http://localhost:4000/snippet/create

*/

func home(responseWriter http.ResponseWriter, request *http.Request) {
  if request.URL.Path != "/" {
    http.NotFound(responseWriter, request) 
    return
     
  }
  
  
  responseWriter.Write([]byte("Welcome"))
}

func showSnippet (responseWriter http.ResponseWriter, request *http.Request)  {
   responseWriter.Write([]byte("Show Snippet"))
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
    
     

func main() {
  //ideia de multiplexador
  //init do server
	mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet", showSnippet)
  mux.HandleFunc("/snippet/create", createSnippet)

  
  log.Println("Escutando na porta 4000")
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)

  
}




