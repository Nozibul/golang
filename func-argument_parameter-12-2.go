/*
                // Go Function Parameters and Arguments
                      // Parameters and Arguments
Information can be passed to functions as a parameter. Parameters act as variables inside the function.

Parameters and their types are specified after the function name, inside the parentheses. You can add as many parameters as you want, just separate them with a comma:

Syntax
func FunctionName(param1 type, param2 type, param3 type) {
  // code to be executed
} 




                 // Function With Parameter Example
The following example has a function with one parameter (fname) of type string. When the familyName() function is called, we also pass along a name (e.g. Liam), and the name is used inside the function, which outputs several different first names, but an equal last name:

Example
package main
import ("fmt")

func familyName(fname string) {
  fmt.Println("Hello", fname, "Refsnes")
}

func main() {
  familyName("Liam")
  familyName("Jenny")
  familyName("Anja")
}
Result:

Hello Liam Refsnes
Hello Jenny Refsnes
Hello Anja Refsnes
Note: When a parameter is passed to the function, it is called an argument. So, from the example above: fname is a parameter, while Liam, Jenny and Anja are arguments.



                         // Multiple Parameters
Inside the function, you can add as many parameters as you want:

Example
package main
import ("fmt")

func familyName(fname string, age int) {
  fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

func main() {
  familyName("Liam", 3)
  familyName("Jenny", 14)
  familyName("Anja", 30)
}
Result:

Hello 3 year old Liam Refsnes
Hello 14 year old Jenny Refsnes
Hello 30 year old Anja Refsnes


Note: When you are working with multiple parameters,
the function call must have the same number of arguments as there
are parameters, and the arguments must be passed in the same order.

*/