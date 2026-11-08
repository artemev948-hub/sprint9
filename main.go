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

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	maxValue := data[0]

	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkCount := CHUNKS

	if len(data) < chunkCount {
		chunkCount = len(data)
	}

	maxValues := make([]int, chunkCount)

	var wg sync.WaitGroup

	wg.Add(chunkCount)

	for i := 0; i < chunkCount; i++ {
		start := i * len(data) / chunkCount
		end := (i + 1) * len(data) / chunkCount

		go func(index, start, end int) {
			defer wg.Done()
			maxValues[index] = maximum(data[start:end])
		}(i, start, end)
	}
	wg.Wait()

	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")

	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)

	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
