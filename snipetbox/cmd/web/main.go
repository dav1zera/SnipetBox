package main

import (
  "flag"
	"log"
	"net/http"
  "os"
  
)	

type application struct {
  errorsLog *log.Logger
  infoLogs *log.Logger
}

 func main() {
  //ideia de multiplexador
  //init do server
  portAdress := flag.String("addr", ":4000", "Porta da Rede")
  flag.Parse()


  infoLogs := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
  errorsLog := log.New(os.Stderr,"ERRO:\t", log.Ldate|log.Ltime|log.Lshortfile)

app := &application {
  errorsLog: errorsLog,
  infoLogs: infoLogs,
}
   
	mux := http.NewServeMux()
  mux.HandleFunc("/", app.home)
  mux.HandleFunc("/snippet", app.showSnippet)
  mux.HandleFunc("/snippet/create", app.createSnippet)

  fileServer := http.FileServer(http.Dir("./ui/static/"))
  mux.Handle("/static/", http.StripPrefix("/static", fileServer))

  infoLogs.Printf("Escutando na porta %s\n", *portAdress)
  err := http.ListenAndServe(*portAdress, mux)
  errorsLog.Fatal(err)
 
}



  

    


    





