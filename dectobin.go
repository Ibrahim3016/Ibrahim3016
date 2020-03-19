package main

import "fmt"

func main() {

	fmt.Println("This program lets you input a decimal integer and outputs the number in binary format")

	var decimal int

	fmt.Println("Please input a decimal integer and press ENTER")

	fmt.Scan(&decimal)

	fmt.Printf("decimal = %d, in binary format = %b\n", decimal, decimal)
}
