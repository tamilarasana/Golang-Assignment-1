package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	var num int
	fmt.Println("Enter number")
	fmt.Scan(&num)
	fmt.Println("Factorial of ", num, " is : ", fact(num))
}