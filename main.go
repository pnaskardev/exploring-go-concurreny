package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// func workerPoolBuffered(numbers []int) (time.Duration, int) {
// 	start := time.Now()

// 	workerCount := runtime.NumCPU()
// 	jobs := make(chan int, workerCount)
// 	results := make(chan int, workerCount)

// 	var wg sync.WaitGroup
// 	var processed int64

// 	for range workerCount {
// 		wg.Go(func() {
// 			for n := range jobs {
// 				results <- mockAPICall(n) // rarely blocks
// 				atomic.AddInt64(&processed, 1)
// 			}
// 		})
// 	}

// 	go func() {
// 		for _, n := range numbers {
// 			jobs <- n
// 		}
// 		close(jobs)
// 	}()

// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	// Same slow collector
// 	for range results {
// 		// time.Sleep(2 * time.Millisecond)
// 	}

// 	return time.Since(start), int(processed)
// }

// func workerPoolUnbuffered(numbers []int) (time.Duration, int) {
// 	start := time.Now()

// 	workerCount := runtime.NumCPU()
// 	jobs := make(chan int)
// 	results := make(chan int)

// 	var wg sync.WaitGroup
// 	var processed int64

// 	// Workers
// 	for range workerCount {
// 		wg.Go(func() {
// 			for n := range jobs {
// 				results <- mockAPICall(n) // BLOCKS until collector receives
// 				atomic.AddInt64(&processed, 1)
// 			}
// 		})
// 	}

// 	// Producer
// 	go func() {
// 		for _, n := range numbers {
// 			jobs <- n
// 		}
// 		close(jobs)
// 	}()

// 	// Closer
// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	// ❗ Slow collector
// 	for range results {
// 		// time.Sleep(2 * time.Millisecond)
// 	}

// 	return time.Since(start), int(processed)
// }

func mockAPICall(n int) int {
	// Simulate network latency (50ms)
	time.Sleep(5 * time.Millisecond)
	return n * n
}

func sequential(numbers []int) (time.Duration, int) {
	start := time.Now()

	processed := 0
	// results := []int{}

	for _, n := range numbers {
		// Same work
		// results = append(results, n*n)
		_ = mockAPICall(n)
		processed++
	}

	return time.Since(start), processed
}

func workerPool(numbers []int) (time.Duration, int) {
	start := time.Now()

	workerCount := runtime.NumCPU()
	jobs := make(chan int, workerCount)
	results := make(chan int, workerCount)

	var wg sync.WaitGroup
	var processed int64

	// Workers
	for range workerCount {

		wg.Go(func() {
			for n := range jobs {
				results <- mockAPICall(n)
				atomic.AddInt64(&processed, 1)
			}
		})
	}

	// Producer
	go func() {
		for _, n := range numbers {
			jobs <- n
		}
		close(jobs)
	}()

	// Closer
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collector
	for range results {
		// drain results
	}

	return time.Since(start), int(processed)
}

func main() {
	numbers := []int{}
	for i := range 2000 {
		numbers = append(numbers, i+1)
	}

	seqTime, seqCount := sequential(numbers)

	wpTime, wpCount := workerPool(numbers)

	// bufTime, bufCount := workerPoolBuffered(numbers)
	// unbufTime, unbufCount := workerPoolUnbuffered(numbers)

	fmt.Printf("Numbers           : %d\n", len(numbers))
	fmt.Printf("CPU cores         : %d\n\n", runtime.NumCPU())

	fmt.Println("Sequential")
	fmt.Printf("Jobs processed    : %d\n", seqCount)
	fmt.Printf("Total time        : %v\n", seqTime)
	fmt.Printf("Throughput        : %.2f jobs/sec\n\n",
		float64(seqCount)/seqTime.Seconds(),
	)

	fmt.Println("WORKER POOL")
	fmt.Printf("Jobs processed    : %d\n", wpCount)
	fmt.Printf("Total time        : %v\n", wpTime)
	fmt.Printf("Throughput        : %.2f jobs/sec\n",
		float64(wpCount)/wpTime.Seconds(),
	)

	// fmt.Println("UNBUFFERED")
	// fmt.Printf("Jobs processed    : %d\n", unbufCount)
	// fmt.Printf("Total time        : %v\n", unbufTime)
	// fmt.Printf("Throughput        : %.2f jobs/sec\n\n",
	// 	float64(unbufCount)/unbufTime.Seconds(),
	// )

	// fmt.Println("BUFFERED")
	// fmt.Printf("Jobs processed    : %d\n", bufCount)
	// fmt.Printf("Total time        : %v\n", bufTime)
	// fmt.Printf("Throughput        : %.2f jobs/sec\n\n",
	// 	float64(bufCount)/bufTime.Seconds(),
	// )
}
