package snippet

import "fmt"

var m = make(map[string]int)

func K(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[K(list)]++
}

func Count(list []string) int {
	return m[K(list)]
}
