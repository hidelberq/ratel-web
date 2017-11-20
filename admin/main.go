package main

import (
	"database/sql"
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

	http.FileServer(http.Dir("admin/static/"))
	opt := handler.Option{DB: *db}
	http.Handle("/soundcloud-track", handler.NewSoundcloud(opt))
	http.ListenAndServe(":9090", nil)

}
