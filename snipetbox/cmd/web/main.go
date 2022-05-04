package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"snipetbox/pkg/models/mysql"

	_ "github.com/go-sql-driver/mysql"
)	

type applicationAux struct {
  errorsLog *log.Logger
  infoLogs *log.Logger
  snippets *mysql.SnippetModel
  
  
}

 func main() {
  portAdress := flag.String("addr", ":4000", "Porta da Rede")
  dsn := flag.String("dsn", "sowvTXioiS:6UF6FTDZPt@tcp(remotemysql.com)/sowvTXioiS?parseTime=true", "MySql DSN")
  flag.Parse()

  

  infoLogs := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
  errorsLog := log.New(os.Stderr,"ERRO:\t", log.Ldate|log.Ltime|log.Lshortfile)

    dataBase, err := openDB(*dsn) 
   if err != nil {
    errorsLog.Fatal(err)
   }

  defer dataBase.Close()
   
app := &applicationAux {
  errorsLog: errorsLog,
  infoLogs: infoLogs,
  snippets: &mysql.SnippetModel{DB:dataBase},
}
   
  serverConfigs :=  &http.Server {
    Addr: *portAdress,
    ErrorLog: errorsLog,
    Handler: app.routes(),
  }
   
  
  infoLogs.Printf("Escutando na porta %s\n", *portAdress)
  err = serverConfigs.ListenAndServe()
  errorsLog.Fatal(err)
  
 
}

func openDB(dsn string) (*sql.DB, error) {
  dataBase, err := sql.Open("mysql", dsn)
  if err != nil {
    return nil, err 
  }
  if err = dataBase.Ping(); err != nil {
    return nil, err
  }
  return dataBase, nil
}



  

    


    





