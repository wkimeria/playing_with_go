package main

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {

	max, err1 := strconv.Atoi(os.Args[1])
	if err1 != nil {
		return
	}

	threads, err2 := strconv.Atoi(os.Args[2])
	if err2 != nil {
		return
	}

	iters, err3 := strconv.Atoi(os.Args[3])
	if err3 != nil {
		return
	}

	itersThreads, err4 := strconv.Atoi(os.Args[4])
	if err4 != nil {
		return
	}

	currentMax := max
	currentThreads := threads

	for i := 0; i <= iters; i++ {

		for x := 0; x <= itersThreads; x++ {
			computePrimesForRange(currentMax, currentThreads)
			currentThreads = currentThreads * 2
		}
		currentMax = currentMax + max
		currentThreads = threads
	}
}

func computePrimesForRange(max int, threads int) {
	batchSize := max / threads

	start := time.Now()

	limits := make(map[int]int)

	lowerLimit := 2

	var wg sync.WaitGroup

	for i := batchSize; i <= max; i += batchSize {
		limits[lowerLimit] = i
		lowerLimit = i
	}

	for k, v := range limits {

		wg.Add(1)
		go func(lowerLimit int, upperLimit int) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			ComputePrimes(lowerLimit, upperLimit)
		}(k, v)
	}

	wg.Wait()

	elapsed := time.Since(start)
	log.Printf("Took %s threads = %s max = %s", elapsed, strconv.Itoa(threads), strconv.Itoa(max))
}

func ComputePrimes(lowerLimit int, upperLimit int) {
	for x := lowerLimit; x <= upperLimit; x++ {
		isPrime := true
		for y := 2; y < ((x / 2) + 1); y++ {
			if x%y == 0 {
				isPrime = false
				break
			}
		}
		if isPrime == true {
			//log.Printf("%s", strconv.Itoa(x))
		}
	}
}
