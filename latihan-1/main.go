package main

import "fmt"

/*
	gender = male || female

	age = int

	isCanWork?

	1. male, umur > 17
	2. female, umur > 20

	"gender male boleh bekerja ? true"
*/

// ====== My Answer ======
func main() {
	gender := "female"
	age := 20

	isCanWork(gender, age)
}

func isCanWork(gender string, age int) {
	var canWork bool = false
	if !(gender == "male" || gender == "female") {
		fmt.Println("gender harus male atau female")
		return
	}
	if gender == "male" && age > 17 {
		canWork = true
	}
	if gender == "female" && age > 20 {
		canWork = true
	}

	fmt.Printf("gender %s boleh bekerja ? %t \n", gender, canWork)
}

// ====== Instructor Answer ======
/*
func main() {
	gender := "female"
	age := 20

	isMale := gender == "male"

	isCanWork := (isMale && age > 17) || (!isMale && age > 20)

	fmt.Printf("gender %s boleh bekerja ? %t \n", gender, isCanWork)
}
*/
