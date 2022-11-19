/* 
      Go Maps

Maps are used to store data values in key:value pairs.

Each element in a map is a key:value pair.

A map is an unordered and changeable collection that does not allow duplicates.

The length of a map is the number of its elements. You can find it using the len() function.

The default value of a map is nil.

Maps hold references to an underlying hash table.

Go has multiple ways for creating maps.

Creating Maps Using var and :=
Syntax:
var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
b := map[KeyType]ValueType{key1:value1, key2:value2,...}



                 Example :
This example shows how to create maps in Go. Notice the order in the code and in the output

package main
import ("fmt")

func main() {
  var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
  b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

  fmt.Printf("a\t%v\n", a)
  fmt.Printf("b\t%v\n", b)
}
Result:
a   map[brand:Ford model:Mustang year:1964]
b   map[Bergen:2 Oslo:1 Stavanger:4 Trondheim:3]

Note: The order of the map elements defined in the code is different from
the way that they are stored. The data are stored in a way to have efficient
 data retrieval from the map.
দ্রষ্টব্য: কোডে সংজ্ঞায়িত মানচিত্রের উপাদানগুলির ক্রম সেগুলি যেভাবে সংরক্ষণ করা হয় তার থেকে আলাদা। মানচিত্র থেকে দক্ষ ডেটা 
পুনরুদ্ধার করার জন্য ডেটা এমনভাবে সংরক্ষণ করা হয়।






           // Creating Maps Using Using make()Function:
Syntax
var a = make(map[KeyType]ValueType)
b := make(map[KeyType]ValueType)

This example shows how to create maps in Go using the make()function.

package main
import ("fmt")

func main() {
  var a = make(map[string]string) // The map is empty now
  a["brand"] = "Ford"
  a["model"] = "Mustang"
  a["year"] = "1964"
                                 // a is no longer empty
  b := make(map[string]int)
  b["Oslo"] = 1
  b["Bergen"] = 2
  b["Trondheim"] = 3
  b["Stavanger"] = 4

  fmt.Printf("a\t%v\n", a)
  fmt.Printf("b\t%v\n", b)
}
Result:

a   map[brand:Ford model:Mustang year:1964]
b   map[Bergen:2 Oslo:1 Stavanger:4 Trondheim:3]



                        // Creating an Empty Map
There are two ways to create an empty map. One is by using the make()function and the other is by using the following syntax.

Syntax:
var a map[KeyType]ValueType
Note: The make()function is the right way to create an empty map. If you make an empty map in a different way and write to it, it will causes a runtime panic.

Example :
This example shows the difference between declaring an empty map using with the make()function and without it.

package main
import ("fmt")

func main() {
  var a = make(map[string]string)
  var b map[string]string

  fmt.Println(a == nil)
  fmt.Println(b == nil)
}
Result:
false
true




                    // Allowed Key Types
The map key can be of any data type for which the equality operator 
(==) is defined. These include:

Booleans
Numbers
Strings
Arrays
Pointers
Structs
Interfaces (as long as the dynamic type supports equality)

Invalid key types are:
Slices
Maps
Functions
These types are invalid because the equality operator (==)
is not defined for them.


               // Allowed Value Types
The map values can be any type.



               // Accessing Map Elements
You can access map elements by:

Syntax
value = map_name[key]
Example
package main
import ("fmt")

func main() {
  var a = make(map[string]string)
  a["brand"] = "Ford"
  a["model"] = "Mustang"
  a["year"] = "1964"

  fmt.Printf(a["brand"])
}
Result:  Ford




                 // Updating and Adding Map Elements
Updating or adding an elements are done by:

Syntax
map_name[key] = value
Example
This example shows how to update and add elements to a map.

package main
import ("fmt")

func main() {
  var a = make(map[string]string)
  a["brand"] = "Ford"
  a["model"] = "Mustang"
  a["year"] = "1964"

  fmt.Println(a)

  a["year"] = "1970" // Updating an element
  a["color"] = "red" // Adding an element

  fmt.Println(a)
}
Result:
map[brand:Ford model:Mustang year:1964]
map[brand:Ford color:red model:Mustang year:1970]



                      // Remove Element from Map
Removing elements is done using the delete() function.

Syntax :
delete(map_name, key)
Example
package main
import ("fmt")

func main() {
  var a = make(map[string]string)
  a["brand"] = "Ford"
  a["model"] = "Mustang"
  a["year"] = "1964"

  fmt.Println(a)

  delete(a,"year")

  fmt.Println(a)
}
Result:
map[brand:Ford model:Mustang year:1964]
map[brand:Ford model:Mustang]





                 // Check For Specific Elements in a Map
You can check if a certain key exists in a map using:

Syntax:
val, ok :=map_name[key]
If you only want to check the existence of a certain key, you can use the blank identifier (_) in place of val.

Example
package main
import ("fmt")

func main() {
  var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day":""}

  val1, ok1 := a["brand"] // Checking for existing key and its value
  val2, ok2 := a["color"] // Checking for non-existing key and its value
  val3, ok3 := a["day"]   // Checking for existing key and its value
  _, ok4 := a["model"]    // Only checking for existing key and not its value

  fmt.Println(val1, ok1)
  fmt.Println(val2, ok2)
  fmt.Println(val3, ok3)
  fmt.Println(ok4)
}
Result:

Ford true
 false
 true
true




                     // Maps Are References
Maps are references to hash tables.

If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.

Example
package main
import ("fmt")

func main() {
  var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
  b := a

  fmt.Println(a)
  fmt.Println(b)

  b["year"] = "1970"
  fmt.Println("After change to b:")

  fmt.Println(a)
  fmt.Println(b)
}
Result:

map[brand:Ford model:Mustang year:1964]
map[brand:Ford model:Mustang year:1964]
After change to b:
map[brand:Ford model:Mustang year:1970]
map[brand:Ford model:Mustang year:1970]




                     // Iterating Over Maps
You can use range to iterate over maps.

Example
This example shows how to iterate over the elements in a map. Note the order of the elements in the output.

package main
import ("fmt")

func main() {
  a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

  for k, v := range a {
    fmt.Printf("%v : %v, ", k, v)
  }
}
Result:

two : 2, three : 3, four : 4, one : 1,




package main
import ("fmt")

func main() {
  a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

  var b []string                // defining the order
  b = append(b, "one", "two", "three", "four")

  for _, element := range b {  // loop with the defined order
    fmt.Printf("%v : %v, ", element, a[element])
  }
}
result: one : 1, two : 2, three : 3, four : 4,
*/