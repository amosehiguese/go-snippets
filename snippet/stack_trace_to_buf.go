package snippet

import (
	"fmt"
	"runtime"
)

func StackTraceToBuffMain() {
	Buffoo()
}

func Buffoo() {
	Buffbar()
}

func Buffbar() {
	buf := make([]byte, 1024)
	runtime.Stack(buf, false)
	fmt.Printf("Trace:\n %s\n", buf)
}
