package snippet

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	User 		string
	Threads []string
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `
	<html>
  <head>
		<title>Go</title>
	</head>
  <body>
    <h1>Hello World</h1>
  </body>
</html>`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://goole.com")
	w.WriteHeader(302)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User: "Amos",
		Threads: []string{"new", "old", "young"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func WebMain() {
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", GetPosts)

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	fmt.Println("running server - ", server.Addr)
	server.ListenAndServe()
}