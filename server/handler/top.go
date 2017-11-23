package handler

import (
	"github.com/hidelbreq/ratel-web/model"
	"html/template"
	"net/http"
	"path/filepath"
	"fmt"
	"log"
	"encoding/json"
)

const (
	TOP_MUSIC_LOG = 3
)

type Top struct {
	soundCloudModel *model.SoundcloudTrackModel
	messageModel *model.MessageModel
}

func NewTop(opt Option) *Top {
	return &Top{
		soundCloudModel: model.NewSoundCloudTrackModel(opt.DB),
		messageModel: model.NewMessageModel(opt.DB),
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
	paths := []string{
		filepath.Join("server", "view", "index.html"),
		filepath.Join("server", "view", "amp-custom.html"),
	}
	tmp, err := template.ParseFiles(paths...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := &struct {
		SoundCloudTracks []*model.SoundCloudTrack
	}{
		SoundCloudTracks: ts,
	}
	tmp.Execute(w, data)
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
