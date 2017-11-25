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
	"time"
	"fmt"
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
	id := strings.TrimPrefix(r.URL.Path, "/tracks/")
	if id == "" {
		switch r.Method {
		case "GET":
			trackHandler.showAll(w, r)
		}
	}

	i, err := strconv.Atoi(id);
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

func (trackHandler *Track) show(w http.ResponseWriter, r *http.Request, id int)  {
	track, err := trackHandler.trackModel.FindById(id)
	log.Println(track.DisplayAt)
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

func (trackHandler *Track) update(w http.ResponseWriter, r *http.Request, id int)  {
	name := r.FormValue("name")
	title := r.FormValue("title")
	author := r.FormValue("author")
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Println(err)
		return
	}

	displayAtStr := r.FormValue("display-at")
	displayAt, err  := time.ParseInLocation("2006-01-02T15:04", displayAtStr, loc)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}

	description := r.FormValue("description")
	updatedAt := time.Now()
	t := &model.Track{
		TrackId: id,
		Name: name,
		Title: title,
		Author:author,
		DisplayAt:displayAt,
		Description:description,
		UpdatedAt:updatedAt,
	}

	err = trackHandler.trackModel.Update(t)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/tracks/%d", id), 302)
}

