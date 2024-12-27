/*
Go-তে ভেরিয়েবলের স্কোপ মানে হল প্রোগ্রামের কোন অংশে সেই ভেরিয়েবলকে ব্যবহার
করা যাবে।

এটা সহজভাবে এমন – ভেরিয়েবল হলো একটি নাম বা ট্যাগ যা কোন মান (value)
ধরে রাখে। আর স্কোপ হলো সেই নাম বা ট্যাগটি প্রোগ্রামের কোন অংশে দেখা বা
ব্যবহার করা যাবে তা নির্ধারণ করে।
*/

package main

import "fmt"

// Global variable declaration
var myVariable int = 100

func main() {
	// Local variable inside the main function
	var localVar int = 200
	fmt.Printf("Inside main - Global variable: %d\n", myVariable)
	fmt.Printf("Inside main - Local variable: %d\n", localVar)
	display()
}

func display() {
	fmt.Printf("Inside display - Global variable: %d\n", myVariable)
}

// Local Variables
/*
লোকাল ভেরিয়েবল হচ্ছে এমন ভেরিয়েবল, যা একটি ফাংশন বা ব্লকের ভেতরে ডিক্লেয়ার
(ঘোষণা) করা হয় এবং সেটি সেই ফাংশন বা ব্লকের বাইরের অংশে দেখা যায় না।

কীভাবে কাজ করে:
যদি তুমি কোন ফাংশন, লুপ (যেমন for), বা শর্ত (যেমন if) এর ভেতরে ভেরিয়েবল
তৈরি করো, সেটি শুধুমাত্র সেই ব্লকের মধ্যেই সীমাবদ্ধ থাকবে।
ফাংশন বা ব্লকের বাইরে গেলে সেই ভেরিয়েবল আর কাজ করবে না।
*/

func main() {
	var localVar int = 200       // Local variable
	fmt.Printf("%d\n", localVar) // Accessible here
}

//Global Variables
/*

গ্লোবাল ভেরিয়েবল হলো এমন ভেরিয়েবল যা ফাংশন বা ব্লকের বাইরে ডিক্লেয়ার করা হয়।
এ কারণে, এই ভেরিয়েবলটি পুরো প্রোগ্রামে যে কোনো জায়গা থেকে ব্যবহার করা যায়।

        গ্লোবাল ভেরিয়েবল কিভাবে কাজ করে:
এটি ফাংশনের বাইরে ডিফাইন করা হয়।
প্রোগ্রামের যেকোনো ফাংশন থেকে গ্লোবাল ভেরিয়েবলে অ্যাক্সেস করা যায় বা তার মান
পরিবর্তন করা যায়।
*/

// Global variable declaration
var myVariables int = 100 // Global variable

func main() {
	fmt.Printf("%d\n", myVariables) // Accessible here
}

//Local Variable Preference
/*
যদি একটি লোকাল ভেরিয়েবলের নাম এবং একটি গ্লোবাল ভেরিয়েবলের নাম একই হয়,
তাহলে লোকাল ভেরিয়েবলটি তার স্কোপের (scope) মধ্যে প্রাধান্য পায়। এটি গ্লোবাল
ভেরিয়েবলকে "ছাপিয়ে" যায়, অর্থাৎ, ফাংশন বা ব্লকের ভিতরে গ্লোবাল ভেরিয়েবলটি
ব্যবহার করা সম্ভব হয় না যতক্ষণ না লোকাল ভেরিয়েবলের স্কোপের বাইরে যাওয়া হয়।
*/
// Global variable declaration
var myVariabless int = 100 // Global variable

func main() {
	var myVariabless int = 200                                        // Local variable
	fmt.Printf("Local variable takes precedence: %d\n", myVariabless) // Accesses local variable
}

// output: 200
