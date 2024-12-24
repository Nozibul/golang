package main

import "fmt"

// Function definition with a return type of string
func HelloWorld(name string) string {
	// Concatenate "Hello" and the name and return the result
	return "Hello " + name
}

// Another example
func helloWorld2() {
	fmt.Print("hello world")
}
func main() {
	// Calling the function and printing the result
	fmt.Println(HelloWorld("World"))
	fmt.Println(HelloWorld("Go"))
	helloWorld2()
}
