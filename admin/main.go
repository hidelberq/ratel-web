package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/hidelbreq/ratel-web/admin/handler"
)

func main() {
	db, err := sql.Open("mysql", "user:pass@tcp(mysql:3306)/ratel?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	if err := db.Ping(); err != nil {
		panic(err.Error())
	}

	log.Println("DB connection succeeded.")
	opt := handler.Option{DB: *db}
	http.Handle("/tracks/", handler.NewSoundcloud(opt))
	http.Handle("/entries/", handler.NewEntriesHandler(opt))
	http.ListenAndServe(":8080", nil)

}
