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

                                // Misc Operators

/* ধরা যাক, আপনার কাছে একটি সাধারণ ভেরিয়েবল x আছে যার মান 10। 
আপনি যদি জানতে চান, x ভেরিয়েবলের ঠিক কোন স্থানে (অথবা কোন অ্যাড্রেসে) 
মান 10 রাখা আছে, তাহলে আপনি & অপারেটর ব্যবহার করবেন।

& অপারেটরের কাজ:
& অপারেটরটি দিয়ে আপনি একটি ভেরিয়েবলের অ্যাড্রেস (অথবা পয়েন্টার) পাবেন। 
এটি আসলে ভেরিয়েবলের অবস্থান বা জায়গা সম্পর্কে তথ্য দেয়, যা পয়েন্টার 
হিসেবে কাজ করতে পারে।

উদাহরণ:
var x int = 10  // x ভেরিয়েবল, যার মান 10
var p *int = &x // p হলো পয়েন্টার, যা x এর অ্যাড্রেস ধারণ করছে

এখানে:
1. &x ব্যবহার করে আপনি x ভেরিয়েবলের অ্যাড্রেস বা স্থান পাবেন।
2. p এখন x এর অ্যাড্রেস ধারণ করে। অর্থাৎ, p এ x এর স্থান সংরক্ষিত থাকবে, কিন্তু p তে আসল মান (10) থাকবে না। p তে x এর অবস্থান বা অ্যাড্রেস থাকবে।
এভাবে, & অপারেটরটি ভেরিয়েবলের অবস্থান (অথবা অ্যাড্রেস) প্রদান করে, যা পরে পয়েন্টার হিসেবে ব্যবহার করা হয়। */

/*  


              স্টার অপারেটর (*) কীভাবে কাজ করে, তা সহজভাবে বোঝাতে:

ধরা যাক, আমাদের একটি ভেরিয়েবল x আছে যার মান 10। আমরা চাই সেই 
ভেরিয়েবলের মানটি একটি পয়েন্টারের মাধ্যমে অ্যাক্সেস করতে।

এখানে পয়েন্টার হলো এমন একটি ভেরিয়েবল যা অন্য একটি ভেরিয়েবলের অ্যাড্রেস 
(অথবা অবস্থান) ধারণ করে।

স্টার অপারেটর (*) এর কাজ:
1. & অপারেটর দিয়ে আমরা একটি ভেরিয়েবলের অ্যাড্রেস (অথবা লোকেশন) পাই।
2. তারপর * অপারেটর দিয়ে সেই অ্যাড্রেস থেকে আসল মানটি নিতে পারি।
*/

var x int = 10  // x নামের একটি ভেরিয়েবল, যার মান 10
var p *int = &x // p হলো একটি পয়েন্টার যা x এর অ্যাড্রেস ধারণ করছে

fmt.Println(*p)  // p এর মাধ্যমে x এর আসল মান পাওয়া যাবে, অর্থাৎ 10

/*  
এখানে:

&x দিয়ে আমরা x এর অ্যাড্রেস পেয়ে p পয়েন্টারকে সেট করেছি।
পরে, *p দিয়ে আমরা p পয়েন্টারের মাধ্যমে x এর আসল মান (যা 10) দেখতে পাচ্ছি।
স্টার অপারেটর * এমনভাবে কাজ করে যে এটি পয়েন্টারের মাধ্যমে ঐ অ্যাড্রেসে থাকা আসল মানটি আমাদের দেয়।

*/

/* 
গো ভাষায় <- (রিসিভ) অপারেটরটি চ্যানেল থেকে মান গ্রহণ করতে ব্যবহৃত হয়। 
এটি একটি চ্যানেলের মাধ্যমে মান পাঠানো এবং গ্রহণ করার কাজ করতে ব্যবহৃত হয়।
এই অপারেটরটি সাধারণত গোরউটিনের মধ্যে ব্যবহৃত হয়, যেখানে এক বা একাধিক গোরউটিন (concurrent execution) থেকে চ্যানেল মাধ্যমে মান পাঠানো এবং গ্রহণ করা হয়।

রিসিভ অপারেটর কাজ কীভাবে করে:
চ্যানেল একটি যোগাযোগ মাধ্যম হিসেবে কাজ করে যেখানে এক গোরউটিন থেকে 
অন্য গোরউটিনে ডেটা পাঠানো বা গ্রহণ করা যায়।
<- অপারেটরটি চ্যানেল থেকে মান গ্রহণ করতে ব্যবহৃত হয়। এটি চ্যানেল থেকে 
একটি মান নিয়ে আসে।

*/

// উদাহরণ:

ch := make(chan int)  // একটি নতুন চ্যানেল তৈরি করা হয়েছে (চ্যানেলটি int টাইপের)

go func() {            // একটি নতুন গোরউটিন শুরু করা
    ch <- 42            // চ্যানেলে 42 পাঠানো হচ্ছে
}()

value := <-ch          // চ্যানেল থেকে মান গ্রহণ করা হচ্ছে
fmt.Println(value)     // 42 প্রদর্শিত হবে

/* 
এখানে কী হচ্ছে:

1. প্রথমে ch := make(chan int) দিয়ে একটি নতুন চ্যানেল তৈরি করা হচ্ছে। 
এই চ্যানেলটি int ধরনের মান ধারণ করবে।

2. তারপর একটি গোরউটিন go func() {...} শুরু করা হচ্ছে। এই গোরউটিনের 
মধ্যে ch <- 42 এর মাধ্যমে চ্যানেলে মান 42 পাঠানো হচ্ছে।

3. মূল গোরউটিনে, value := <-ch ব্যবহার করে আমরা চ্যানেল থেকে মান গ্রহণ করছি। 
যখন ch থেকে মান আসবে (এখানে 42), তখন তা value ভেরিয়েবলে রাখা হবে।

4. শেষে fmt.Println(value) দিয়ে value এর মান (42) প্রদর্শন করা হবে।

                    গুরুত্বপূর্ণ বিষয়:
<-ch রিসিভ অপারেটরটি চ্যানেল থেকে মান গ্রহণ করতে অপেক্ষা করে যতক্ষণ না কোনো 
মান চ্যানেলে পাঠানো হয়। যদি চ্যানেলে মান না থাকে, তবে এটি ব্লক হয়ে যাবে, 
অর্থাৎ চ্যানেলে মান আসা না পর্যন্ত কোডটি চলতে থাকবে না।
*/

/*
<- (রিসিভ) অপারেটর চ্যানেল থেকে মান গ্রহণ করতে ব্যবহৃত হয় এবং এটি গোরউটিনের 
মধ্যে সমান্তরালভাবে কাজ করে। এটি রিয়েল-টাইম ডাটা প্রসেসিং, চ্যাট সিস্টেম, স্টক ট্রেডিং, 
ডাটাবেস রিডিং, গেম ডেভেলপমেন্ট, এবং ফাইল প্রসেসিং ইত্যাদি ক্ষেত্রে ব্যাপকভাবে ব্যবহৃত 
হতে পারে।  
*/