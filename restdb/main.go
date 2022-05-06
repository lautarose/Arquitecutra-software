package main

import (
   "github.com/gin-gonic/gin"
   "net/http"
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
)

type Body struct {
  // json tag to serialize json body
   Number int `json:"number"`
}

func main() {
  
   db, err := sql.Open("mysql", "root:pass@tcp(dbmysql:3306)/nums")
   if err != nil {
      panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
   }
   defer db.Close()


   stmtOut, err := db.Prepare("SELECT squareNumber FROM squarenum WHERE number = ?")
   if err != nil {
      panic(err.Error()) // proper error handling instead of panic in your app
   }
   defer stmtOut.Close()

 

   engine:=gin.New()
   engine.GET("/test", func(context *gin.Context) {
  
      // using BindJson method to serialize body with struct

    squareNum:=0 // we "scan" the result in here
    body:=Body{}
   err = stmtOut.QueryRow(13).Scan(&squareNum) // WHERE number = 13
   if err != nil {
      panic(err.Error()) // proper error handling instead of panic in your app
   }
      body.Number=squareNum;
      context.JSON(http.StatusAccepted,&body)
   })
   engine.Run(":3000")
}

//curl -X POST -H "Content-Type: application/json" -d '{"name":"lalala"}'