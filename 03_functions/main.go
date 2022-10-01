package main

import "fmt"

func main(){
	fmt.Printf("Welcome to your first function Biodun \nThe sum of two numbers 5 and 6 is: %d \n", summation(5,6))
}

func summation (num1, num2 int) int {
	sum := num1 + num2
	return sum
}