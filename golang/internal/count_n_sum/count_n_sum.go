package main

import (
	"log"
	"sync"
)

func main() {
	incoming := []int{10, -67}
	log.Printf("%+v", CountNSum(incoming))
}

var wg = sync.WaitGroup{}

func CountNSum(numbers []int) []int {
	if len(numbers) < 1 {
		return []int{}
	}

	resultCh := make(chan int, 2)
	posValCh := make(chan int)
	negValCh := make(chan int)

	wg.Add(2)
	go negativesHandler(resultCh, negValCh)
	go positivesHandler(resultCh, posValCh)

	for _, val := range numbers {
		if val > 0 {
			posValCh <- val
		} else {
			negValCh <- val
		}
	}
	close(posValCh)
	close(negValCh)
	wg.Wait()
	close(resultCh)
	resp := []int{0, 0}
	for v := range resultCh {
		idx := 0
		if v < 0 {
			idx = 1
		}
		resp[idx] = v
	}
	return resp
}

func positivesHandler(resultCh chan<- int, inputsCh <-chan int) {
	count := 0
	for range inputsCh {
		count++
	}
	resultCh <- count
	wg.Done()
}

func negativesHandler(resultCh chan<- int, inputsCh <-chan int) {
	sum := 0
	for val := range inputsCh {
		sum += val
	}
	resultCh <- sum
	wg.Done()
}
