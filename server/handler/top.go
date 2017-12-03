package handler

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/hidelbreq/ratel-web/model"
	"github.com/hidelbreq/ratel-web/util"
)

const (
	MUSIC_LOG_NUM = 3
	ENTRY_NUM     = 2
)

type Top struct {
	soundCloudModel *model.TrackModel
	messageModel    *model.MessageModel
	entryModel      *model.EntryModel
}

func NewTop(opt Option) *Top {
	return &Top{
		soundCloudModel: model.NewTrackModel(opt.DB),
		messageModel:    model.NewMessageModel(opt.DB),
		entryModel:      model.NewEntryModel(opt.DB),
	}
}

func (t *Top) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		t.show(w, r)
	case "POST":
		t.post(w, r)
	}
}

func (t *Top) show(w http.ResponseWriter, r *http.Request) {
	ts := t.soundCloudModel.FindLatest(MUSIC_LOG_NUM)
	es, err := t.entryModel.FindLatest(ENTRY_NUM)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	paths := []string{
		filepath.Join("server", "view", "index.html"),
		filepath.Join("server", "view", "amp-custom.html"),
	}

	tmp, err := template.New("base").Funcs(util.FuncMap).ParseFiles(paths...)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	data := &struct {
		SoundCloudTracks []*model.Track
		Entries          []*model.Entry
	}{
		SoundCloudTracks: ts,
		Entries:          es,
	}
	if err = tmp.Execute(w, data); err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
}

func (t *Top) post(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")

	m := &model.Message{name, email, message}
	if err := t.messageModel.Create(m); err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	j, err := json.Marshal(struct{ Name string }{name})
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("origin"))
	w.Header().Set("AMP-Access-Control-Allow-Source-Origin", r.FormValue("__amp_source_origin"))
	w.WriteHeader(200)
	w.Write(j)
}

var _ http.Handler = (*Top)(nil)
