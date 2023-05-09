package snippet

import "fmt"



func RepeatTakeMain() {
	repeat := func (done <-chan interface{}, values ...interface{} ) <-chan interface{} {
		valueStream := make(chan interface{})

		go func() {
			defer close(valueStream)

			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}

	take := func (done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{}  {
		dataStream := make(chan interface{})

		go func() {
			defer close(dataStream)

			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case dataStream <- <- valueStream:
				}
			}
		}()
		return dataStream
	}

	done := make(chan interface{})
	defer close(done)

	gen := take(done, repeat(done, 1, 2, 3, 4), 10)

	for v := range gen {
		fmt.Println(v)
	}

}
