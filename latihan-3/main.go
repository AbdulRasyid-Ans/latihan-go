package main

import "fmt"

//Buat Function untuk mereturn bilangan prima dari 1 - max number
func main() {
	fmt.Println(listNumber(100, isPrime))
}

func isPrime(num int) (result bool) {
	if num == 1 {
		result = true
		return
	}
	count := 0
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			count++
		}
	}
	if count == 2 {
		result = true
	}
	return
}

func listNumber(maxNumber int, cb func(int) bool) (res []int) {
	for i := 1; i <= maxNumber; i++ {
		if cb(i) {
			res = append(res, i)
		}
	}
	return
}
