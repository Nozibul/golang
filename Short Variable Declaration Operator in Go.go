/*
গোল্যাং-এ শর্ট ভেরিয়েবল ডিক্লেয়ারেশন অপারেটর (:=)

গোল্যাং-এ := শর্ট ভেরিয়েবল ডিক্লেয়ারেশন অপারেটরটি ব্যবহৃত হয় নতুন ভেরিয়েবল
ডিক্লেয়ার এবং ইনিশিয়ালাইজ করার জন্য। এর মাধ্যমে টাইপ উল্লেখ করার প্রয়োজন নেই,
কারণ কম্পাইলার এক্সপ্রেশন থেকে টাইপটি নির্ধারণ করে। সাধারণত, এটি লোকাল
(local) ভেরিয়েবল ডিক্লেয়ার করতে ব্যবহৃত হয়, যেগুলি শুধুমাত্র ফাংশনের ভিতরে
কাজ করে এবং ফাংশন ব্লকের স্কোপে থাকে।

variable_name := expression

এখানে, ডিক্লেয়ারেশনের সময় টাইপ উল্লেখ করার প্রয়োজন নেই, এবং একই সময়ে
ভেরিয়েবলটি ইনিশিয়ালাইজও করতে হয়।
*/

package main

import "fmt"

func main() {
	// শর্ট ডিক্লেয়ারেশন অপারেটর দিয়ে ভেরিয়েবল ডিক্লেয়ার এবং ইনিশিয়ালাইজ করা
	a := 30
	fmt.Println("The Value of a is: ", a) // আউটপুট: 30

	// টাইপ নির্ধারণ করা ছাড়া স্ট্রিং ভেরিয়েবল ডিক্লেয়ার করা
	Language := "Go Programming"
	fmt.Println("The Value of Language is: ", Language) // আউটপুট: Go Programming
}

/*
                একাধিক ভেরিয়েবল ডিক্লেয়ারেশন:

শর্ট ডিক্লেয়ারেশন অপারেটর দিয়ে একসাথে একাধিক ভেরিয়েবলও ডিক্লেয়ার করা যায়,
এবং এগুলির টাইপটি এক্সপ্রেশনের মাধ্যমে নির্ধারিত হয়। একে একসাথে একই টাইপের
বা ভিন্ন টাইপের ভেরিয়েবল ডিক্লেয়ার করা যায়।
*/

func main() {
	// একাধিক ভেরিয়েবল ডিক্লেয়ার এবং ইনিশিয়ালাইজ করা
	geek1, geek2, geek3 := 117, 7834, 5685

	// একাধিক ভেরিয়েবল ডিক্লেয়ার এবং ইনিশিয়ালাইজ করা
	geek4, geek5, geek6 := "GFG", 859.24, 1234

	fmt.Printf("The value of geek1 is : %d\n", geek1)
	fmt.Printf("The type of geek1 is : %T\n", geek1)

	fmt.Printf("\nThe value of geek2 is : %d\n", geek2)
	fmt.Printf("The type of geek2 is : %T\n", geek2)

	fmt.Printf("\nThe value of geek3 is : %d\n", geek3)
	fmt.Printf("The type of geek3 is : %T\n", geek3)

	fmt.Printf("\nThe value of geek4 is : %s\n", geek4)
	fmt.Printf("The type of geek4 is : %T\n", geek4)

	fmt.Printf("\nThe value of geek5 is : %f\n", geek5)
	fmt.Printf("The type of geek5 is : %T\n", geek5)

	fmt.Printf("\nThe value of geek6 is : %d\n", geek6)
	fmt.Printf("The type of geek6 is : %T\n", geek6)
}

/*
The value of geek1 is : 117
The type of geek1 is : int

The value of geek2 is : 7834
The type of geek2 is : int

The value of geek3 is : 5685
The type of geek3 is : int

The value of geek4 is : GFG
The type of geek4 is : string

The value of geek5 is : 859.240000
The type of geek5 is : float64

The value of geek6 is : 1234
The type of geek6 is : int

*/

/*
       গুরুত্বপূর্ণ পয়েন্টসমূহ:
1. শর্ট ডিক্লেয়ারেশন অপারেটর শুধুমাত্র তখন ব্যবহার করা যেতে পারে যখন অন্তত
একটি নতুন ভেরিয়েবল ডিক্লেয়ার করা হয়। যদি সমস্ত ভেরিয়েবল পূর্বে ডিক্লেয়ার করা
থাকে, তবে এটি এ্যাসাইনমেন্ট এর মতো আচরণ করবে।

2. যদি একই নামের ভেরিয়েবল পুনরায় ডিক্লেয়ার করা হয়, তাহলে ত্রুটি হবে।
উদাহরণস্বরূপ:

package main

import "fmt"

func main() {
    p, q := 100, 200
    fmt.Println("Value of p:", p, "Value of q:", q)

    // এখানে ত্রুটি হবে, কারণ p ও q আগেই ডিক্লেয়ার করা হয়েছে
    p, q := 500, 600
    fmt.Println("New value of p:", p, "New value of q:", q)
}

ত্রুটি:
./prog.go:17:10: no new variables on left side of :=

*/

/*
3. অ্যাসাইনমেন্ট এর মতো আচরণ করার সময়, যদি ভেরিয়েবল নতুন না হয়, তবে
এক্সপ্রেশনগুলির মান ভেরিয়েবলে সরাসরি অ্যাসাইন করা হবে।

4. গোল্যাং একটি শক্তিশালী টাইপ সিস্টেম। এক ভেরিয়েবলে অন্য টাইপের মান অ্যাসাইন
করা যতটুকু সম্ভব, ততটুকু অনুমোদিত নয়। যেমন:

package main

import "fmt"

func main() {
    z := 50
    fmt.Printf("Value of z is %d", z)

    // এই লাইনে ত্রুটি হবে, কারণ আপনি একটি ইনটিজার টাইপের ভেরিয়েবলে স্ট্রিং অ্যাসাইন করছেন
    z := "Golang"
}
ত্রুটি:

./prog.go:16:4: no new variables on left side of :=
./prog.go:16:7: cannot use “Golang” (type string) as type int in assignment
*/
