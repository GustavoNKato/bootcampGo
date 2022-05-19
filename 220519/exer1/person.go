package exer1

import "fmt"

var (
	name string
	age  int8
)

func Person(name string, age int8) {
	fmt.Printf("Meu nome é %s e tenho %v anos!", name, age)
}
