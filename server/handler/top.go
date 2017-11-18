package handler

import (
	"github.com/hidelbreq/ratel-web/model"
	"html/template"
	"net/http"
	"path/filepath"
	"time"
	"fmt"
)

type Top struct {
	soundCloudModel *model.SoundCloudTrackModel
}

func NewTop(opt Option) *Top {
	return &Top{soundCloudModel: model.NewSoundCloudTrackModel(opt.DB)}
}

func (t *Top) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		t.show(w, r)
	}
}

func dateOfMonsh(t time.Time) string {
	return t.Format("01-02")
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

var _ http.Handler = (*Top)(nil)
