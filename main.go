package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates a slice of random positive integers of the given size.
// Returns an empty slice if size is zero or negative.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(1_000_000) + 1
	}
	return data
}

// maximum returns the maximum value in the given slice.
// Returns 0 if the slice is empty.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	maxVal := data[0]
	for _, v := range data[1:] {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

// maxChunks divides the slice into CHUNKS parts, finds the maximum in each part
// concurrently using goroutines, and returns the overall maximum.
// Returns 0 if the slice is empty.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkSize := len(data) / CHUNKS
	maxValues := make([]int, CHUNKS)

	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}

		chunk := data[start:end]
		idx := i

		go func(chunk []int, idx int) {
			defer wg.Done()
			maxValues[idx] = maximum(chunk)
		}(chunk, idx)
	}

	wg.Wait()

	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	maxSingle := maximum(data)
	elapsedSingle := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxSingle, elapsedSingle)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	maxParallel := maxChunks(data)
	elapsedParallel := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxParallel, elapsedParallel)
}
