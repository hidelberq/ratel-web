package handler

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/hidelbreq/ratel-web/model"
	"github.com/hidelbreq/ratel-web/util"
	"strings"
	"strconv"
)

type Track struct {
	trackModel *model.TrackModel
}

func NewSoundcloud(opt Option) *Track {
	return &Track{
		trackModel: model.NewTrackModel(opt.DB),
	}
}

func (trackHandler *Track) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id := strings.TrimPrefix(r.URL.Path, "/tracks/")
		if id == "" {
			trackHandler.showAll(w, r)
		} else {
			trackHandler.show(w, r, id)
		}
	}


}

func (trackHandler *Track) showAll(w http.ResponseWriter, r *http.Request)  {
	ts := trackHandler.trackModel.FindLatest(100)

	paths := []string {
		filepath.Join("admin", "view", "base.html"),
		filepath.Join("admin", "view", "tracks.html"),
	}
	tmpl, err := template.New("base").Funcs(util.FuncMap).ParseFiles(paths...)
	if err != nil {
		log.Println(err)
		return
	}

	data := &struct {
		Tracks []*model.Track
	}{
		ts,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		return
	}
}

func (trackHandler *Track) show(w http.ResponseWriter, r *http.Request, id string)  {
	i, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	track, err := trackHandler.trackModel.FindById(i)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	if track == nil {
		w.WriteHeader(404)
		return
	}

	paths := []string {
		filepath.Join("admin", "view", "base.html"),
		filepath.Join("admin", "view", "track.html"),
	}
	tmpl, err := template.New("base").Funcs(util.FuncMap).ParseFiles(paths...)
	if err != nil {
		log.Println(err)
		return
	}

	data := &struct {
		Track *model.Track
	}{
		track,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		return
	}
}