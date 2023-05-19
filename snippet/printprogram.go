package snippet

import (
	"fmt"
	"reflect"
	"strings"
)

type NewInt int

type NewPerson struct {
	Name    *Name
	Address *Address
}

type NewName struct {
	Title, First, Last string
}

type NewAddress struct {
	Street, Region string
}

func NewPrintProgramMain() {
	fmt.Println("Walking a simple integer")
	var one Int = 1
	walk(one, 0)

	fmt.Println("Walking a simple struct")
	two := struct{Name string}{"foo"}
	walk(two, 0)

	fmt.Println("Walking a struct with struct fields")
	p := &Person{
		Name: &Name{"Count", "Tyrone", "Rugen"},
		Address: &Address{"Humperdink Castle", "Florian"},
	}
	newwalk(p, 0)
}

func newwalk(u interface{}, depth int) {
	val := reflect.Indirect(reflect.ValueOf(u))
	t := val.Type()
	tabs := strings.Repeat("\t", depth+1)

	fmt.Printf("%sValue is type %q (%s)\n", tabs, t, val.Kind())
	if val.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldVal := reflect.Indirect(val.Field(i))

			tabs := strings.Repeat("\t", depth+2)
			fmt.Printf("%sField %q is type %q (%s)\n", tabs, field.Name, field.Type, fieldVal.Kind())

			if fieldVal.Kind() == reflect.Struct{
				walk(fieldVal.Interface(), depth+1)
			}
		}
	}
}