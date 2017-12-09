package handler

import (
	"net/http"

	"log"

	"html/template"

	"github.com/hidelbreq/ratel-web/model"
)

const (
	ENTRY_NUM = 5
)

type Entry struct {
	tmpl       *template.Template
	entryModel *model.EntryModel
}

func NewEntry(opt Option) *Entry {
	return &Entry{
		tmpl:       opt.Tmpl,
		entryModel: model.NewEntryModel(opt.DB),
	}
}

// TODO: limit, offset
func (e *Entry) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	es, err := e.entryModel.FindLatest(ENTRY_NUM)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	data := &struct {
		Entries []*model.Entry
	}{
		es,
	}

	if err = e.tmpl.ExecuteTemplate(w, "entries", data); err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}

}
