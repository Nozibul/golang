/* 
               // Go Conditions
Conditional statements are used to perform different actions based on different conditions.

Go Conditions
A condition can be either true or false.

Go supports the usual comparison operators from mathematics:

Less than <
Less than or equal <=
Greater than >
Greater than or equal >=
Equal to ==
Not equal to !=
Additionally, Go supports the usual logical operators:

Logical AND &&
Logical OR ||
Logical NOT !
You can use these operators or their combinations to create conditions for different decisions.

Example	Try it
x > y	
x != y	
(x > y) && (y > z)	
(x == y) || z	


Go has the following conditional statements:

Use if to specify a block of code to be executed, if a specified condition is true
Use else to specify a block of code to be executed, if the same condition is false
Use else if to specify a new condition to test, if the first condition is false
Use switch to specify many alternative blocks of code to be executed
*/

/*

              // Go if statement
The if Statement
Use the if statement to specify a block of Go code to be executed if a condition is true.

Syntax
if condition {
  // code to be executed if condition is true
}
Note that if is in lowercase letters. Uppercase letters (If or IF) will generate an error.

In the example below, we test two values to find out if 20 is greater than 18. If the condition is true, print some text:

Example
package main
import ("fmt")

func main() {
  if 20 > 18 {
    fmt.Println("20 is greater than 18")
  }
}
We can also test variables:

Example
package main
import ("fmt")

func main() {
  x:= 20
  y:= 18
  if x > y {
    fmt.Println("x is greater than y")
  }
}
Example explained
In the example above we use two variables, x and y, to test whether x is greater than y (using the > operator). As x is 20, and y is 18, and we know that 20 is greater than 18, we print to the screen that "x is greater than y".


*/


/* 
                 // Go if else Statement
The else Statement
Use the else statement to specify a block of code to be executed if the condition is false.

Syntax
if condition {
  // code to be executed if condition is true
} else {
  // code to be executed if condition is false
}
Using The if else Statement
Example
In this example, time (20) is greater than 18, so the if condition is false. Because of this, we move on to the else condition and print to the screen "Good evening". If the time was less than 18, the program would print "Good day":

package main
import ("fmt")

func main() {
  time := 20
  if (time < 18) {
    fmt.Println("Good day.")
  } else {
    fmt.Println("Good evening.")
  }
}



Example
In this example, the temperature is 14 so the condition for if is false so the code line inside the else statement is executed:

package main
import ("fmt")

func main() {
  temperature := 14
  if (temperature > 15) {
    fmt.Println("It is warm out there")
  } else {
    fmt.Println("It is cold out there")
  }
}
The brackets in the else statement should be like } else {:

Example
Having the else brackets in a different line will raise an error:

package main
import ("fmt")

func main() {
  temperature := 14
  if (temperature > 15) {
    fmt.Println("It is warm out there.")
  } // this raises an error // else ta khane hobe
  else {
    fmt.Println("It is cold out there.")
  }
}
Result:

./prog.go:9:3: syntax error: unexpected else, expecting }

*/

/* 
Go else if Statement
The else if Statement
Use the else if statement to specify a new condition if the first condition is false.

Syntax
if condition1 {
   // code to be executed if condition1 is true
} else if condition2 {
   // code to be executed if condition1 is false and condition2 is true
} else {
   // code to be executed if condition1 and condition2 are both false
}
Using The else if Statement
Example:
This example shows how to use an else if statement.

package main
import ("fmt")

func main() {
  time := 22
  if time < 10 {
    fmt.Println("Good morning.")
  } else if time < 20 {
    fmt.Println("Good day.")
  } else {
    fmt.Println("Good evening.")
  }
}
Result:

Good evening.


Another example for the use of else if.

package main
import ("fmt")

func main() {
  a := 14
  b := 14
  if a < b {
    fmt.Println("a is less than b.")
  } else if a > b {
    fmt.Println("a is more than b.")
  } else {
    fmt.Println("a and b are equal.")
  }
}
Result:

a and b are equal.
*/

/*
Go Nested if Statement
The Nested if Statement
You can have if statements inside if statements, this is called a nested if.

Syntax
if condition1 {
   // code to be executed if condition1 is true
  if condition2 {
     // code to be executed if both condition1 and condition2 are true
  }
}
Example
This example shows how to use nested if statements:

package main
import ("fmt")

func main() {
  num := 20
  if num >= 10 {
    fmt.Println("Num is more than 10.")
    if num > 15 {
      fmt.Println("Num is also more than 15.")
     }
  } else {
    fmt.Println("Num is less than 10.")
  }
}
Result:

Num is more than 10.
Num is also more than 15.

*/