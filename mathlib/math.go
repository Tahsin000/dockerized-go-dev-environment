package mathlib

import "fmt"

// Exported variable
// কারণ Money capital M দিয়ে শুরু
var Money = 100

// Unexported variable
// কারণ secretMoney ছোট হাতের
var secretMoney = 500

// Exported function
// কারণ Add capital A দিয়ে শুরু
func Add(x int, y int) {
	z := x + y
	fmt.Println("mathlib.Add result:", z)
}

// Unexported function
// কারণ sum ছোট হাতের
func sum(x int, y int) {
	fmt.Println("sum result:", x+y)
}