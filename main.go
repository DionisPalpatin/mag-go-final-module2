package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size < 0 {
		return nil
	}

	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(math.MaxInt) + 1
	}

	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	dataLen := len(data)
	if dataLen == 0 {
		return -1
	}

	maxValue := data[0]
	for i := 1; i < dataLen; i++ {
		maxValue = max(maxValue, data[i])
	}

	return maxValue
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) < CHUNKS {
		return maximum(data)
	}

	localMaximums := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS
	var wg sync.WaitGroup

	wg.Add(CHUNKS)
	for i := range CHUNKS {
		start := i * chunkSize
		end := start + chunkSize

		if i == CHUNKS-1 {
			end = len(data)
		}

		go func() {
			defer wg.Done()
			localMaximums[i] = maximum(data[start:end])
		}()
	}
	wg.Wait()

	maxValue := localMaximums[0]
	for i := 1; i < CHUNKS; i++ {
		maxValue = max(maxValue, localMaximums[i])
	}

	return maxValue
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)

	array := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")

	start := time.Now()
	maxValue := maximum(array)
	elapsed := time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxValue, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)

	start = time.Now()
	maxValue = maxChunks(array)
	elapsed = time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxValue, elapsed)
}
