/*
Go has three functions to output text:

Print()
Println()
Printf() 

         The Print() Function
The Print() function prints its arguments with their default format

 If we want to print the arguments in new lines, we need to use \n.
If we want to add a space between string arguments, we need to use " ":


                             The Printf() Function
The Printf() function first formats its argument based on the given formatting verb and then prints them.

Here we will use two formatting verbs:

%v is used to print the value of the arguments
%T is used to print the type of the arguments

*/

package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  // If we want to print the arguments in new lines, we need to use \n.
  fmt.Print(i, "\n") ;// Hello
   // or
  fmt.Print(i, "\n", j); // World

  //If we want to add a space between string arguments, we need to use " ":
  fmt.Println(i, " ", j)



  var n = "Go"
  var o = 34
  var k string = "Hello"
  var m int = 15

  fmt.Printf("%v has value: %v and type: %T\n", n, o, n);// Go has value: 34 and type: string
  fmt.Printf("i has value: %v and type: %T\n", k, k);// i has value: Hello and type: string
  fmt.Printf("j has value: %v and type: %T", m, m); // j has value: 15 and type: int
}