package snippet

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Content-Type")	!= "application/json" {
				w.WriteHeader(http.StatusUnsupportedMediaType)
				w.Write([]byte("415 - Unsupported Media Type. Please send JSON"))
				return
			}
			handler.ServeHTTP(w, r)
		})
}

func setServerTimeCookie(handler http.Handler) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		cookie := http.Cookie{
			Name: "Server-Time(UTC)",
			Value: strconv.FormatInt(time.Now().Unix(), 10),
		}
		fmt.Println(time.Now().Unix(), 10)
		http.SetCookie(w, &cookie)
		
	})
	
}

type city struct {
	Name  string
	Area  uint64
}

func mainLogic(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "POST" {
		var tempCity city
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempCity)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		fmt.Printf("Got %s city with area of %d sq miles!\n", tempCity.Name, tempCity.Area)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201 - Created"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
	}
	
}

func Middlewaremain() {
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/", setServerTimeCookie(filterContentType(mainLogicHandler)))
	http.ListenAndServe(":8000", nil)
	// server := &http.Server{
	// 	Addr: "127.0.0.1:8000",
	// 	Handler: r,
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout: 15 * time.Second,
	// }
	// log.Fatal(server.ListenAndServe())
}