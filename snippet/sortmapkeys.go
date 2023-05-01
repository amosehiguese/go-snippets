package snippet

import (
	"fmt"
	"sort"
)

func SortMapKeysMain() {
	ages := map[string]int{"amos": 1, "josh": 2, "luke": 3, "mike": 4}
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

}