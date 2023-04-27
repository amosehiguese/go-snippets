package snippet

import (
	"fmt"
	"os"
	"strings"
)

func ImprovedEcho() {
	output := strings.Join(os.Args[1:], "")
	fmt.Printf("%T, %[1]s", output)
}