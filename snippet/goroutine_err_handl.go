package snippet

import (
	"fmt"
	"net/http"
)

type Result struct {
	err error
	resp *http.Response
}

func ErrHandlMain() {
	checkStatus := func (done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)

		go func() {
			defer close(results)

			for _, url := range urls {
				resp, err := http.Get(url)

				result := Result{err, resp}
				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.google.com", "https://www.bing.com"}
	
	errCounter := 0
	for results := range checkStatus(done, urls...) {
		if results.err != nil {
			fmt.Printf("Err: %v", results.err)
			errCounter++

			if errCounter >= 3 {
				fmt.Println("too many errors, breaking")
				break
			}
			continue
		}
		fmt.Printf("Response: %v\n", results.resp.Status)
	}
}
