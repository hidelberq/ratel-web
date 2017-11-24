package util

import (
	"html/template"
	"strings"
)

var FuncMap = map[string]interface{}{
	"nl2br": func(text string) template.HTML {
		return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br>", -1))
	},
}
