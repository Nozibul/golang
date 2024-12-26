/*
              // Go Operators
Operators are used to perform operations on variables and values.

The + operator adds together two values, like in the example below:

Example
package main
import ("fmt")

func main() {
  var a = 15 + 25
  fmt.Println(a)
}

Although the + operator is often used to add
together two values, it can also be used to add together a
variable and a value, or a variable and another variable:

Example
package main
import ("fmt")

func main() {
  var (
    sum1 = 100 + 50 // 150 (100 + 50)
    sum2 = sum1 + 250 // 400 (150 + 250)
    sum3 = sum2 + sum2 // 800 (400 + 400)
  )
  fmt.Println(sum3)
}


*/

/*
                    // Arithmetic Operators
Arithmetic operators are used to perform common mathematical operations.

Operator	Name	Description	Example	Try it
+	Addition	Adds together two values	x + y
-	Subtraction	Subtracts one value from another	x - y
*	Multiplication	Multiplies two values	x * y
/	Division	Divides one value by another	x / y
%	Modulus	Returns the division remainder	x % y
++	Increment	Increases the value of a variable by 1	x++
--	Decrement	Decreases the value of a variable by 1	x--


        // Assignment Operators
Assignment operators are used to assign values to variables.

In the example below, we use the assignment operator (=) to assign the value 10 to a variable called x:

Example
package main
import ("fmt")

func main() {
  var x = 10
  fmt.Println(x)
}


 // The addition assignment operator (+=) adds a value to a variable:

Example
package main
import ("fmt")

func main() {
  var x = 10
  x +=5
  fmt.Println(x)
}




 list of all assignment operators:

Operator	Example	Same As	Try it
=	x = 5	x = 5
+=	x += 3	x = x + 3
-=	x -= 3	x = x - 3
*=	x *= 3	x = x * 3
/=	x /= 3	x = x / 3
%=	x %= 3	x = x % 3
&=	x &= 3	x = x & 3
|=	x |= 3	x = x | 3
^=	x ^= 3	x = x ^ 3
>>=	x >>= 3	x = x >> 3
<<=	x <<= 3	x = x << 3


*/

/*
Go Comparison Operators
Comparison Operators
Comparison operators are used to compare two values.

Note: The return value of a comparison is either true (1) or false (0).

In the following example, we use the greater than operator (>) to find out if 5 is greater than 3:

Example
package main
import ("fmt")

func main() {
  var x = 5
  var y = 3
  fmt.Println(x>y) // returns 1 (true) because 5 is greater than 3
}
A list of all comparison operators:

Operator	Name	Example	Try it
==	Equal to	x == y
!=	Not equal	x != y
>	Greater than	x > y
<	Less than	x < y
>=	Greater than or equal to	x >= y
<=	Less than or equal to	x <= y
*/

/*
                 Go Logical Operators
Logical Operators
Logical operators are used to determine the logic between variables or values:

Operator	Name	Description	Example	Try it
&& 	Logical and	Returns true if both statements are true	x < 5 &&  x < 10
|| 	Logical or	Returns true if one of the statements is true	x < 5 || x < 4
!	Logical not	Reverse the result, returns false if the result is true	!(x < 5 && x < 10)

*/

/*
               // Go Bitwise Operators
Bitwise Operators
Bitwise operators are used on (binary) numbers:

Operator	Name	Description	Example	Try it
& 	AND	Sets each bit to 1 if both bits are 1	x & y
|	OR	Sets each bit to 1 if one of two bits is 1	x | y
 ^	XOR	Sets each bit to 1 if only one of two bits is 1	x ^ b
<<	Zero fill left shift	Shift left by pushing zeros in from the right	x << 2
>>	Signed right shift	Shift right by pushing copies of the leftmost bit in from the left,
and let the rightmost bits fall off	x >> 2
*/

// Go program to illustrate the
// use of bitwise operators
package main

import "fmt"

func main() {
	p := 34
	q := 20

	// & (bitwise AND)
	result1 := p & q
	fmt.Printf("Result of p & q = %d", result1)

	// | (bitwise OR)
	result2 := p | q
	fmt.Printf("\nResult of p | q = %d", result2)

	// ^ (bitwise XOR)
	result3 := p ^ q
	fmt.Printf("\nResult of p ^ q = %d", result3)

	// << (left shift)
	result4 := p << 1
	fmt.Printf("\nResult of p << 1 = %d", result4)

	// >> (right shift)
	result5 := p >> 1
	fmt.Printf("\nResult of p >> 1 = %d", result5)

	// &^ (AND NOT)
	result6 := p &^ q
	fmt.Printf("\nResult of p &^ q = %d", result6)

}

/* Result of p & q = 0
Result of p | q = 54
Result of p ^ q = 54
Result of p << 1 = 68
Result of p >> 1 = 17
Result of p &^ q = 34 */

/*
                           ১. p & q (Bitwise AND)
result1 := p & q

p = 34 (00100010)
q = 20 (00010100)

   00100010  (34)
&  00010100  (20)
--------------
   00000000  (0)


                            ২. p | q (Bitwise OR)

result2 := p | q
OR অপারেশনে যদি কোনো একটি বিট 1 হয়, তাহলে ফলাফল 1 হয়।

   00100010  (34)
|  00010100  (20)
--------------
   00110110  (54)
ফলাফল: 54


                            ৩. p ^ q (Bitwise XOR)

result3 := p ^ q
XOR অপারেশনে উভয় বিট ভিন্ন হলে 1 হয়, একই হলে 0।
markdown
Copy code
   00100010  (34)
^  00010100  (20)
--------------
   00110110  (54)
ফলাফল: 54


                                 ৪. p << 1 (Left Shift)

result4 := p << 1
Left Shift মানে p-এর সব বিট এক ধাপ বামে সরানো হয় এবং ডান দিকে 0 যোগ করা হয়।

   00100010  (34)
<< 1
--------------
   01000100  (68)
ফলাফল: 68


                                ৫. p >> 1 (Right Shift)

result5 := p >> 1
Right Shift মানে p-এর সব বিট এক ধাপ ডানে সরানো হয় এবং বাম দিকে 0 যোগ করা হয়।

   00100010  (34)
>> 1
--------------
   00010001  (17)
ফলাফল: 17


                                ৬. p &^ q (AND NOT)

result6 := p &^ q
AND NOT (p &^ q) মানে p থেকে q-এর বিট 1 থাকলে সেটি বাদ দেওয়া হয়।

   00100010  (34)
&^ 00010100  (20)
--------------
   00100010  (34)

যেহেতু q-তে এমন কোনো বিট নেই যা p-তে 1 এবং q-তে 1, তাই p অপরিবর্তিত থাকে।
ফলাফল: 34


                            সংক্ষেপে ফলাফল:
p & q  = 0
p | q  = 54
p ^ q  = 54
p << 1 = 68
p >> 1 = 17
p &^ q = 34

এই প্রোগ্রামটি দেখায় কিভাবে Go প্রোগ্রামিং ভাষায় বিট লেভেল অপারেশন করে বিভিন্ন ফলাফল পাওয়া যায়।


*/
