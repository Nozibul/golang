package main

import "fmt"

func main() {
	var i, j string = "Hello", "World" // দুটি স্ট্রিং ভেরিয়েবল ডিক্লেয়ার করা হয়েছে

	// যদি আর্গুমেন্টগুলিকে নতুন লাইনে প্রিন্ট করতে চাই, তাহলে \n ব্যবহার করতে হবে
	fmt.Print(i, "\n")    // এটি "Hello" প্রিন্ট করবে (নতুন লাইনে)
	fmt.Print(i, "\n", j) // এটি "Hello" এবং "World" প্রিন্ট করবে (নতুন লাইনে)

	// যদি স্ট্রিং আর্গুমেন্টগুলির মধ্যে একটি স্পেস যোগ করতে চাই, তাহলে " " ব্যবহার করতে হবে
	fmt.Println(i, " ", j) // এটি "Hello World" প্রিন্ট করবে (স্পেস সহ)

	// অন্যান্য ভেরিয়েবলের উদাহরণ:
	var n = "Go"
	var o = 34
	var k string = "Hello"
	var m int = 15

	// Printf() ফাংশন: এখানে %v ব্যবহার করা হয়েছে মান এবং %T টাইপ প্রিন্ট করার জন্য
	fmt.Printf("%v has value: %v and type: %T\n", n, o, n) // "Go has value: 34 and type: string"
	fmt.Printf("k has value: %v and type: %T\n", k, k)     // "Hello has value: Hello and type: string"
	fmt.Printf("m has value: %v and type: %T", m, m)       // "15 has value: 15 and type: int"
}

/*
                              fmt.Print():
fmt.Print() ফাংশন আর্গুমেন্টগুলিকে তাদের ডিফল্ট ফরম্যাটে প্রিন্ট করে।
যদি আমরা আর্গুমেন্টগুলিকে নতুন লাইনে প্রিন্ট করতে চাই, তাহলে \n ব্যবহার করতে হয়।
উদাহরণ: fmt.Print(i, "\n") এটি "Hello" প্রিন্ট করবে এবং নতুন লাইনে যাবে।

                              fmt.Println():
fmt.Println() ফাংশন আর্গুমেন্টগুলিকে প্রিন্ট করার পরে একটি নতুন লাইন যোগ করে দেয়।
উদাহরণ: fmt.Println(i, " ", j) এটি "Hello World" প্রিন্ট করবে, যেখানে দুটি আর্গুমেন্টের মধ্যে একটি স্পেস থাকবে।

                              fmt.Printf():
fmt.Printf() ফাংশন ফরম্যাটিংয়ের জন্য ব্যবহার হয়। এখানে দুটি ফরম্যাটিং verb ব্যবহার করা হয়েছে:
%v: এটি আর্গুমেন্টের মান (value) প্রিন্ট করবে।
%T: এটি আর্গুমেন্টের টাইপ (type) প্রিন্ট করবে।
উদাহরণ: fmt.Printf("%v has value: %v and type: %T\n", n, o, n) এটি "Go has value: 34 and type: string" প্রিন্ট করবে।

                                সারাংশ:
fmt.Print(): ডিফল্ট ফরম্যাটে প্রিন্ট করে এবং \n দিয়ে নতুন লাইনে যায়।
fmt.Println(): প্রিন্ট করার পরে একটি নতুন লাইন যোগ করে।
fmt.Printf(): আর্গুমেন্টের মান এবং টাইপ প্রিন্ট করার জন্য ফরম্যাটিং verb ব্যবহার করা হয়।

*/
