package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// create a channel for results
	resultsChan := make(chan bool, 1)
	wg.Add(1)
	octet := "01"
	go OctetValidator(resultsChan, octet)
	wg.Wait()
	close(resultsChan)

	v, _ := <-resultsChan
	fmt.Printf("octet '%s' is '%t'-y\n", octet, v)
}

func OctetValidator(ch chan<- bool, octet string) {
	// invalidate octets with leading zeroes unless they are the exact value '0'
	matched := (regexp.MustCompile(`^0$|^[^0]\d{0,2}$`)).MatchString(octet)
	// convert to int for easier validation of 1-256 range
	oct, err := strconv.Atoi(octet)
	ch <- err == nil && matched && 0 <= oct && oct <= 256
	wg.Done()
}

func IpValidation(str string) bool {
	octets := strings.Split(str, ".")
	if len(octets) != 4 {
		return false
	}

	ch := make(chan bool, len(octets))
	wg.Add(len(octets))
	for _, v := range octets {
		go OctetValidator(ch, v)
	}
	wg.Wait()
	close(ch)

	results := true
	for chVal := range ch {
		results = results && chVal
	}
	return results
}
