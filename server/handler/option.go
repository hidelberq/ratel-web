package handler

import (
	"database/sql"
	"html/template"
)

type Option struct {
	DB   sql.DB
	Tmpl *template.Template
}
