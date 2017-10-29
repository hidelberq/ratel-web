package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hidelbreq/ratel-web/server/handler"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "user:pass@tcp(localhost:3306)/ratel?parseTime=true&loc=Asia%2FTokyo")
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
