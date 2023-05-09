package snippet

import "sync"

var FanIn = func(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	multiplexStream := make(chan interface{})

	multiplex := func(c <-chan interface{}) {
		defer wg.Done()
		for v := range c {
			select {
			case <-done:
				return
			case multiplexStream <- v:
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexStream)
	}()

	return multiplexStream
}