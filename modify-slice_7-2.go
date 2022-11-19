/* 
//                   Go Access, Change, Append and Copy Slices

                    // Access Elements of a Slice
You can access a specific slice element by referring to the index number.

In Go, indexes start at 0. That means that [0] is the first
element, [1] is the second element, etc.


Example:
This example shows how to access the first and third elements in the prices slice:

package main
import ("fmt")

func main() {
  prices := []int{10,20,30}

  fmt.Println(prices[0])
  fmt.Println(prices[2])
}
Result:

10
30


                // Append Elements To a Slice
You can append elements to the end of a slice using the append()function:

Syntax
slice_name = append(slice_name, element1, element2, ...)
Example
This example shows how to append elements to the end of a slice:

package main
import ("fmt")

func main() {
  myslice1 := []int{1, 2, 3, 4, 5, 6}
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))

  myslice1 = append(myslice1, 20, 21)
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))
}
Result:

myslice1 = [1 2 3 4 5 6]
length = 6
capacity = 6
myslice1 = [1 2 3 4 5 6 20 21]
length = 8
capacity = 12




            // Append One Slice To Another Slice
To append all the elements of one slice to another slice,
use the append()function:

Syntax
slice3 = append(slice1, slice2...)
Note: The '...' after slice2 is necessary when appending the
elements of one slice to another.


            Example :
This example shows how to append one slice to another slice:

package main
import ("fmt")

func main() {
  myslice1 := []int{1,2,3}
  myslice2 := []int{4,5,6}
  myslice3 := append(myslice1, myslice2...)

  fmt.Printf("myslice3=%v\n", myslice3)
  fmt.Printf("length=%d\n", len(myslice3))
  fmt.Printf("capacity=%d\n", cap(myslice3))
}
Result:
myslice3=[1 2 3 4 5 6]
length=6
capacity=6



                 // Change The Length of a Slice
Unlike arrays, it is possible to change the length of a slice.

Example:
This example shows how to change the length of a slice:

package main
import ("fmt")

func main() {
  arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array  
  myslice1 := arr1[1:5] // Slice array
  fmt.Printf("myslice1 = %v\n", myslice1)   // myslice1 = [10 11 12 13]
  fmt.Printf("length = %d\n", len(myslice1)) // length = 4
  fmt.Printf("capacity = %d\n", cap(myslice1)) // capacity = 5

  myslice1 = arr1[1:3] // Change length by re-slicing the array
  fmt.Printf("myslice1 = %v\n", myslice1) // myslice1 = [10 11]
  fmt.Printf("length = %d\n", len(myslice1)) // length = 2
  fmt.Printf("capacity = %d\n", cap(myslice1)) // capacity = 5

  myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items
  fmt.Printf("myslice1 = %v\n", myslice1) // myslice1 = [10 11 20 21 22 23]
  fmt.Printf("length = %d\n", len(myslice1)) // length = 6
  fmt.Printf("capacity = %d\n", cap(myslice1)) // capacity = 10
}




                // Memory Efficiency
When using slices, Go loads all the underlying elements into the memory.
স্লাইস ব্যবহার করার সময়, গো মেমরিতে অন্তর্নিহিত উপাদানগুলি লোড করে।

If the array is large and you need only a few elements, 
It is better to copy those elements using the copy() function.
যদি অ্যারেটি বড় হয় এবং আপনার শুধুমাত্র কয়েকটি উপাদানের প্রয়োজন হয়, 
তাহলে কপি() ফাংশন ব্যবহার করে সেই উপাদানগুলি অনুলিপি করা ভাল।

The copy() function creates a new underlying array with only the required elements for the slice.
This will reduce the memory used for the program.

কপি() ফাংশন স্লাইসের জন্য শুধুমাত্র প্রয়োজনীয় উপাদান সহ একটি নতুন অন্তর্নিহিত অ্যারে তৈরি করে। 
এটি প্রোগ্রামের জন্য ব্যবহৃত মেমরি হ্রাস করবে।


Syntax:
copy(dest, src)
The copy() function takes in two slices dest and src, 
and copies data from src to dest. It returns the number of elements copied. 
কপি() ফাংশনটি দুটি স্লাইস dest এবং src নেয়,
এবং src থেকে গন্তব্যে ডেটা কপি করে। এটি অনুলিপি করা উপাদানের সংখ্যা প্রদান করে।

Example:
This example shows how to use the copy() function:

package main
import ("fmt")

func main() {
  numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
  // Original slice
  fmt.Printf("numbers = %v\n", numbers) // numbers = [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]
  fmt.Printf("length = %d\n", len(numbers)) // length = 15
  fmt.Printf("capacity = %d\n", cap(numbers)) // capacity = 15

  // Create copy with only needed numbers
  neededNumbers := numbers[:len(numbers)-10] // len(15-10 = 5) 
  numbersCopy := make([]int, len(neededNumbers)) 
  copy(numbersCopy, neededNumbers) // neededNumbers theke element nia numbersCopy te rekhe dibo

  fmt.Printf("numbersCopy = %v\n", numbersCopy) // numbersCopy = [1 2 3 4 5]
  fmt.Printf("length = %d\n", len(numbersCopy)) // length = 5
  fmt.Printf("capacity = %d\n", cap(numbersCopy)) // capacity = 5
}



*/