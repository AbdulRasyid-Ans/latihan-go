package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	var arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	var wg sync.WaitGroup

	for i, a := range arr {
		wg.Add(3)
		go process1(i, a, 8, &wg)
		go process2(i, a, 8, &wg)
		go process3(i, a, 8, &wg)
	}

	wg.Wait()

	last := time.Since(now)
	fmt.Println("time for execution :", last.Seconds())

}

func process1(index, num, find int, wg *sync.WaitGroup) {
	time.Sleep(500 * time.Millisecond)
	if num == find {
		fmt.Println("process 1 done... ")
		fmt.Println("find :", num)
	}
	wg.Done()
}

func process2(index, num, find int, wg *sync.WaitGroup) {
	time.Sleep(500 * time.Millisecond)
	if num == find {
		fmt.Println("process 2 done... ")
		fmt.Println("find :", num)
	}
	wg.Done()
}

func process3(index, num, find int, wg *sync.WaitGroup) {
	time.Sleep(500 * time.Millisecond)
	if num == find {
		fmt.Println("process 3 done... ")
		fmt.Println("find :", num)
	}
	wg.Done()
}
