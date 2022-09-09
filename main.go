package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func partOne() {
	rand.Seed(time.Now().UnixNano())
	x := 0
	fmt.Print("Please enter an integer: ")
	// Scan input from command line
	fmt.Scan(&x)

	// Init slice
	s := make([]int, 0)
	// Add values
	for i := 0; i < x; i++ {
		s = append(s, rand.Intn(100))
	}

	// Without using goroutines
	loop_sum := 0
	// Start timer
	loop_start := time.Now()
	// Sum using loop
	for _, num := range s {
		loop_sum += num
	}
	// End timer
	loop_end := time.Since(loop_start)

	// First print
	fmt.Println("Loop Sum:", loop_sum)
	fmt.Println("Loop Calculation time:", loop_end)

	// Using goroutines
	// Init channel
	c := make(chan int, 2)

	// Start timer
	parallel_start := time.Now()
	go calcSum(s[:len(s)/2], c)
	go calcSum(s[len(s)/2:], c)

	sum1 := <-c
	sum2 := <-c

	sum3 := sum1 + sum2

	// End timer
	parallel_end := time.Since(parallel_start)

	// Second Print
	fmt.Println("Parallel Sum:", sum3)
	fmt.Println("Parallel Calculation time:", parallel_end)

}

// Helper function for partOne()
func calcSum(s []int, c chan<- int) {
	sum := 0
	// Sum using loop
	for _, num := range s {
		sum += num
	}
	// Send sum through channel back to main function
	c <- sum
}

func partTwo() {
	rand.Seed(time.Now().UnixNano())
	x := 0
	fmt.Print("Please enter an integer: ")
	// Scan input from command line
	fmt.Scan(&x)

	// Init slice
	s := make([]int, 0)
	// Add values
	for i := 0; i < x; i++ {
		s = append(s, rand.Intn(100))
	}

	// Making a copy of slice s
	s2 := make([]int, len(s))
	copy(s2, s)

	// Start timer
	slice1_start := time.Now()
	// Sort Slice
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	// End timer
	slice1_end := time.Since(slice1_start)

	// Start timer
	slice2_start := time.Now()
	// Sort Slice
	sort.SliceStable(s2, func(i, j int) bool { return s2[i] < s2[j] })
	// End timer
	slice2_end := time.Since(slice2_start)

	// Print outputs
	fmt.Println("Slice Sort Time:", slice1_end)
	fmt.Println("SliceStable Sort Time:", slice2_end)
}

func main() {
	partOne()
	partTwo()
}
