package util

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

var FuncMap = map[string]interface{}{
	"nl2br": func(text string) template.HTML {
		return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br>", -1))
	},
	"truncate": func(text string, length int) template.HTML {
		replaced := strings.Replace(text, "\n", "", -1)
		r := []rune(replaced)
		return template.HTML(r[0:length])
	},
	"format4datetimelocal": func(t time.Time) string {
		return t.Format(FormatDateTimeLocal)
	},
}

func ExecuteOnAdminBase(w http.ResponseWriter, data interface{}, path ...string) error {
	basePath := filepath.Join("admin", "view", "base.html")
	tmpl, err := template.New("base").Funcs(FuncMap).ParseFiles(append(path, basePath)...)
	if err != nil {
		return err
	}

	return tmpl.Execute(w, data)
}

func ExecuteAdminInternalError(w http.ResponseWriter, err error) {
	log.Println(err)
	data := &struct {
		Err error
	}{
		err,
	}
	ExecuteOnAdminBase(w, data, filepath.Join("admin", "view", "error-500.html"))
	w.WriteHeader(http.StatusInternalServerError)
}
