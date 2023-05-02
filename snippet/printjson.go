package snippet

import "fmt"

func PrintJSON(v interface{}) {
	switch vv := v.(type) {
	case string:
		fmt.Println("is string", vv)
	case float64:
		fmt.Println("is float64", vv)
	case []interface{}:
		fmt.Println("is an array:")
		for i, u := range vv {
			fmt.Print(i, " ")
			PrintJSON(u)
		}
	case map[string]interface{}:
		fmt.Println("is an object:")
		for i, u := range vv {
			fmt.Print(i, " ")
			PrintJSON(u)
		}
	default:
		fmt.Println("Unknown type")
	}

}