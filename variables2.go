/*
                  // Go Variable Types
In Go, there are different types of variables, for example:

int- stores integers (whole numbers), such as 123 or -123
float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
string - stores text, such as "Hello World". String values are surrounded by double quotes
bool- stores values with two states: true or false


                      // Declaring (Creating) Variables
In Go, there are two ways to declare a variable:
1. With the var keyword:  
 Example: 
 var variablename type := value ;
 var newSrt string := "Go developer"

 Note: You always have to specify either type or value (or both).

2. With the := sign:
Use the := sign, followed by the variable value:

Syntax
variablename := value

Note: In this case, the type of the variable is inferred from the value (means that the compiler decides the type of the variable, based on the value).
দ্রষ্টব্য: এই ক্ষেত্রে, ভেরিয়েবলের ধরনটি মান থেকে অনুমান করা হয় (মানে কম্পাইলার মানের উপর ভিত্তি করে পরিবর্তনশীলের প্রকার নির্ধারণ করে)।

Note: It is not possible to declare a variable using :=, without assigning a value to it.
দ্রষ্টব্য: এটিতে একটি মান নির্ধারণ না করে := ব্যবহার করে একটি ভেরিয়েবল ঘোষণা করা সম্ভব নয়।

*/

package main
import "fmt"

var a int
var e int = 2
var f = 3

func main() {

   //1. Variable Declaration With Initial Value

   var student1 string = "John" //type is string
   var student2 = "Jane" //type is inferred means অনুমান করা
   x := 2 //type is inferred
 
   fmt.Println(student1) // John
   fmt.Println(student2) // Jane
   fmt.Println(x) // 2

 
   //2. Variable Declaration Without Initial Value
   var g string
   var b int
   var c bool

   fmt.Println(a) ;// empty string
   fmt.Println(b) ;// 0
   fmt.Println(c) ;// false


   //3. Value Assignment After Declaration
   //It is possible to assign a value to a variable after it is declared. 
   //This is helpful for cases the value is not initially known.
   /*
    একটি ভেরিয়েবল ঘোষণা করার পরে এটি একটি মান নির্ধারণ করা সম্ভব।
    এটি এমন ক্ষেত্রে সহায়ক যেটির মান প্রাথমিকভাবে জানা যায় না। 
    */

   //  Note: It is not possible to declare a variable using ":=" without assigning a value to it.

    var student4 string
    student4 = "John"
    fmt.Println(student4); // John


    // This example shows declaring variables outside of a function, with the var keyword:

   //  g = 1;
    fmt.Println(g) ;// 1
    fmt.Println(e) ;// 2
    fmt.Println(f) ;// 3

   // Since := is used outside of a function, running the program results in an error.



   /*                       // Go Multiple Variable Declaration
    In Go, it is possible to declare multiple variables in the same line.  
   */
   var s, t, p, q int = 1, 3, 5, 7

  fmt.Println(s) ; // 1
  fmt.Println(t) ; // 2
  fmt.Println(p) ; // 3
  fmt.Println(q) ; // 4



// Note: If you use the type keyword, it is only possible to declare one type of variable per line.
// দ্রষ্টব্য: আপনি যদি টাইপ কীওয়ার্ড ব্যবহার করেন, তবে প্রতি লাইনে শুধুমাত্র এক ধরনের ভেরিয়েবল ঘোষণা করা সম্ভব।

var aa, bb = 6, "Hello"
  cc, dd := 7, "World!"

  fmt.Println(aa) // 6
  fmt.Println(bb) // Hello
  fmt.Println(cc) // 7
  fmt.Println(dd) // World




/*
                    // Go Variable Declaration in a Block
Multiple variable declarations can also be grouped together into a block for greater readability: 
*/
var (
   a1 int
   b1 int = 1
   c1 string = "hello"
 )

fmt.Println(a1); // 0
fmt.Println(b1); // 1
fmt.Println(c1); // hello



}


/*
                       Go Variable Naming Rules
A variable can have a short name (like x and y) or a more descriptive name (age, price, carname, etc.).

Go variable naming rules:

A variable name must start with a letter or an underscore character (_)
A variable name cannot start with a digit
A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
Variable names are case-sensitive (age, Age and AGE are three different variables)
There is no limit on the length of the variable name
A variable name cannot contain spaces
The variable name cannot be any Go keywords 


                   Multi-Word Variable Names
Variable names with more than one word can be difficult to read.

There are several techniques you can use to make them more readable:

                      // Camel Case
Each word, except the first, starts with a capital letter:
myVariableName = "John"


                     // Pascal Case
Each word starts with a capital letter:
MyVariableName = "John"

 
                     // Snake Case
Each word is separated by an underscore character:
my_variable_name = "John"

 
*/

 


                //Difference Between var and :=


/*
                 var	                                                 :=
Can be used inside and outside of functions	         Can only be used inside functions
// ফাংশন ভিতরে এবং বাইরে ব্যবহার করা যেতে পারে           // ফাংশন ভিতরে ব্যবহার করা যেতে পারে 
Variable declaration and value assignment             Variable declaration and value assignment  
can be done separately 	                              cannot be done separately (must be done in the same line)
*/                