package main

import "fmt"

type myType struct{
	s string
}

func main() {
	var (
		a *myType
		b *myType
	)
	if a == b {
		fmt.Println(a, b)
		fmt.Println("lol")
	}
}