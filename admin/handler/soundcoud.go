package handler

import (
	"github.com/hidelbreq/ratel-web/model"
	"net/http"
	"io"
)

type Soundcloud struct {
	soundcloudModel *model.SoundcloudTrackModel
}

func NewSoundcloud(opt Option) *Soundcloud {
	return &Soundcloud{
		soundcloudModel: model.NewSoundCloudTrackModel(opt.DB),
	}
}

func (a *Soundcloud) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "This HTTP response has both headers before this text and trailers at the end.\n")
}