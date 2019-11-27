package main

import "fmt"

// notice that the type is declared after the var
// if all the var types are the same, you can omit all but last declaration
func add(x int, y int) int {
	return x + y
}

// A funtion can return any number of resutls
func swap(x, y string) (string, string) {
	return y, x
}

func functions() {
	fmt.Println(add(42, 99))
	fmt.Println(swap("Hello", "Goodbye"))
}
