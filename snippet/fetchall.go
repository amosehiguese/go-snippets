package snippet

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func FetchAll(url string, ch chan<- string) {
	start := time.Now()

	response, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	numberOfBytes, err :=  io.Copy(io.Discard, response.Body)
	response.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <-fmt.Sprintf("%.2fs  %7d  %s", secs, numberOfBytes, url)
}