package handler

import (
	"net/http"
	"path/filepath"

	"log"

	"github.com/hidelbreq/ratel-web/util"
)

type TopHandler struct{}

func NewTopHandler(opt Option) *TopHandler {
	return &TopHandler{}
}

func (eh *TopHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if checkAuth(r) == false {
		w.Header().Set("WWW-Authenticate", `Basic realm="MY REALM"`)
		w.WriteHeader(401)
		w.Write([]byte("401 Unauthorized\n"))
		return
	}

	if err := util.ExecuteOnAdminBase(w, nil, filepath.Join("admin", "view", "index.html")); err != nil {
		log.Println(err)
	}
}
