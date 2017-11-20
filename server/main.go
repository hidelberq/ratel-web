package main

import (
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"database/sql"
	"github.com/hidelbreq/ratel-web/server/handler"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "user:pass@tcp(mysql:3306)/ratel?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("db connection suceeded.")

	http.FileServer(http.Dir("server/static/"))
	opt := handler.Option{DB: *db}
	http.Handle("/", handler.NewTop(opt))
	http.ListenAndServe(":8080", nil)
}
