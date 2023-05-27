package snippet

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var ErrTimeout = errors.New("the request timed out")
var ErrRejected = errors.New("the request was rejected")

var seed = time.Now().Unix()

func ErrVMain() {
	response, err := SendRequest("Hello")
	for err == ErrTimeout {
		fmt.Println("Timeout. Retrying...")
		response, err = SendRequest("Hello")
	}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}

func SendRequest(req string) (string, error) {
	rand.NewSource(seed)

	switch rand.Intn(50) % 3 {
	case 0:
		return "Success", nil
	case 1:
		return "", ErrRejected
	default:
		return "", ErrTimeout
	}
}