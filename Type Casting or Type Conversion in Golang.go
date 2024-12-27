/*
Golang-এ Type Conversion (টাইপ কনভার্সন) কী এবং কেন দরকার?
টাইপ কনভার্সন হলো যখন আমরা একটি ডেটা টাইপের মানকে অন্য ডেটা টাইপে অ্যাসাইন (assign) করি।
অনেক স্ট্যাটিক্যালি টাইপড (Statically Typed) ভাষা যেমন C/C++, Java স্বয়ংক্রিয়ভাবে Implicit Type Conversion (ইমপ্লিসিট টাইপ কনভার্সন) সাপোর্ট করে।
তবে, Golang এ ব্যাপারটা আলাদা। Go স্বয়ংক্রিয়ভাবে ডেটা টাইপ পরিবর্তন করে না।

Golang-এ Implicit Type Conversion নেই, কেন?
Golang-এ Implicit Type Conversion (স্বয়ংক্রিয় টাইপ কনভার্সন) নেই কারণ:
👉 Golang একটি Strongly Typed (স্ট্রং টাইপড) ভাষা।
👉 এটি অটোমেটিক টাইপ কনভার্সন করতে দেয় না, এমনকি যদি ডেটা টাইপ কম্প্যাটিবল (compatible) হয় তবুও।

Type Conversion করতে হলে কী করতে হবে?
Golang-এ যদি এক ডেটা টাইপ থেকে অন্য ডেটা টাইপে পরিবর্তন করতে হয়, তবে Explicit Conversion (এক্সপ্লিসিট কনভার্সন) করতে হবে।
অর্থাৎ, আপনাকে স্পষ্টভাবে কনভার্সন করতে হবে।

Golang-এ Typecasting নেই, শুধু Type Conversion আছে!
👉 Golang স্পেসিফিকেশনে "Type Casting" বলে কিছু নেই।
👉 যদি আপনি Golang ডকুমেন্টেশনে "Type Casting" খুঁজেন, কিছুই পাবেন না।
👉 Golang-এ শুধুমাত্র Type Conversion (টাইপ কনভার্সন) শব্দটি ব্যবহৃত হয়।
👉 অন্যান্য ভাষায়, Type Casting এবং Type Conversion একই জিনিস বলে মনে করা হয়।

Type Conversion-এর দরকার কেন?
কখনো কখনো ডেটা টাইপের নির্দিষ্ট বৈশিষ্ট্যগুলোর সুবিধা নিতে হলে, এক ডেটা টাইপ থেকে অন্য টাইপে কনভার্ট করতে হয়।
উদাহরণস্বরূপ:
ইন্টিজার (int) থেকে ফ্লোট (float)
স্ট্রিং (string) থেকে বাইট (byte)

Type Casting vs Type Conversion (Bangla Explanation)
অনেক সময় প্রোগ্রামিং ভাষায় একটি ডেটা টাইপের মানকে অন্য ডেটা টাইপে পরিবর্তন করতে হয়। এই পরিবর্তনের জন্য সাধারণত দুটি টার্ম ব্যবহার হয়:

1. Type Casting (টাইপ ক্যাস্টিং)
2. Type Conversion (টাইপ কনভার্সন)

	Type Casting (টাইপ ক্যাস্টিং):

Type Casting হলো এমন একটি প্রক্রিয়া যেখানে একটি ডেটা টাইপ জোরপূর্বক (forcibly) অন্য ডেটা টাইপে পরিবর্তন করা হয়, এমনকি যদি সেই টাইপ একে অপরের সাথে পুরোপুরি সামঞ্জস্যপূর্ণ না হয়।

🔹 কীভাবে কাজ করে:

কম্পাইলার বা ইন্টারপ্রেটার মনে করে আপনি জানেন আপনি কী করছেন।
এটি Unsafe (অনিরাপদ) হতে পারে কারণ ভুল টাইপ ক্যাস্টিং প্রোগ্রামে বাগ বা ক্র্যাশ তৈরি করতে পারে।
🔹 কোথায় ব্যবহৃত হয়:

C/C++, Java এবং অন্যান্য ভাষায় টাইপ ক্যাস্টিং প্রচলিত।

🔹 উদাহরণ (C/C++):

int x = 10;
float y = (float)x;  // টাইপ ক্যাস্টিং - int থেকে float-এ জোরপূর্বক পরিবর্তন

	Type Conversion (টাইপ কনভার্সন):

Type Conversion হলো এমন একটি প্রক্রিয়া যেখানে একটি ডেটা টাইপকে স্বাভাবিক উপায়ে (safely) অন্য টাইপে পরিবর্তন করা হয়। এটি Implicit (স্বয়ংক্রিয়) বা Explicit (স্পষ্টভাবে) হতে পারে।

🔹 Implicit Conversion (স্বয়ংক্রিয়):

কম্পাইলার নিজে থেকে নিরাপদভাবে টাইপ পরিবর্তন করে।
C/C++, Java এই ধরনের কনভার্সন সাপোর্ট করে।
🔹 Explicit Conversion (স্পষ্ট):

প্রোগ্রামারকে স্পষ্টভাবে টাইপ পরিবর্তন করতে হয়।
Golang-এ শুধুমাত্র Explicit Type Conversion অনুমোদিত।

🔹 উদাহরণ (Golang):

var x int = 10
var y float64 = float64(x)  // টাইপ কনভার্সন - int থেকে float-এ নিরাপদ পরিবর্তন


Golang-এর ক্ষেত্রে গুরুত্বপূর্ণ বিষয়:
Golang-এ Type Casting নেই, শুধুমাত্র Type Conversion আছে।
স্বয়ংক্রিয় টাইপ পরিবর্তন হয় না। আপনাকে নিজ হাতে টাইপ কনভার্ট করতে হবে।
এটি কোডের নিরাপত্তা (safety) এবং পরিষ্কারতা (clarity) নিশ্চিত করে।
*/

package main

import "fmt"

func main() {
	var x int = 42
	var y float64 = float64(x) // int থেকে float-এ কনভার্সন
	fmt.Println(y)             // আউটপুট: 42.0

	// taking the required
	// data into variables
	var totalsum int = 846
	var number int = 19
	var avg float32

	// explicit type conversion
	avg = float32(totalsum) / float32(number)

	// Displaying the result
	fmt.Printf("Average = %f\n", avg)
}

/*
কেন Golang-এ এই নিয়ম?
Golang স্পষ্টতা (clarity) এবং নিরাপত্তা (safety) নিশ্চিত করতে চায়।
Implicit Conversion অনেক সময় বাগ (bug) সৃষ্টি করতে পারে, তাই Go আপনাকে
সবসময় এক্সপ্লিসিট কনভার্সন করতে বাধ্য করে।
Golang-এ Type Conversion বাধ্যতামূলক হওয়ার ফলে, কোড পরিষ্কার এবং বাগ মুক্ত থাকে।
*/
