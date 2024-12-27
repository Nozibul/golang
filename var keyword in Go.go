/*
Golang-এ var কিওয়ার্ডটি একটি নির্দিষ্ট টাইপের ভেরিয়েবল তৈরি করতে ব্যবহৃত হয়,
যার একটি সঠিক নাম এবং প্রাথমিক মান থাকে। ভেরিয়েবল ডিক্লেয়ার করার সময়
var কিওয়ার্ড ব্যবহার করে ইনিশিয়ালাইজেশন ঐচ্ছিক।

Go একটি স্ট্যাটিক্যালি type language, তবে এটি এখনও একটি সুবিধা প্রদান করে যা
ভেরিয়েবল ডিক্লেয়ার করার সময় ডেটা টাইপের ডিক্লেয়ারেশন বাদ দেয়, যেমন নিচের
সিনট্যাক্সে প্রদর্শিত হয়েছে। এটিকে সাধারণত টাইপ ইনফারেন্স বলা হয়।
exm: var identifier = initialValue

*/

/*
   // Multiple variable declarations using var Keyword

1. Declaring multiple variables using var keyword along with the type:
   exp:     var geek1, geek2, geek3, geek4 int

2. Declaring multiple variables using var keyword along with the type and initial values:
   exp:     var geek1, geek2, geek3, geek4 int = 10, 20, 30, 40


আপনি উপরের মতো টাইপ ইনফারেন্সও ব্যবহার করতে পারেন, যা কম্পাইলারকে টাইপ
সম্পর্কে জানাতে সাহায্য করবে, অর্থাৎ একাধিক ভেরিয়েবল ডিক্লেয়ার করার সময়
টাইপটি সরিয়ে দেওয়ার একটি অপশন রয়েছে।
   exp:     var geek1, geek2, geek3, geek4 = 10, 20, 30.30, true



আপনি var কিওয়ার্ড ব্যবহার করে বিভিন্ন টাইপের মান ডিক্লেয়ার এবং ইনিশিয়ালাইজ
করার জন্য একাধিক লাইনও ব্যবহার করতে পারেন,
যেমন নিচে দেখানো হয়েছে:
		var(
			geek1 = 100
			geek2 = 200.57
			geek3 bool
			geek4 string = "GeeksforGeeks"
		)


ডিক্লেয়ারেশনের সময় টাইপ ব্যবহার করলে আপনি শুধুমাত্র একই টাইপের একাধিক
ভেরিয়েবল ডিক্লেয়ার করতে পারবেন। কিন্তু ডিক্লেয়ারেশনের সময় টাইপ সরিয়ে দিলে,
আপনি বিভিন্ন টাইপের একাধিক ভেরিয়েবল ডিক্লেয়ার করতে পারবেন।

package main

import "fmt"

func main() {

 // Multiple variables of the same type
    // are declared and initialized
    // in the single line along with type
    var geek1, geek2, geek3 int = 232, 784, 854

    // Multiple variables of different type
    // are declared and initialized
    // in the single line without specifying
    // any type
    var geek4, geek5, geek6 = 100, "GFG", 7896.46


   // Display the values of the variables
   fmt.Printf("The value of geek1 is : %d\n", geek1)

   fmt.Printf("The value of geek2 is : %d\n", geek2)

   fmt.Printf("The value of geek3 is : %d\n", geek3)

   fmt.Printf("The value of geek4 is : %d\n", geek4)

   fmt.Printf("The value of geek5 is : %s\n", geek5)

   fmt.Printf("The value of geek6 is : %f", geek6)

}
Output:

The value of geek1 is : 232
The value of geek2 is : 784
The value of geek3 is : 854
The value of geek4 is : 100
The value of geek5 is : GFG
The value of geek6 is : 7896.460000

*/

/*
var কিওয়ার্ড সম্পর্কে গুরুত্বপূর্ণ পয়েন্টগুলি:

যখন আপনি var কিওয়ার্ড ব্যবহার করে ভেরিয়েবল ডিক্লেয়ার করেন, তখন আপনি টাইপ
বা = এক্সপ্রেশন যেকোনো একটি সরাতে পারবেন, কিন্তু উভয়ই একসাথে সরানো যাবে না।
যদি আপনি উভয়ই সরান, তাহলে কম্পাইলার একটি এরর দিবে।

যদি আপনি এক্সপ্রেশন সরিয়ে দেন, তাহলে ভেরিয়েবলটি ডিফল্টভাবে সংখ্যা জন্য শূন্য
মান, বুলিয়ান জন্য false, স্ট্রিং জন্য "" এবং ইন্টারফেস বা রেফারেন্স টাইপের জন্য
nil ধারণ করবে। তাই Go ভাষায় কোনো অইনিশিয়ালাইজড ভেরিয়েবলের ধারণা নেই।


import "fmt"

func main() {

    // Variable declared but
    // no initialization
    var geek1 int
    var geek2 string
    var geek3 float64
    var geek4 bool

    // Display the zero-value of the variables
    fmt.Printf("The value of geek1 is : %d\n", geek1)

    fmt.Printf("The value of geek2 is : %s\n", geek2)

    fmt.Printf("The value of geek3 is : %f\n", geek3)

    fmt.Printf("The value of geek4 is : %t", geek4)

}
Output:

The value of geek1 is : 0
The value of geek2 is :
The value of geek3 is : 0.000000
The value of geek4 is : false

*/