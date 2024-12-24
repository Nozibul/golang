/*
Go Arrays
Arrays are used to store multiple values of the same type in a single variable,
instead of declaring separate variables for each value.
অ্যারেগুলি প্রতিটি মানের জন্য পৃথক ভেরিয়েবল ঘোষণা করার পরিবর্তে একটি একক
ভেরিয়েবলে একই ধরণের একাধিক মান সংরক্ষণ করতে ব্যবহৃত হয়।


              // Declare an Array
In Go, there are two ways to declare an array:

               // 1. With the var keyword:
Syntax:
var array_name = [length]datatype{values} // here length is defined

or
var array_name = [...]datatype{values} // here length is inferred


               // 2. With the := sign:
Syntax:
array_name := [length]datatype{values} // here length is defined

or

array_name := [...]datatype{values} // here length is inferred


Note: The length specifies the number of elements to store in the array.
In Go, arrays have a fixed length. The length of the array is either
defined by a number or is inferred (means that the compiler decides the
length of the array, based on the number of values).

               //Array Examples
Example: 01
This example declares two arrays (arr1 and arr2) with defined lengths:

package main
import ("fmt")

func main() {
  var arr1 = [3]int{1,2,3}
  arr2 := [5]int{4,5,6,7,8}

  fmt.Println(arr1) // [1 2 3]
  fmt.Println(arr2) // [4 5 6 7 8]
}


Example : 02
This example declares two arrays (arr1 and arr2) with inferred lengths:

package main
import ("fmt")

func main() {
  var arr1 = [...]int{1,2,3}
  arr2 := [...]int{4,5,6,7,8}

  fmt.Println(arr1)
  fmt.Println(arr2)
}
Result:

[1 2 3]
[4 5 6 7 8]


Example : 03
This example declares an array of strings:

package main
import ("fmt")

func main() {
  var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  fmt.Print(cars)
}
Result:

[Volvo BMW Ford Mazda]


            // Access Elements of an Array
You can access a specific array element by referring to the index number.

In Go, array indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.

Example: 04
This example shows how to access the first and third elements in the prices array:

package main
import ("fmt")

func main() {
  prices := [3]int{10,20,30}

  fmt.Println(prices[0]) // 10
  fmt.Println(prices[2]) // 30
}


          //Change Elements of an Array
You can also change the value of a specific array element by referring to the index number.

Example: 05
This example shows how to change the value of the third element in the prices array:

package main
import ("fmt")

func main() {
  prices := [3]int{10,20,30}

  prices[2] = 50
  fmt.Println(prices)
}
Result:

[10 20 50]


           // Array Initialization
If an array or one of its elements has not been initialized in the code, it is assigned the default value of its type.

Tip: The default value for int is 0, and the default value for string is "".

Example
package main
import ("fmt")

func main() {
  arr1 := [5]int{} //not initialized
  arr2 := [5]int{1,2} //partially initialized
  arr3 := [5]int{1,2,3,4,5} //fully initialized

  fmt.Println(arr1)
  fmt.Println(arr2)
  fmt.Println(arr3)
}
Result:

[0 0 0 0 0]
[1 2 0 0 0]
[1 2 3 4 5]



                 // Initialize Only Specific Elements
It is possible to initialize only specific elements in an array.

Example
This example initializes only the second and third elements of the array:

package main
import ("fmt")

func main() {
  arr1 := [5]int{1:10,2:40}

  fmt.Println(arr1)
}
Result:

[0 10 40 0 0]
Example Explained:
 The array above has 5 elements.

1:10 means: assign 10 to array index 1 (second element).
2:40 means: assign 40 to array index 2 (third element).



                         // Find the Length of an Array
The len() function is used to find the length of an array:

Example
package main
import ("fmt")

func main() {
  arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  arr2 := [...]int{1,2,3,4,5,6}

  fmt.Println(len(arr1))
  fmt.Println(len(arr2))
}
Result:

4
6

*/
