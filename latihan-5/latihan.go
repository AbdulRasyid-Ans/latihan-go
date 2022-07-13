package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	defer fmt.Println("Fungsi selesai")
	time.Sleep(500 * time.Millisecond)
	var arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	for i, a := range arr {
		go process1(i, a, 8, c1)
		go process2(i, a, 8, c2)
		go process3(i, a, 8, c3)
	}

	select {
	case msg := <-c1:
		fmt.Println("c1", msg)
		fmt.Println("process 1 done... ")
		break
	case msg := <-c2:
		fmt.Println("c2", msg)
		fmt.Println("process 2 done... ")
		break
	case msg := <-c3:
		fmt.Println("c3", msg)
		fmt.Println("process 3 done... ")
		break
	}
	fmt.Println("time for execution :", time.Since(now).Seconds())
}

func process1(index, num, find int, c chan string) {
	time.Sleep(500 * time.Millisecond)
	if num == find {
		// fmt.Println("find :", num)
		// c <- "process 1 done... "
		c <- fmt.Sprintln("find :", num)
	}
}

func process2(index, num, find int, c chan string) {
	time.Sleep(500 * time.Millisecond)
	if num == find {
		// fmt.Println("find :", num)
		// c <- "process 2 done... "
		c <- fmt.Sprintln("find :", num)
	}
}

func process3(index, num, find int, c chan string) {
	time.Sleep(500 * time.Millisecond)
	if num == find {
		// fmt.Println("find :", num)
		// c <- "process 3 done... "
		c <- fmt.Sprintln("find :", num)
	}
}
