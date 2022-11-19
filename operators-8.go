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