package main

import "fmt"

func values() {
	fmt.Println("Go" + "lang") // Golang
	fmt.Println(10 + 21)       // 31
	fmt.Println(10 - 21)       // -11
	fmt.Println(10 * 21)       //210
	fmt.Println(42 / 21)       // 2
	fmt.Println(42 % 21)       // 0

	// ভগ্নাংশের (float) সাথে কাজ
	fmt.Println(15.5 + 4.5) // 20
	fmt.Println(9.0 / 2.0)  // 4.5

	fmt.Println(true && false) // flase
	fmt.Println(true || false) // true
	fmt.Println(!true)         // false
}

func comparision() {
	fmt.Println(10 > 5)   // true
	fmt.Println(10 < 5)   // false
	fmt.Println(10 == 10) // true
	fmt.Println(10 != 5)  // true
	fmt.Println(5 >= 5)   // true
	fmt.Println(5 <= 4)   // false
}

//(Type Casting):
func typeCasting() {
	var age int = 10
	var number float64 = 20.6
	var result float64 = float64(age) + number

	fmt.Println(result) // 30.6
}

/*
টাইপ মিসম্যাচ এড়াতে type casting ব্যবহার করুন।
float64(age) এর মাধ্যমে int কে float এ কনভার্ট করা হয়েছে।
*/
func main() {
	values()
	comparision()
	typeCasting()
	// Using var for variable declaration
	var name string = "Go Developer"
	var age int = 30
	var isActive bool = true

	// Using shorthand := for variable declaration
	country := "Bangladesh"
	lang := "Go"

	// ভুল: ভ্যারিয়েবল ইনিšিয়ালাইজ না করেই ব্যবহার করা
	var score int
	fmt.Println(score) // এটি 0 হবে কারণ Go তে ইনিশিয়ালাইজ না করলে ডিফল্ট 0 মান নেবে

	// Printing values
	fmt.Println("Name:", name)       // Go Developer
	fmt.Println("Age:", age)         // 30
	fmt.Println("Active:", isActive) // true
	fmt.Println("Country:", country) // Bangladesh
	fmt.Println("Language:", lang)   // Go
}

/*
ইন্টিজার এবং ফ্লোট ভিন্নভাবে ব্যবহার করতে হবে।
 int এবং float64 একসাথে যোগ করলে type mismatch হবে।

 স্ট্রিং কনক্যাটেনেশনের (concat) সময় "+" ব্যবহার করুন।
স্পষ্টতা বজায় রাখার জন্য স্ট্রিংগুলোর মাঝে স্পেস যোগ করুন।

তুলনামূলক অপারেটর ব্যবহার করলে Boolean (true/false) মান রিটার্ন করবে।
সবসময় == দিয়ে তুলনা করুন এবং != দিয়ে অসাম্য যাচাই করুন।

&& (AND), || (OR), ! (NOT) অপারেটর ব্যবহার করে শর্ত পরীক্ষা করুন।

Use meaningful variable names (name, age, isActive).
Prefer shorthand := for local variables if you don't need to declare the type explicitly.

2. Always Initialize Variables
সবসময় ভ্যারিয়েবল ইনিšিয়ালাইজ (অর্থাৎ, মান নির্ধারণ) করা উচিত।
Go তে যদি আপনি কোন ভ্যারিয়েবল ডিক্লেয়ার করেন এবং তাকে মান না দেন,
তবে সেটি ডিফল্ট ভ্যালু গ্রহণ করবে (যেমন, 0 ইনটিজারের জন্য, false বুলিয়ান জন্য, অথবা "" স্ট্রিংয়ের জন্য)।
তবে, এভাবে ভ্যারিয়েবল ব্যবহার করা সঠিক নয়,
কারণ এতে কোডে অস্পষ্টতা তৈরি হতে পারে।

1. Variable Names: Concise but Descriptive
ভ্যারিয়েবল নাম রাখতে হবে সংক্ষিপ্ত কিন্তু বর্ণনামূলক। এর মানে হলো, ভ্যারিয়েবল নাম এমনভাবে
 হওয়া উচিত যাতে তা ভ্যারিয়েবলের উদ্দেশ্য বা ব্যবহার সহজেই বুঝতে পারে। তবে,
 নামটি যেন খুব বড় বা জটিল না হয়।

*/
