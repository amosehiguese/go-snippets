package snippet

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	Lissajous(w)
}