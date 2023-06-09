package snippet

import "fmt"



func BridgeMain() {
	orDone := func (done, c <-chan interface{}) <-chan interface{}  {
		valueStream := make(chan interface{})

		go func() {
			defer close(valueStream)

			for{
				select {
				case <-done:
					return
				case v, ok := <-c:
					if !ok {
						return
					}
					select {
					case valueStream <- v:
					case <-done:	
					}
				}
			}
		}()
		return valueStream
	}

	bridge := func (done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{}  {
		
		valueStream:= make(chan interface{})

		go func() {
			defer close(valueStream)
			for {
				var stream <-chan interface{}
				select {
				case maybeStream, ok := <-chanStream:
					if !ok{
						return
					}
					stream = maybeStream
				case <-done:
					return
				}
				for val := range orDone(done, stream) {
					select {
					case valueStream <-val:
					case <-done:
					}
				}
			}
		}()
		return valueStream
	}

	genVals := func() <-chan <-chan interface{} {
		chanStream := make(chan(<-chan interface{}))
		go func ()  {
			defer close(chanStream)
			for i := 0; i < 10; i++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				chanStream <- stream
			}
		}()
		return chanStream
	}

	for v := range bridge(nil, genVals()) {
		fmt.Printf("%v ", v )
	}

}
