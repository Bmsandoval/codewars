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

	// create a channel for the two handlers to place their responses
	resultCh := make(chan int, 2)
	// create a channel to pipe positive inputs to appropriate handler
	posValCh := make(chan int)
	// create a channel to pipe negative inputs to appropriate handler
	negValCh := make(chan int)

	// wait group will be used to track when the handlers are done processing inputs
	wg.Add(2)
	// launch our two handler subroutines in the background
	go negativesHandler(resultCh, negValCh)
	go positivesHandler(resultCh, posValCh)

	// loop over the numbers and send positives and negatives to their rightful place
	for _, val := range numbers {
		if val > 0 {
			posValCh <- val
		} else {
			negValCh <- val
		}
	}
	// tell the positives handler that we are done sending inputs
	close(posValCh)
	// tell the negatives handler the same
	close(negValCh)
	// wait for the two handlers to finish processing their inputs
	wg.Wait()
	// close the channel of results (so we can freely 'range' over it)
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

// positivesHandler accepts any number or positive values and will return a count of them
func positivesHandler(resultCh chan<- int, inputsCh <-chan int) {
	count := 0
	// handle inputs until the channel is empty&closed
	for range inputsCh {
		count++
	}
	// post results when you have them
	resultCh <- count
	// notify main thread that we are done
	wg.Done()
}

// negativesHandler accepts any number of negative values and will return their sum
func negativesHandler(resultCh chan<- int, inputsCh <-chan int) {
	sum := 0
	// handle inputs until the channel is empty&closed
	for val := range inputsCh {
		sum += val
	}
	// post results when you have them
	resultCh <- sum
	// notify main thread that we are done
	wg.Done()
}
