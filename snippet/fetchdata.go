package snippet

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var apis map[int]string
var ch chan map[int]interface{}

func FetchData(API int) {
	url := apis[API]
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()

		if body, err := io.ReadAll(resp.Body); err == nil {
			var result map[string]interface{}

			var re = make(map[int]interface{})
			json.Unmarshal([]byte(body), &result)

			switch API {
			case 1:
				if result["success"] == true {
					re[API] = result["rates"].(map[string]interface{})["USD"]
				} else {
					re[API] = result["error"].(map[string]interface{})["info"]
				}
				ch <- re
				fmt.Println("Result for API 1 stored")
			case 2:
				if result["main"] != nil {
					re[API] = result["main"].(map[string]interface{})["temp"]
				} else {
					re[API] = result["message"]
				}
				ch <- re
				fmt.Println("Result for API 2 stored")
			} 
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	} 
}

func FetchMain() {
	apis = make(map[int]string)
	ch = make(chan map[int]interface{})

	apis[1] = "http://data.fixer.io/api/latest?access_key=" + os.Getenv("ACCESS_KEY")

	apis[2] = "http://api.openweatermap.org/data/2.5/weather?" + "q=SINGAPORE&appid=" + os.Getenv("API_KEY")

	go FetchData(1)
	go FetchData(2)

	for i := 0; i < 2; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("All done!")

	fmt.Scanln()
}