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





