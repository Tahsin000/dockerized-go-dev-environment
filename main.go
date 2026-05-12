// package main

// import (
// "fmt"
// "net/http"
// )

// func main() {
// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// fmt.Fprintln(w, "Ami hotash! hello world")
// })

// fmt.Println("Server running on :8080")
// http.ListenAndServe(":8080", nil)
// }

// ------------------ class: 07

// package main

// func main (){
// 	fmt.Println("Ami hotash! hello world")
// }


// ------------------ class: 08

// package main

// import "fmt"

// func main (){
// 	a:=80;
// 	fmt.Println(a)
// 	/*
// 		int 
// 		float32
// 		bool
// 		string
// 	*/

// 	var x int = 90;
// 	fmt.Println(x)

// 	var intVar = 100
// 	fmt.Println(intVar)

// 	boolVar := true
// 	fmt.Println(boolVar)

// 	strVar := "Hello, Go!"
// 	fmt.Println(strVar)

// 	strVar = "Updated string value"
// 	fmt.Println(strVar)

// 	const pi = 3.14
// 	fmt.Println(pi)
// }



// ------------------ class: 09

// package main

// import "fmt"

// func main() {
// 	age := 20;
// 	gender := "male"


// // ------- if-else statement
// 	fmt.Printf("Your age is %d and your gender is %s.\n", age, gender)

// 	if (age >= 18) {
// 		fmt.Println("You are an adult.")
// 	} else {
// 		fmt.Println("You are a minor.")
// 	}

// // ------ switch case
// 	switch gender {
// 	case "male":
// 		fmt.Println("You are a man.")
// 	case "female":
// 		fmt.Println("You are a woman.")
// 	default:
// 		fmt.Println("Your gender is not specified.")
// 	}
// }

// ------------------ class: 10

// package main

// import "fmt"


// func add(x int, y int) int {
// 	return x + y
// }

// func display(v interface{}) {
// 	fmt.Printf("The values are: %v\n", v)
// }

// func main() {
// // ----- sum example
// 	a := 10
// 	b := 20
// 	sum := a + b
// 	fmt.Printf("The sum of %d and %d is %d.\n", a, b, sum)

// // ----- function example
// 	result := add(a, b);
// 	display(result);
// }


// ------------------ class: 11


// package main

// import "fmt"

// func getSum(x int, y int) int {
// 	return x + y
// }

// func sayHello(message string) {
// 	fmt.Println(message)
// }

// func main() {
// 	a:=10
// 	b:=20

// 	result := getSum(a, b)
// 	fmt.Printf("The sum of %d and %d is %d.\n", a, b, result)

// 	sayHello("Hello, Go!")
// }


// ------------------ class: 13

// package main

// import "fmt"

// func printWelcomeMessage() {
// 	fmt.Println("Welcome to the Application")
// }

// func getUserName() string {
// 	var name string

// 	fmt.Print("Enter your name: ")
// 	fmt.Scanln(&name)

// 	return name
// }

// func getTwoNumbers() (int, int) {
// 	var numberOne int
// 	var numberTwo int

// 	fmt.Print("Enter first number: ")
// 	fmt.Scanln(&numberOne)

// 	fmt.Print("Enter second number: ")
// 	fmt.Scanln(&numberTwo)

// 	return numberOne, numberTwo
// }

// func add(numberOne int, numberTwo int) int {
// 	return numberOne + numberTwo
// }

// func displayResult(name string, sum int) {
// 	fmt.Println("Hello", name)
// 	fmt.Println("Summation =", sum)
// }

// func printGoodbyeMessage() {
// 	fmt.Println("Thank you for using the application")
// 	fmt.Println("Goodbye")
// }

// func main() {
// 	printWelcomeMessage()

// 	name := getUserName()

// 	numberOne, numberTwo := getTwoNumbers()

// 	sum := add(numberOne, numberTwo)

// 	displayResult(name, sum)

// 	printGoodbyeMessage()
// }




// ------------------ class: 14

// package main

// import "fmt"

// // Global scope
// // a, b, and add() can be accessed from any function in this file/package.
// var a = 20
// var b = 30

// func add(x int, y int) {
// 	// Function/local scope
// 	// z only exists inside this add() function.
// 	z := x + y

// 	fmt.Println("sum:", z)
// }

// func main() {
// 	// Local scope of main()
// 	p := 30
// 	q := 40

// 	add(p, q) // 30 + 40 = 70

// 	add(a, b) // 20 + 30 = 50

// 	add(a, p) // 20 + 30 = 50

// 	// This will NOT work:
// 	// add(b, z)
// 	//
// 	// Why?
// 	// Because z was created inside add().
// 	// So z's scope is only inside add().
// 	// main() cannot access z.
// }



// ------------------ class: 15

// package main

// import "fmt"

// // Package scope / global-like variable
// // এই variable পুরো package-এর যেকোনো function থেকে access করা যাবে
// var appName = "Scope Demo"
// var version = 1

// func main() {
// 	fmt.Println("App:", appName)
// 	fmt.Println("Version:", version)

// 	// Function scope / local variable
// 	// x শুধু main function-এর ভিতরে access করা যাবে
// 	x := 18

// 	fmt.Println("x from main:", x)

// 	if x >= 18 {
// 		// if block scope
// 		// p শুধু এই if block-এর ভিতরে access করা যাবে
// 		p := 10

// 		fmt.Println("Inside if block")
// 		fmt.Println("p =", p)
// 		fmt.Println("x =", x)          // main function-এর variable পাওয়া যাবে
// 		fmt.Println("appName =", appName) // package scope variable পাওয়া যাবে
// 	}

// 	// এখানে p পাওয়া যাবে না
// 	// কারণ p শুধু if block-এর ভিতরে ছিল

// 	// fmt.Println(p) 
// 	// Error: undefined: p

// 	switch x {
// 	case 18:
// 		// switch case/block scope
// 		message := "You are exactly 18"
// 		fmt.Println(message)

// 	default:
// 		otherMessage := "You are not 18"
// 		fmt.Println(otherMessage)
// 	}

// 	// এখানে message পাওয়া যাবে না
// 	// কারণ message switch block-এর ভিতরে declare করা হয়েছে

// 	// fmt.Println(message)
// 	// Error: undefined: message

// 	printInfo()
// }

// func printInfo() {
// 	// এখানে appName পাওয়া যাবে কারণ এটা package scope
// 	fmt.Println("From printInfo:", appName)

// 	// এখানে x পাওয়া যাবে না
// 	// কারণ x main function-এর local variable

// 	// fmt.Println(x)
// 	// Error: undefined: x
// }



// ------------------ class: 15



// package main

// import (
// 	"fmt"

// 	"example.com/first-program/mathlib"
// )

// var a = 10
// var b = 20

// func main() {
// 	fmt.Println("Main package started")

// 	// Same package-এর function
// 	add(4, 7)

// 	// Different package-এর exported function
// 	mathlib.Add(5, 6)

// 	// Different package-এর exported variable
// 	fmt.Println("Money:", mathlib.Money)

// 	// এটা কাজ করবে না, কারণ sum ছোট হাতের
// 	// mathlib.sum(2, 3)

// 	// এটাও কাজ করবে না, কারণ secretMoney ছোট হাতের
// 	// fmt.Println(mathlib.secretMoney)
// 	// go mod init example.com/first-program
// }



// ------------------ class: 17

// package main

// import "fmt"

// // Global scope
// var a int = 10
// var b int = 20

// func add(x int, y int) {
// 	// Local scope of add function
// 	result := x + y

// 	printNumber(result)
// }

// func main() {
// 	add(a, b)
// }

// func printNumber(number int) {
// 	// Local scope of printNumber function
// 	fmt.Println(number)
// }




// ------------------ class: 18

// package main

// import "fmt"

// // Global variable
// var a int = 10

// func main() {
// 	age := 30

// 	if age > 18 {
// 		// This local variable shadows the global variable `a`
// 		var a int = 47

// 		fmt.Println(a) // prints local `a`
// 	}

// 	fmt.Println(a) // prints global `a`
// }



// ------------------ class: 19

// package main

// import "fmt"

// // add হচ্ছে একটি standard / named function
// // কারণ এই function-এর নাম আছে: add
// func add(a int, b int) {
// 	result := a + b
// 	fmt.Println("Sum:", result)
// }

// // greet হচ্ছে আরেকটি named function
// func greet(name string) {
// 	fmt.Println("Hello,", name)
// }

// // multiply function value return করছে
// func multiply(a int, b int) int {
// 	return a * b
// }

// func main() {
// 	// main নিজেও একটি standard / named function
// 	// কারণ এর নাম আছে: main

// 	add(4, 7)

// 	greet("Rahim")

// 	result := multiply(5, 6)
// 	fmt.Println("Multiply:", result)
// }


// ------------------ class: 20

// package main

// import "fmt"

// func init() {
//     fmt.Println("I am the first function that executes first")
// }

// func main() {
//     fmt.Println("Hello init function")
// }

// ------------------ class: 21

// package main

// import "fmt"

// // Named Function
// func unit() {
// 	fmt.Println("I will be called first")
// }

// // Named Function
// func add(a int, b int) {
// 	fmt.Println("Sum:", a+b)
// }

// func main() {

// 	// First function call
// 	unit()

// 	// Named function invocation
// 	add(5, 7)

// 	// -----------------------------
// 	// Anonymous Function
// 	// -----------------------------

// 	func(a int, b int) {

// 		c := a + b
// 		fmt.Println("Anonymous Function Sum:", c)

// 	}(4, 7)

// 	// -----------------------------
// 	// Another IIFE Example
// 	// -----------------------------

// 	func() {
// 		fmt.Println("This is an IIFE function")
// 	}()

// }

// ------------------ class: 22

// package main

// import "fmt"

// // Global Named Function
// func unit() {
// 	fmt.Println("I will be called first")
// }

// // Global Function Expression
// var globalAdd = func(a int, b int) {
// 	fmt.Println("Global Add:", a+b)
// }

// // Another Global Function
// func sum() {
// 	fmt.Println("Inside sum()")

// 	// Calling global function expression
// 	globalAdd(2, 4)
// }

// func main() {

// 	// First function call
// 	unit()

// 	// Calling another function
// 	sum()

// 	// --------------------------------
// 	// Local Function Expression
// 	// --------------------------------

// 	add := func(a int, b int) {

// 		c := a + b

// 		fmt.Println("Local Add:", c)

// 	}

// 	// Function Invocation
// 	add(4, 5)

// 	// --------------------------------
// 	// Shadowing Example
// 	// --------------------------------

// 	globalAdd := func(a int, b int) {
// 		fmt.Println("Shadowed Add:", a*b)
// 	}

// 	// This will call local shadowed function
// 	globalAdd(3, 4)

// }



// ------------------ class: 23

// package main

// import "fmt"

// // First-order function
// // শুধু normal data নিয়ে কাজ করে
// func Add(a int, b int) int {
// 	return a + b
// }

// // Higher-order function
// // কারণ এটা function কে parameter হিসেবে নিচ্ছে
// func ProcessOperation(
// 	a int,
// 	b int,
// 	operation func(int, int) int,
// ) int {

// 	// callback function execute
// 	result := operation(a, b)

// 	return result
// }

// // Function return করছে another function
// // এটাও Higher-order function
// func GetMultiplier() func(int, int) int {

// 	return func(x int, y int) int {
// 		return x * y
// 	}
// }

// func main() {

// 	// -----------------------------
// 	// Parameter vs Argument
// 	// -----------------------------

// 	// এখানে a,b হচ্ছে parameter
// 	// নিচে 2,5 হচ্ছে argument

// 	sum := Add(2, 5)

// 	fmt.Println("Add Result:", sum)

// 	// -----------------------------
// 	// First-class function
// 	// -----------------------------

// 	// Function কে variable এ assign করা হচ্ছে
// 	myFunc := Add

// 	fmt.Println("First-class Function:", myFunc(10, 20))

// 	// -----------------------------
// 	// Higher-order function
// 	// -----------------------------

// 	// anonymous function pass করছি
// 	// এটা callback function

// 	result := ProcessOperation(
// 		4,
// 		6,
// 		func(x int, y int) int {
// 			return x + y
// 		},
// 	)

// 	fmt.Println("Higher-order Result:", result)

// 	// -----------------------------
// 	// Function returning function
// 	// -----------------------------

// 	multiply := GetMultiplier()

// 	fmt.Println("Multiply Result:", multiply(3, 4))
// }




// ------------------ class: 24


package main

import "fmt"

// 🔵 Global variable (conceptually in DATA SEGMENT)
var globalValue = 10

// 🔵 Function (conceptually in CODE SEGMENT)
func add(x int, y int) int {
	// 🟡 Stack frame for add() is created here

	sum := x + y // local variable (STACK)

	fmt.Println("Inside add(), sum =", sum)

	// 🟡 Stack frame destroyed when function returns
	return sum
}

// 🔵 Another function (also CODE SEGMENT)
func init() {
	fmt.Println("init() runs first")
}

func main() {
	// 🟡 Stack frame for main() starts here

	fmt.Println("Global value:", globalValue)

	// First call → new stack frame for add(5,4)
	result1 := add(5, 4)
	fmt.Println("Result1:", result1)

	// Second call → new stack frame for add(10,3)
	result2 := add(globalValue, 3)
	fmt.Println("Result2:", result2)

	// 🟡 main() stack frame ends here
}



// ------------------ class: 25






