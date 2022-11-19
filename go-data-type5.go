/* 
                         Go Data Types
Data type is an important concept in programming. Data type specifies the size and type of variable values.

Go is statically typed, meaning that once a variable type is defined, it can only store data of that type.

       // Go has three basic data types:

bool: represents a boolean value and is either true or false
Numeric: represents integer types, floating point values, and complex types
string: represents a string value
*/


// Boolean Data Type
// A boolean data type is declared with the bool keyword and can only take the values true or false.
// The default value of a boolean data type is false.

/* package main
import ("fmt")

func main() {
  var b1 bool = true // typed declaration with initial value
  var b2 = true // untyped declaration with initial value
  var b3 bool // typed declaration without initial value
  b4 := true // untyped declaration with initial value

  fmt.Println(b1) // Returns true
  fmt.Println(b2) // Returns true
  fmt.Println(b3) // Returns false
  fmt.Println(b4) // Returns true
} */


/*
           // Go Integer Data Types
Integer data types are used to store a whole number without decimals, like 35, -50, or 1345000.

The integer data type has two categories:

Signed integers - can store both positive and negative values
Unsigned integers - can only store non-negative values

Tip: The default type for integer is int. If you do not specify a type, the type will be int.  


          // Go has five keywords/types of signed integers:

Type	     Size	                       Range
int	    Depends on platform:           
        32 bits in 32 bit systems and    -2147483648 to 2147483647 in 32 bit systems and
        64 bit in 64 bit systems.	     -9223372036854775808 to 9223372036854775807 in 64 bit systems

int8	8 bits/1 byte	                 -128 to 127
int16	16 bits/2 byte	                 -32768 to 32767
int32	32 bits/4 byte	                 -2147483648 to 2147483647
int64	64 bits/8 byte	                 -9223372036854775808 to 9223372036854775807

*/

                // Signed Integers
// Signed integers, declared with one of the int keywords, can store both positive and negative values:



/*
               Go has five keywords/types of unsigned integers:

Type	           Size	                               Range
uint	 Depends on platform:
         32 bits in 32 bit systems and           0 to 4294967295 in 32 bit systems and
         64 bit in 64 bit systems	             0 to 18446744073709551615 in 64 bit systems

uint8	 8 bits/1 byte	                         0 to 255
uint16	 16 bits/2 byte	                         0 to 65535
uint32   32 bits/4 byte	                         0 to 4294967295
uint64	 64 bits/8 byte	                         0 to 18446744073709551615 

The type of integer to choose, depends on the value the variable has to store.

Example:
This example will result in an error because 1000 is out of range for int8 (which is from -128 to 127):

package main
import ("fmt")

func main() {
  var x int8 = 1000
  fmt.Printf("Type: %T, value: %v", x, x)
}
Result: ./prog.go:5:7: constant 1000 overflows int8
                     
*/

//Unsigned Integers
// Unsigned integers, declared with one of the uint keywords, can only store non-negative values:



/*
                    Go Float Data Types
The float data types are used to store positive and negative numbers with a decimal point, like 35.3, -2.34, or 3597.34987.
The float data type has two keywords:

Type	   Size	          Range
float32	  32 bits    	-3.4e+38 to 3.4e+38.
float64	  64 bits	   -1.7e+308 to +1.7e+308.

Tip: The default type for float is float64. If you do not specify a type, the type will be float64. 


// Which Float Type to Use?
The type of float to choose, depends on the value the variable has to store.

Example
This example will result in an error because 3.4e+39 is out of range for float32:

package main
import ("fmt")

func main() {
  var x float32= 3.4e+39
  fmt.Println(x)
}
Result:

./prog.go:5:7: constant 3.4e+39 overflows float32
*/


/* 
                           String Data Type
The string data type is used to store a sequence of characters (text).
String values must be surrounded by double quotes:

Example
package main
import ("fmt")

func main() {
  var txt1 string = "Hello!"
  var txt2 string
  txt3 := "World 1"

  fmt.Printf("Type: %T, value: %\n", txt1, txt1)
  fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
  fmt.Printf("Type: %T, value: %v\n", txt3, txt3)
}
Result:

Type: string, value: Hello!
Type: string, value:
Type: string, value: World 1


*/



// Example
package main
import ("fmt")

func main() {
  var x int = 500
  var y int = -4500
  fmt.Printf("Type: %T, value: %v", x, x) // Type: int, value: 500
  fmt.Printf("Type: %T, value: %v", y, y) // Type: int, value: -4500

  
  var x1 uint = 500
  var y1 uint = 4500
  fmt.Printf("Type: %T, value: %v", x1, x1)// Type: uint, value: 500
  fmt.Printf("Type: %T, value: %v", y1, y1)// Type: uint, value: 4500


  var xt float32 = 123.78
  var yt float32 = 3.4e+38
  fmt.Printf("Type: %T, value: %v\n", xt, xt); // var x float32 = 123.78
  fmt.Printf("Type: %T, value: %v\n", yt, yt) ; // Type: float32, value: 3.4e+38
 
  var z float64 = 1.7e+308
  fmt.Printf("Type: %T, value: %v", z, z); // Type: float64, value: 1.7e+308

}

