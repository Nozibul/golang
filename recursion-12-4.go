/*
                    // Go Recursion Functions
Recursion Functions
Go accepts recursion functions. A function is recursive if it calls 
itself and reaches a stop condition.

In the following example, testcount() is a function that calls itself. 
We use the x variable as the data, which increments with 1 (x + 1) 
every time we recurse. The recursion ends when the x variable equals
to 11 (x == 11). 

Example
package main
import ("fmt")

func testcount(x int) int {
  if x == 6 {
    return 0
  }
  fmt.Println(x)
  return testcount(x + 1)
}

func main(){
  testcount(1)
}
Result: 1 2 3 4 5





Recursion is a common mathematical and programming concept. 
This has the benefit of meaning that you can loop through data to
reach a result.

The developer should be careful with recursion functions as it 
can be quite easy to slip into writing a function which never terminates,
or one that uses excess amounts of memory or processor power. 
However, when written correctly recursion can be a very efficient and 
mathematically-elegant approach to programming.

In the following example, factorial_recursion() is a function that 
calls itself. We use the x variable as the data, which decrements (-1) 
every time we recurse. The recursion ends when the condition is not 
greater than 0 (i.e. when it is 0).


Example : 
package main
import ("fmt")

func factorial_recursion(x float64) (y float64) {
  if x > 0 {
     y = x * factorial_recursion(x-1)
  } else {
     y = 1
  }
  return
}

func main() {
  fmt.Println(factorial_recursion(4))
}
Result: 24
*/