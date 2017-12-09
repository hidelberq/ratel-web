package main

import (
	"database/sql"
	"io"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	//"github.com/hidelbreq/ratel-web/server/handler"
	"log"

	"github.com/hidelbreq/ratel-web/server/handler"
)

var db *sql.DB

func main() {
	var err error

	logfile, err := os.OpenFile("./logs/server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("cannnot open test.log:" + err.Error())
	}
	defer logfile.Close()
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)

	db, err = sql.Open("mysql", "user:pass@tcp(mysql:3306)/ratel?parseTime=true&loc=Asia%2FTokyo&charset=utf8mb4")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("db connection suceeded.")

	opt := handler.Option{DB: *db}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("server/static"))))
	http.Handle("/", handler.NewTop(opt))
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
