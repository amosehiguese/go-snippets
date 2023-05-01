package snippet

import "fmt"

func Remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func slMain() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(Remove(s, 2))
}