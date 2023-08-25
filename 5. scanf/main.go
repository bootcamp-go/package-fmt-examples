package main

import "fmt"

func main() {

	var (
		name, surname string
		age           int
	)

	fmt.Scanf("%s %s %d", &name, &surname, &age)

	fmt.Println("The name is", name)
	fmt.Println("The surname is", surname)
	fmt.Println("The age is", age)
}
