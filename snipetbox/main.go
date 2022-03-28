package main

import (
	"log"
	"net/http"
)

//ListAndServer = async
//Response e Request

func home(responseWriter http.ResponseWriter, request *http.Request) {
  responseWriter.Write([]byte("Welcome") )
}

func main() {
  //ideia de multiplexador
  //init do server
	mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  log.Println("Escutando na porta 4000")
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)

  
}
