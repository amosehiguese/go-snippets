package snippet

import (
	"bytes"
	"fmt"
	"sync"
)

func ConfinementMain()  {
	printData := func (wg *sync.WaitGroup, data []byte)  {
		defer wg.Done()
		
		var buf bytes.Buffer
		for _, d := range data {
			fmt.Fprintf(&buf, "%c", d)
		}
		fmt.Println(buf.String())
	}

	var wg sync.WaitGroup
	var data = []byte("golang")

	wg.Add(2)
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])

	wg.Wait()
	
}



