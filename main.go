/*
Go Syntax
A Go file consists of the following parts:

Package declaration
Import packages
Functions
Statements and expressions
*/
package main // must use Package declaration

import "fmt" // must be import fmt Import packages .er name holo formate and fmt use na korle
//amra kono kisu print korte parbo na bec fmt er moddhe printli ta ase jer maddhome amra print korbo
//and dubble cotetion er moddhe fmt use korte hobe

// func means function
func main() { // Change mainFn to main
	// Statements and expressions
	name := "Nozibul Islam"
	var age = 54
	fmt.Print("I am a ")
	fmt.Println("Go developer", name, age) // Output: I am a Go developer Nozibul Islam 54
}
