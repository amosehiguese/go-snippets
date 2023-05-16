package snippet

import (
	"fmt"
	"net/http"
)



func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie {
		Name: "first_cookie",
		Value: "Go",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name: "second_cookie",
		Value: "Manning",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	// h := r.Header["Cookie"]
	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
}

func CookieMain() {
	http.HandleFunc("/set-cookie", setCookie)
	http.HandleFunc("/get-cookie", getCookie)

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	fmt.Println("running server - ", server.Addr)
	server.ListenAndServe()
}