/*
// https://leetcode.com/problems/contains-duplicate/
// https://www.youtube.com/watch?v=o1SHWzNTzMs&list=PLEWNW6mPFo-BdAGVEo_ATdxjujNXL_gB6

 Go Constants
If a variable should have a fixed value that cannot be changed, you can use the const keyword.

The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.

Syntax:
const CONSTNAME type = value

Note: The value of a constant must be assigned when you declare it.



                             Constant Rules
Constant names follow the same naming rules as variables
Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
Constants can be declared both inside and outside of a function

                     Constant Types
There are two types of constants:
1.Typed constants
2.Untyped constants

Note: In this case, the type of the constant is inferred from the
(means the compiler decides the type of the constant, based on the value).


          Constants: Unchangeable and Read-only
When a constant is declared, it is not possible to change the value later:
  const A = 1
  A = 2 // not working
  fmt.Println(A) // error this code


*/

package main

import (
	"fmt"
)

// Multiple Constants Declaration
// Multiple constants can be grouped together into a block for readability:
const (
	A1 int = 1
	B1     = 3.14
	C1     = "Hi!"
)

// Typed Constants
// Typed constants are declared with a defined type:
const A int = 1

// Untyped Constants
// Untyped constants are declared without a type:
const PI = 3.1416

func main() {

	fmt.Println(A1) // 1
	fmt.Println(B1) // 3.14
	fmt.Println(C1) // Hi

	fmt.Println(A)  // 1
	fmt.Println(PI) // 3.1416

	const GFG = "GeeksforGeeks"
	fmt.Println("Hello", GFG)

	fmt.Println("Happy", PI, "Day")

	const Correct = true
	fmt.Println("Go rules?", Correct)

	const untypedInteger = 123
	const untypedFloating = 123.12

	const typedInteger int = 123
	const typedFloatingPoint float64 = 123.12

	fmt.Println(untypedInteger, untypedFloating)
}

/*
                              Constants এর বৈশিষ্ট্যসমূহ:
স্থির মান: একবার একটি constant ডিফাইন করলে, তার মান পরিবর্তন হয় না।
যেমন const Pi = 3.14 একবার সেট হলে, তা সব সময় ৩.১৪ থাকবে।

টাইপ ইনফারেন্স: Constants গুলি যে কোনো টাইপের হতে পারে, যেমন সংখ্যা, স্ট্রিং,
বা বুলিয়ান। যদি টাইপ স্পষ্ট না করা হয়, তবে Go স্বয়ংক্রিয়ভাবে টাইপ নির্ধারণ করে।

মেমরি অ্যালোকেশন নেই: Constants মেমরিতে সংরক্ষিত হয় না। প্রোগ্রাম কম্পাইল
হওয়ার সময় এগুলি সরাসরি তাদের মান দিয়ে প্রতিস্থাপিত হয়।

কম্পাইলেশন-টাইম মূল্যায়ন: Constants গুলি প্রোগ্রাম রান টাইমে পরিবর্তন হয় না, এবং
কম্পাইলার এগুলি কম্পাইলেশন সময়ে মূল্যায়ন করে।

এভাবে, constants ব্যবহার করে আপনি প্রোগ্রামে এমন মান সংরক্ষণ করতে পারেন যা কখনো পরিবর্তন হবে না।

*/
