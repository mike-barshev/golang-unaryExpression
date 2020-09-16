package main

import "fmt"

func unaryExpression(x *int) {
	*x++
}

func main() {
	num := 1
	unaryExpression(&num)
	fmt.Println(num)
}
