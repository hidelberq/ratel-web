package handler

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"fmt"
	"github.com/hidelbreq/ratel-web/model"
	"github.com/hidelbreq/ratel-web/util"
	"strconv"
	"strings"
	"time"
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
	trimed := strings.TrimPrefix(r.URL.Path, "/tracks/")
	if trimed == "" {
		switch r.Method {
		case "GET":
			trackHandler.showAll(w, r)
		case "POST":
			trackHandler.post(w, r)
		}
	}

	i, err := strconv.Atoi(trimed)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	switch r.Method {
	case "GET":
		trackHandler.show(w, r, i)
	case "POST":
		trackHandler.update(w, r, i)
	}

}

func (trackHandler *Track) showAll(w http.ResponseWriter, r *http.Request) {
	ts := trackHandler.trackModel.FindLatest(100)
	paths := []string{
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

func (trackHandler *Track) show(w http.ResponseWriter, r *http.Request, id int) {
	track, err := trackHandler.trackModel.FindById(id)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	if track == nil {
		w.WriteHeader(404)
		return
	}

	paths := []string{
		filepath.Join("admin", "view", "base.html"),
		filepath.Join("admin", "view", "track-update.html"),
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

func (trackHandler *Track) update(w http.ResponseWriter, r *http.Request, id int) {
	trackIdStr := r.FormValue("track-id")
	trackId, err := strconv.Atoi(trackIdStr)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}

	name := r.FormValue("name")
	title := r.FormValue("title")
	author := r.FormValue("author")
	displayAtStr := r.FormValue("display-at")

	displayAt, err := util.ParseLocateTimeInJST(displayAtStr)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}

	description := r.FormValue("description")
	updatedAt := time.Now()
	t := &model.Track{
		TrackId:     trackId,
		Name:        name,
		Title:       title,
		Author:      author,
		DisplayAt:   displayAt,
		Description: description,
		UpdatedAt:   updatedAt,
	}

	err = trackHandler.trackModel.Update(id, t)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/tracks/%d", trackId), 302)
}

func (trackHandler *Track) post(w http.ResponseWriter, r *http.Request) {
	trackId := r.FormValue("track-id")
	id, err := strconv.Atoi(trackId)
	if err != nil {
		w.WriteHeader(400)
		log.Println(err)
		return
	}

	name := r.FormValue("name")
	title := r.FormValue("title")
	author := r.FormValue("author")
	displayAtStr := r.FormValue("display-at")
	displayAt, err := util.ParseLocateTimeInJST(displayAtStr)
	if err != nil {
		w.WriteHeader(400)
		log.Println(err)
		return
	}

	t := &model.Track{
		TrackId:   id,
		Name:      name,
		Title:     title,
		Author:    author,
		DisplayAt: displayAt,
		UpdatedAt: time.Now(),
	}
	if err := trackHandler.trackModel.Create(t); err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	http.Redirect(w, r, "/tracks", 302)
}
