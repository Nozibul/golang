/*
                  // Go Variable Types
In Go, there are different types of variables, for example:

int- stores integers (whole numbers), such as 123 or -123
float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
float32 এর তুলনায় float64 বেশি নির্ভুলতা দেয় এবং বৃহৎ সংখ্যার ক্ষেত্রে কম ত্রুটি হয়।
এটি 64-বিট ব্যবহার করে, যার ফলে দশমিকের পরে আরও বেশি সংখ্যা সঠিকভাবে সংরক্ষণ করা যায়।

উদাহরণ:
3.141592653589793 (π এর সুনির্দিষ্ট মান)
-273.15 (শূন্য কেলভিন)

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

Note: In this case, the type of the variable is inferred from the
value (means that the compiler decides the type of the variable, based on the value).
দ্রষ্টব্য: এই ক্ষেত্রে, ভেরিয়েবলের ধরনটি মান থেকে অনুমান করা হয় (মানে কম্পাইলার মানের উপর ভিত্তি করে পরিবর্তনশীলের প্রকার নির্ধারণ করে)।

Note: It is not possible to declare a variable using :=, without assigning a value to it.
দ্রষ্টব্য: এটিতে একটি মান নির্ধারণ না করে := ব্যবহার করে একটি ভেরিয়েবল ঘোষণা করা সম্ভব নয়।

*/

package main

import (
	"fmt"
	"math"
)

var a int
var e int = 2
var f = 3

func main() {

	//1. Variable Declaration With Initial Value

	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred means অনুমান করা
	/*
		🔹 এখানে student2 নামে একটি ভেরিয়েবল ডিক্লেয়ার করা হয়েছে।
		🔹 var কীওয়ার্ড ব্যবহার করা হয়েছে, কিন্তু টাইপ (string) সরাসরি উল্লেখ করা হয়নি।
		🔹 Go স্বয়ংক্রিয়ভাবে টাইপ "Jane" দেখে বুঝতে পারে যে এটি string।
		🔹 একে বলে type inference বা টাইপ অনুমান করা।
	*/
	x := 2 //type is inferred
	/*
	   🔹 এখানে x নামে একটি ভেরিয়েবল ডিক্লেয়ার করা হয়েছে।
	   🔹 := অপারেটর ব্যবহার করে ভেরিয়েবল ডিক্লেয়ার এবং ইনিশিয়ালাইজ করা হয়েছে।
	   🔹 2 দেখে Go বুঝতে পারে যে x এর টাইপ int হবে।
	   🔹 এটি short variable declaration পদ্ধতি, যা শুধুমাত্র function এর ভিতরে ব্যবহার করা যায়।
	*/

	/*

		1.

		go
		Copy code
		var student1 string = "John" //type is string
		🔹 এখানে student1 নামে একটি ভেরিয়েবল ডিক্লেয়ার করা হয়েছে।
		🔹 var কীওয়ার্ড ব্যবহার করে স্পষ্টভাবে string টাইপ উল্লেখ করা হয়েছে।
		🔹 ভেরিয়েবলের মান "John" সেট করা হয়েছে, এবং এটি string টাইপের কারণ এটি স্পষ্টভাবে লেখা আছে।

		2.

		go
		Copy code
		var student2 = "Jane"        //type is inferred means অনুমান করা
		🔹 এখানে student2 নামে একটি ভেরিয়েবল ডিক্লেয়ার করা হয়েছে।
		🔹 var কীওয়ার্ড ব্যবহার করা হয়েছে, কিন্তু টাইপ (string) সরাসরি উল্লেখ করা হয়নি।
		🔹 Go স্বয়ংক্রিয়ভাবে টাইপ "Jane" দেখে বুঝতে পারে যে এটি string।
		🔹 একে বলে type inference বা টাইপ অনুমান করা।

		3.

		go
		Copy code
		x := 2                       //type is inferred
		🔹 এখানে x নামে একটি ভেরিয়েবল ডিক্লেয়ার করা হয়েছে।
		🔹 := অপারেটর ব্যবহার করে ভেরিয়েবল ডিক্লেয়ার এবং ইনিশিয়ালাইজ করা হয়েছে।
		🔹 2 দেখে Go বুঝতে পারে যে x এর টাইপ int হবে।
		🔹 এটি short variable declaration পদ্ধতি, যা শুধুমাত্র function এর ভিতরে ব্যবহার করা যায়।

		সারাংশ:

		var ব্যবহারে টাইপ উল্লেখ করা ঐচ্ছিক, কারণ Go টাইপ অনুমান করতে পারে।
		:= ব্যবহার করলে Go টাইপ স্বয়ংক্রিয়ভাবে নির্ধারণ করে।
	*/

	fmt.Println(student1) // John
	fmt.Println(student2) // Jane
	fmt.Println(x)        // 2

	//2. Variable Declaration Without Initial Value
	var g string
	var b int
	var c bool

	fmt.Println(a) // empty string
	fmt.Println(b) // 0
	fmt.Println(c) // false

	//3. Value Assignment After Declaration
	//It is possible to assign a value to a variable after it is declared.
	//This is helpful for cases the value is not initially known.

	//  Note: It is not possible to declare a variable using ":=" without assigning a value to it.

	var student4 string
	student4 = "John"
	fmt.Println(student4) // John

	// This example shows declaring variables outside of a function, with the var keyword:

	//  g = 1;
	fmt.Println(g) // 1
	fmt.Println(e) // 2
	fmt.Println(f) // 3

	// Since := is used outside of a function, running the program results in an error.

	/*                       // Go Multiple Variable Declaration
	In Go, it is possible to declare multiple variables in the same line.
	*/
	var s, t, p, q int = 1, 3, 5, 7

	fmt.Println(s) // 1
	fmt.Println(t) // 2
	fmt.Println(p) // 3
	fmt.Println(q) // 4

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
		ages   int             // ages ভেরিয়েবল, যার টাইপ int (ইনিশিয়াল মান নেই, তাই 0 হবে)
		number int    = 1      // number ভেরিয়েবল, টাইপ int এবং মান 1
		name   string = "Jhon" // name ভেরিয়েবল, টাইপ string এবং মান "hello"
	)

	fmt.Println(ages)   // 0
	fmt.Println(number) // 1
	fmt.Println(name)   // Jhon

	// A variable can be assigned to any value of its type.
	// In the above program, age can be assigned any integer value.

	var age int                                           // variable declaration
	fmt.Println("My initial age is", age)                 // 0
	age = 30                                              //assignment
	fmt.Println("My age after first assignment is", age)  //30
	age = 45                                              //assignment
	fmt.Println("My age after second assignment is", age) // 45

	//Short hand declaration

	count := 10
	numbers, userName := 45, "John"
	fmt.Println(count, numbers, userName)

	/*
	  এই প্রোগ্রামে একটি ত্রুটি হবে কারণ short hand declaration (:=) ব্যবহৃত হয়েছে,
	  কিন্তু সমস্ত ভেরিয়েবলের জন্য প্রাথমিক মান (initial value) প্রদান করা হয়নি।
	*/
	brotherName, num := "Naveen" //error

	fmt.Println("my name is", brotherName, "age is", num)
	/*
	   := অপারেটরের মাধ্যমে একসাথে ভেরিয়েবল ডিক্লেয়ার এবং ইনিশিয়ালাইজ করা হয়,
	   এবং যখন একাধিক ভেরিয়েবল ডিক্লেয়ার করা হয়, তখন তাদের প্রত্যেকটির জন্য
	   প্রাথমিক মান (initial value) দিতে হয়।
	*/

	a, b := 20, 30                            // declare variables a and b
	fmt.Println("a is", a, "b is", b)         // a is 20 b is 30
	b, c := 40, 50                            // b is already declared but c is new
	fmt.Println("b is", b, "c is", c)         // b is 40 c is 50
	b, c = 80, 90                             // assign new values to already declared variables b and c
	fmt.Println("changed b is", b, "c is", c) // changed b is 80 c is 90

	// Whereas if we run the program below,
	a, b := 20, 30 //a and b declared
	fmt.Println("a is", a, "b is", b)
	a, b := 40, 50 //error, no new variables

	//Variables can also be assigned values which are computed during run
	// time. Consider the following program,

	y, z := 12.3, 3.4
	w := math.Min(y, z)
	fmt.Println(" value is", w) // value is 15.7

	ages := 65      // age is int
	ages = "Naveen" // error since we are trying to assign a string to a variable of type int

}

/*
Go-তে := অপারেটরটি Short Variable Declaration অপারেটর নামে পরিচিত।

🔹 := অপারেটরের বৈশিষ্ট্য:
একসাথে ভেরিয়েবল ডিক্লেয়ার এবং ইনিশিয়ালাইজ করে।
টাইপ উল্লেখ করতে হয় না, কারণ Go স্বয়ংক্রিয়ভাবে টাইপ অনুমান (type inference) করে।
শুধুমাত্র function এর ভিতরে ব্যবহার করা যায়।
*/

// Type inference
/*
যদি একটি ভেরিয়েবলের একটি প্রাথমিক মান (initial value) থাকে, তবে Go সেই মান
দেখে নিজেই ভেরিয়েবলের টাইপ বুঝে ফেলে। এর মানে হলো, যদি ভেরিয়েবলে আগে
থেকেই মান দেওয়া থাকে, তাহলে টাইপ উল্লেখ করার দরকার নেই। Go তা অটো বুঝে নেয়।
*/

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

/*
In Go, variable names generally follow camelCase for local variables and PascalCase for exported (public) variables.

1. CamelCase (Recommended for local variables)
Used for variables that are not exported outside the package.
Example:

userName := "John"
totalAmount := 500
isConnected := true


2. PascalCase (For Exported/Public Variables)
Used for variables that need to be accessed from other packages.
Example:

var ServerPort = 8080
var MaxUsers = 100


3. Snake_case (Rarely Used in Go)
Not common in Go, but sometimes used in constants.
Example:
go
Copy code
const max_limit = 100


🔹 Best Practices for Naming Variables in Go
Keep it short and meaningful
❌ var x int → ✅ var count int

Avoid unnecessary abbreviations
❌ var usrNm string → ✅ var userName string

Boolean variables should start with is or has
✅ isAvailable
✅ hasPermission

Constants in UPPERCASE (optional)
✅ const MAX_SIZE = 500

Idiomatic Go encourages short variable names for small scopes
✅ i, j, k (for loops or counters)

Would you like examples of exported functions or constants with proper naming conventions?

*/
