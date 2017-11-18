package handler

import (
	"github.com/hidelbreq/ratel-web/model"
	"html/template"
	"net/http"
	"path/filepath"
	"fmt"
	"log"
)

type Top struct {
	soundCloudModel *model.SoundcloudTrackModel
}

func NewTop(opt Option) *Top {
	return &Top{soundCloudModel: model.NewSoundCloudTrackModel(opt.DB)}
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
	ts := t.soundCloudModel.FindAll()
	tmp, err := template.ParseFiles(filepath.Join("server", "view", "index.html"))
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
	log.Fatal(r.PostForm)
}

var _ http.Handler = (*Top)(nil)
