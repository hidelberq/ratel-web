package main

import (
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"database/sql"
	"github.com/hidelbreq/ratel-web/server/handler"
)

func main() {
	db, err := sql.Open("mysql", "user:pass@tcp(mysql:3306)/ratel?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	http.FileServer(http.Dir("server/static/"))
	opt := handler.Option{DB: *db}
	http.Handle("/", handler.NewTop(opt))
	http.ListenAndServe(":8080", nil)
}
