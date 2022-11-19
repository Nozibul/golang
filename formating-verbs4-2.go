/*
Formatting Verbs for Printf()
Go offers several formatting verbs that can be used with the Printf() fu 

General Formatting Verbs
The following verbs can be used with all data types:
Verb	Description
 %v 	Prints the value in the default format
 %#v	Prints the value in Go-syntax format
 %T		Prints the type of the value
 %%		Prints the % sign


               Integer Formatting Verbs
The following verbs can be used with the integer data type:
Verb	Description
 %b	      Base 2
 %d	      Base 10
 %+d      Base 10 and always show sign
 %o	      Base 8
 %O	      Base 8, with leading 0o
 %x	      Base 16, lowercase
 %X	      Base 16, uppercase
 %#x      Base 16, with leading 0x
 %4d      Pad with spaces (width 4, right justified)
 %-4      d	Pad with spaces (width 4, left justified)
 %04      d	Pad with zeroes (width 4


     String Formatting Verbs
The following verbs can be used with the string data type:

Verb	Description
 %s	     Prints the value as plain string
 %q	     Prints the value as a double-quoted string
 %8s	 Prints the value as plain string (width 8, right justified)
 %-8s	 Prints the value as plain string (width 8, left justified)
 %x   	 Prints the value as hex dump of byte values
 % x	 Prints the value as hex dump with spaces


                // Boolean Formatting Verbs
The following verb can be used with the boolean data type:

Verb	Description
 %t	    Value of the boolean operator in true or false format (same as using %v)



               // Float Formatting Verbs
The following verbs can be used with the float data type:

Verb	Description
 %e	     Scientific notation with 'e' as exponent
 %f	     Decimal point, no exponent
 %.2f	 Default width, precision 2
 %6.2f	 Width 6, precision 2
 %g	     Exponent as needed, only necessary digits

*/

package main
import ("fmt")

func main() {
  var i = 15.5
  var txts = "Hello World!"

  fmt.Printf("%v\n", i) // 15.5   
  fmt.Printf("%#v\n", i) // 15.5  
  fmt.Printf("%v%%\n", i) // 15.5%
  fmt.Printf("%T\n", i) // float64

  fmt.Printf("%v\n", txts) //Hello World!
  fmt.Printf("%#v\n", txts) //"Hello World!"
  fmt.Printf("%T\n", txts)  // string


  var r = 15
 
  fmt.Printf("%b\n", r) // 1111 -->  %b  Base 2
  fmt.Printf("%d\n", r) // 15 -->   %d  Base 10
  fmt.Printf("%+d\n", r)// +15 -->   %+d  Base 10 and always show sign
  fmt.Printf("%o\n", r) // 17 -->   %o  Base 8
  fmt.Printf("%O\n", r) // 0o17 -->   %O  Base 8, with leading 0o
  fmt.Printf("%x\n", r) // f -->  %x  Base 16, lowercase
  fmt.Printf("%X\n", r) // F -->  %X  Base 16, uppercase
  fmt.Printf("%#x\n", r)// 0xf -->  %#x  Base 16, with leading 0x
  fmt.Printf("%4d\n", r)// 15 -->  %4d   Pad with spaces (width 4, right justified)
  fmt.Printf("%-4d\n", r)// 15 -->  %-4  d Pad with spaces (width 4, left justified)
  fmt.Printf("%04d\n", r) // 0015 -->   %04   d	Pad with zeroes (width 4)




  var txt = "Hello"
 
  fmt.Printf("%s\n", txt) // Hello -->  %s  Prints the value as plain string
  fmt.Printf("%q\n", txt) // "Hello" -->  %q  Prints the value as a double-quoted string
  fmt.Printf("%8s\n", txt)//    Hello -->  %8s	Prints the value as plain string (width 8, right justified)
  fmt.Printf("%-8s\n", txt)// Hello  -->   %-8s	 Prints the value as plain string (width 8, left justified)
  fmt.Printf("%x\n", txt) // 48656c6c6f --> %x  Prints the value as hex dump of byte values
  fmt.Printf("% x\n", txt)// 48 65 6c 6c 6f -->  %x Prints the value as hex dump with spaces


  var h = false
  var s = true

  // %t	  Value of the boolean operator in true or false format (same as using %v)
  fmt.Printf("%t\n", h)// false
  fmt.Printf("%t\n", s) // true




  var v = 3.141

  fmt.Printf("%e\n", v) //3.141000e+00 --> %e  Scientific notation with 'e' as exponent
  fmt.Printf("%f\n", v) //3.141000 --> %f  Decimal point, no exponent
  fmt.Printf("%.2f\n", v) //3.14 --> %.2f	Default width, precision 2
  fmt.Printf("%6.2f\n", v)//  3.14 --> %6.2f	 Width 6, precision 2
  fmt.Printf("%g\n", v) //3.141 --> %g  Exponent as needed, only necessary digits


}