package snippet

import (
	"fmt"
	"math/rand"
)



func RepeatFnMain() {
	repeatfn := func (done <-chan interface{}, fn func()interface{}) <- chan interface{}  {
		fnStream := make(chan interface{})

		go func() {
			defer close(fnStream)

			for{
				select {
				case <-done:
					return
				case fnStream <- fn():
				}
			}
		}()

		return fnStream
	}

	take := func (done <-chan interface{}, fnStream <-chan interface{}, num int) <-chan interface{}  {
		stream := make(chan interface{})

		go func() {
			defer close(stream)

			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case stream <- <- fnStream:
				}
			}
		}()
		return stream
	}

	done := make(chan interface{})
	defer close(done)

	randInt := func () interface{}  {
		return rand.Int()
	}

	for v := range take(done, repeatfn(done, randInt), 3) {
		fmt.Println(v)
	}

}
