package main

import (
	"database/sql"
	"io"
	"os"
	"log"
	"net/http"

	"github.com/hidelbreq/ratel-web/admin/handler"
)

var db *sql.DB

func main() {
	var err error

	logfile, err := os.OpenFile("./logs/admin.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("cannnot open test.log:" + err.Error())
	}
	defer logfile.Close()
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)

	db, err = sql.Open("mysql", "user:pass@tcp(mysql:3306)/ratel?parseTime=true&loc=Asia%2FTokyo&charset=utf8mb4")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	if err := db.Ping(); err != nil {
		panic(err.Error())
	}

	log.Println("DB connection succeeded.")
	opt := handler.Option{DB: *db}
	http.HandleFunc("/", handler.BasicAuth(handler.NewTopHandler(opt)))
	http.HandleFunc("/tracks/", handler.BasicAuth(handler.NewSoundcloud(opt)))
	http.HandleFunc("/entries/", handler.BasicAuth(handler.NewEntriesHandler(opt)))
	http.ListenAndServe(":8080", nil)

}
