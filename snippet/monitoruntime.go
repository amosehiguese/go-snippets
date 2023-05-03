package snippet

import (
	"log"
	"runtime"
	"time"
)

func MonitorRuntime() {
	log.Println("Number of CPUs:", runtime.NumCPU())
	m := &runtime.MemStats{}
	for {
		r := runtime.NumGoroutine()
		log.Println("Number of goroutines", r)

		// This runtime.ReadMemStats momentarily halt the Go runtime and can have a performance impact on your app. Don't do often and only when in a debug mode
		runtime.ReadMemStats(m)
		log.Println("Allocated memory", m.Alloc)
		time.Sleep(10 * time.Second)
	}
}

func MonitorRuntimeMain() {
	go MonitorRuntime()

	i := 0
	for i < 40 {
		go func() {
			time.Sleep(15 * time.Second)
		}()
		i = i + 1
		time.Sleep(1 * time.Second)
	}
}