package algorithm

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

const LARGEST_PRIME = 10_000_000

var cores int
var primeNumbers []int
var m sync.Mutex
var wg sync.WaitGroup

func SieveOfEratosthenes(n int) []int {
	primes := make([]bool, n+1)
	for i := 0; i < n+1; i++ {
		primes[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if primes[p] == true {
			for i := p * p; i <= n; i += p {
				primes[i] = false
			}
		}
	}

	var primeNumbers []int
	for p := 2; p <= n; p++ {
		if primes[p] == true {
			primeNumbers = append(primeNumbers, p)
		}
	}

	return primeNumbers
}

func primeBetween(prime []int, low, high int) {
	defer wg.Done()
	limit := high - low
	segment := make([]bool, limit+1)
	for i := 0; i < len(segment); i++ {
		segment[i] = true
	}

	for i := 0; i < len(prime); i++ {
		lowlimit := int(math.Floor(float64(low)/float64(prime[i])) * float64(prime[i]))
		if lowlimit < low {
			lowlimit += prime[i]
		}
		for j := lowlimit; j < high; j += prime[i] {
			segment[j-low] = false
		}
	}

	m.Lock()
	for i := 0; i < high; i++ {
		if segment[i-low] == true {
			primeNumbers = append(primeNumbers, i)
		}
	}
	m.Unlock()
}

func SegmentedSieve(n int) {
	limit := int(math.Floor(float64(n) / float64(cores)))
	prime := SieveOfEratosthenes(limit)
	for i := 0; i < len(prime); i++ {
		primeNumbers = append(primeNumbers, prime[i])
	}
	for low := limit; low < n; low += limit {
		high := low + limit
		if high >= n {
			high = n
		}
		wg.Add(1)
		go primeBetween(prime, low, high)
	}
	wg.Wait()
}

func Ex() {
	cores = runtime.NumCPU()
	start := time.Now()
	SegmentedSieve(LARGEST_PRIME)
	elapsed := time.Since(start)
	fmt.Println("\nComputation time for concurrent: ", elapsed)
	fmt.Println("Number of primes = ", len(primeNumbers))
}
