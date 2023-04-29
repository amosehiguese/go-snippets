package snippet

import "fmt"

const boilingF = 212.0

func BoilingMain() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gF or %gC\n", f, c)
}