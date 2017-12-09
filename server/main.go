package main

import (
	"database/sql"
	"io"
	"log"
	"net/http"
	"os"

	"html/template"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hidelbreq/ratel-web/server/handler"
	"github.com/hidelbreq/ratel-web/util"
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

	pattern := filepath.Join("server", "view", "*.html")
	tmp, err := template.New("").Funcs(util.FuncMap).ParseGlob(pattern)
	if err != nil {
		log.Fatal(err)
	}

	opt := handler.Option{DB: *db, Tmpl: tmp}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("server/static"))))
	http.Handle("/", handler.NewTop(opt))
	http.Handle("/entries/", handler.NewEntry(opt))
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
