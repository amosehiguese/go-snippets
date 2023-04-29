package snippet

import "fmt"

func GCDMain() {
	x := 1
	y := 10
	result := gcd(x, y)
	fmt.Println(result)
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x % y 
	}
	return x
}