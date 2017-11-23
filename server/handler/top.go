package handler

import (
	"github.com/hidelbreq/ratel-web/model"
	"html/template"
	"net/http"
	"path/filepath"
	"log"
	"encoding/json"
	"strings"
)

const (
	TOP_MUSIC_LOG = 3
)

type Top struct {
	soundCloudModel *model.SoundcloudTrackModel
	messageModel *model.MessageModel
	entryModel *model.EntryModel
}

func NewTop(opt Option) *Top {
	return &Top{
		soundCloudModel: model.NewSoundCloudTrackModel(opt.DB),
		messageModel: model.NewMessageModel(opt.DB),
		entryModel: model.NewBlogModel(opt.DB),
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
	ts := t.soundCloudModel.FindAll(TOP_MUSIC_LOG)
	es, err := t.entryModel.FindLatest(2)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	paths := []string{
		filepath.Join("server", "view", "index.html"),
		filepath.Join("server", "view", "amp-custom.html"),
	}

	funcMap := map[string]interface{}{
		"nl2br": func(text string) template.HTML {
			return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br>", -1))
		},
	}

	tmp, err := template.New("base").Funcs(funcMap).ParseFiles(paths...)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	data := &struct {
		SoundCloudTracks []*model.SoundCloudTrack
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
	w.WriteHeader(200)
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

	j, err := json.Marshal(struct {Name string} {name})
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
