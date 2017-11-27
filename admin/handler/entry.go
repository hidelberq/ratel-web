package handler

import (
	"log"
	"net/http"
	"path/filepath"

	"database/sql"
	"strconv"
	"strings"

	"time"

	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/hidelbreq/ratel-web/model"
	"github.com/hidelbreq/ratel-web/util"
)

type EntriesHandler struct {
	entryModel *model.EntryModel
}

func NewEntriesHandler(opt Option) *EntriesHandler {
	return &EntriesHandler{
		entryModel: model.NewEntryModel(opt.DB),
	}
}

func (eh *EntriesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	strId := strings.TrimPrefix(r.URL.Path, "/entries/")
	if strId == "" {
		switch r.Method {
		case http.MethodGet:
			eh.showAll(w, r)
		case http.MethodPost:
			eh.post(w, r)
		}
		return
	}

	id, err := strconv.Atoi(strId)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		eh.show(w, r, id)
	case http.MethodPost:
		eh.update(w, r, id)
	}
}

func (eh *EntriesHandler) showAll(w http.ResponseWriter, r *http.Request) {
	es, err := eh.entryModel.FindAll(100)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	data := &struct {
		Entries []*model.Entry
	}{
		es,
	}
	err = util.ExecuteOnAdminBase(w, data, filepath.Join("admin", "view", "entries.html"))
	if err != nil {
		log.Println(err)
		return
	}
}

func (eh *EntriesHandler) post(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	author := r.FormValue("author")
	body := r.FormValue("body")
	strDisplayAt := r.FormValue("display-at")
	deleted := r.FormValue("deleted")

	displayAt, err := util.ParseLocateTimeInJST(strDisplayAt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(strDisplayAt))
		return
	}

	var deletedAt mysql.NullTime
	if deleted == "on" {
		deletedAt.Valid = true
		deletedAt.Time = time.Now()
	}

	e := &model.Entry{
		Title:     title,
		Author:    author,
		Body:      body,
		DisplayAt: displayAt,
		DeletedAt: deletedAt,
	}
	_, err = eh.entryModel.Create(e)
	if err != nil {
		util.ExecuteAdminInternalError(w, err)
		return
	}

	http.Redirect(w, r, "/entries/", 302)
}

func (eh *EntriesHandler) show(w http.ResponseWriter, r *http.Request, id int) {
	entry, err := eh.entryModel.FindById(id)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	data := &struct {
		Entry *model.Entry
	}{
		entry,
	}
	err = util.ExecuteOnAdminBase(w, data, filepath.Join("admin", "view", "entry.html"))
	if err != nil {
		log.Println(err)
	}
}

func (eh *EntriesHandler) update(w http.ResponseWriter, r *http.Request, id int) {
	e := makeEntry(r)
	if e == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := eh.entryModel.Update(e, id)
	if err != nil {
		util.ExecuteAdminInternalError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/entries/%d", id), http.StatusFound)
}

func makeEntry(r *http.Request) *model.Entry {
	title := r.FormValue("title")
	author := r.FormValue("author")
	body := r.FormValue("body")
	strDisplayAt := r.FormValue("display-at")
	deleted := r.FormValue("deleted")

	displayAt, err := util.ParseLocateTimeInJST(strDisplayAt)
	if err != nil {
		return nil
	}

	var deletedAt mysql.NullTime
	if deleted == "on" {
		deletedAt.Valid = true
		deletedAt.Time = time.Now()
	}

	return &model.Entry{
		Title:     title,
		Author:    author,
		Body:      body,
		DisplayAt: displayAt,
		DeletedAt: deletedAt,
	}
}
