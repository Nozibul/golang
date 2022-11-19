/*
                   // Go For Loops
The for loop loops through a block of code a specified number of times.
The for loop is the only loop available in Go.

                   // Go for Loop
Loops are handy if you want to run the same code over and over again, 
each time with a different value.
আপনি যদি একই কোড বারবার চালাতে চান, প্রতিবার আলাদা মান সহ লুপগুলি কার্যকর।

Each execution of a loop is called an iteration.

The for loop can take up to three statements:

        //   Syntax :
for statement1; statement2; statement3 {
   // code to be executed for each iteration
}
statement1 Initializes the loop counter value.

statement2 Evaluated for each loop iteration. If it evaluates to TRUE, 
the loop continues. If it evaluates to FALSE, the loop ends.

statement3 Increases the loop counter value.

Note: These statements don't need to be present as loops arguments. 
However, they need to be present in the code in some form.




          // for Loop Examples
          // Example 1
This example will print the numbers from 0 to 4:  

package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    fmt.Println(i)
  }
}
Result: 0 1 2 3 4

Example 1 explained
i:=0; - Initialize the loop counter (i), and set the start value to 0
i < 5; - Continue the loop as long as i is less than 5
i++ - Increase the loop counter value by 1 for each iteration



       Example 2 :
This example counts to 100 by tens: 

package main
import ("fmt")

func main() {
  for i:=0; i <= 100; i+=10 {
    fmt.Println(i)
  }
}

i:=0; - Initialize the loop counter (i), and set the start value to 0
i <= 100; - Continue the loop as long as i is less than or equal to 100
i+=10 - Increase the loop counter value by 10 for each iteration




            // The continue Statement
The continue statement is used to skip one or more iterations in the loop. It then continues with the next iteration in the loop.

Example
This example skips the value of 3:

package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    if i == 3 {
      continue
    }
   fmt.Println(i)
  }
}
result: 0 1 2 4



                // The break Statement
The break statement is used to break/terminate the loop execution.

Example
This example breaks out of the loop when i is equal to 3:

package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    if i == 3 {
      break
    }
   fmt.Println(i)
  }
}
Result: 0 1 2

Note: continue and break are usually used with conditions.




                     // Nested Loops
It is possible to place a loop inside another loop.

Here, the "inner loop" will be executed one time for each iteration of the "outer loop":

Example
package main
import ("fmt")

func main() {
  adj := [2]string{"big", "tasty"}
  fruits := [3]string{"apple", "orange", "banana"}
  for i:=0; i < len(adj); i++ {
    for j:=0; j < len(fruits); j++ {
      fmt.Println(adj[i],fruits[j])
    }
  }
}
Result:
big apple
big orange
big banana
tasty apple
tasty orange
tasty banana




              // The Range Keyword
The range keyword is used to more easily iterate over an array, slice or map. It returns both the index and the value.

The range keyword is used like this:

Syntax
for index, value := array|slice|map {
   // code to be executed for each iteration
}



                // Example
This example uses range to iterate over an array and print both the indexes and the values at each (idx stores the index, val stores the value):

package main
import ("fmt")

func main() {
  fruits := [3]string{"apple", "orange", "banana"}
  for idx, val := range fruits {
     fmt.Printf("%v\t%v\n", idx, val)
  }
}
Result:
0      apple
1      orange
2      banana


Tip: To only show the value or the index,
you can omit the other output using an underscore (_).
Example :
Here, we want to omit the indexes (idx stores the index, val stores the value):

package main
import ("fmt")

func main() {
  fruits := [3]string{"apple", "orange", "banana"}
  for _, val := range fruits {
     fmt.Printf("%v\n", val)
  }
}
Result:

apple
orange
banana

*/