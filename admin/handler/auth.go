package handler

import "net/http"

func BasicAuth(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if checkAuth(r) == false {
			w.Header().Set("WWW-Authenticate", `Basic realm="MY REALM"`)
			w.WriteHeader(401)
			w.Write([]byte("401 Unauthorized\n"))
			return
		}

		h.ServeHTTP(w, r)
	}
}
func checkAuth(r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		return false
	}

	return username == "grandcolline" && password == "Iwann@B1g"
}
