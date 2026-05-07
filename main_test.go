package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGenerateRandomElements tests the generateRandomElements function.
func TestGenerateRandomElements(t *testing.T) {
	t.Run("normal size", func(t *testing.T) {
		size := 100
		data := generateRandomElements(size)
		assert.Len(t, data, size)
		for i, v := range data {
			assert.Positive(t, v, "expected positive number at index %d", i)
		}
	})

	t.Run("zero size returns empty slice", func(t *testing.T) {
		data := generateRandomElements(0)
		assert.Empty(t, data)
	})

	t.Run("negative size returns empty slice", func(t *testing.T) {
		data := generateRandomElements(-5)
		assert.Empty(t, data)
	})

	t.Run("size of one", func(t *testing.T) {
		data := generateRandomElements(1)
		assert.Len(t, data, 1)
		assert.Positive(t, data[0])
	})
}

// TestMaximum tests the maximum function.
func TestMaximum(t *testing.T) {
	t.Run("empty slice returns 0", func(t *testing.T) {
		assert.Equal(t, 0, maximum([]int{}))
	})

	t.Run("single element", func(t *testing.T) {
		assert.Equal(t, 42, maximum([]int{42}))
	})

	t.Run("maximum is first element", func(t *testing.T) {
		assert.Equal(t, 9, maximum([]int{9, 3, 1, 5, 2}))
	})

	t.Run("maximum is last element", func(t *testing.T) {
		assert.Equal(t, 9, maximum([]int{1, 3, 2, 5, 9}))
	})

	t.Run("maximum is in the middle", func(t *testing.T) {
		assert.Equal(t, 9, maximum([]int{1, 3, 9, 5, 2}))
	})

	t.Run("all elements equal", func(t *testing.T) {
		assert.Equal(t, 7, maximum([]int{7, 7, 7, 7}))
	})

	t.Run("two elements", func(t *testing.T) {
		assert.Equal(t, 8, maximum([]int{4, 8}))
	})
}

// TestMaxChunks tests the maxChunks function.
func TestMaxChunks(t *testing.T) {
	t.Run("empty slice returns 0", func(t *testing.T) {
		assert.Equal(t, 0, maxChunks([]int{}))
	})

	t.Run("result matches sequential maximum", func(t *testing.T) {
		data := generateRandomElements(1_000)
		assert.Equal(t, maximum(data), maxChunks(data))
	})

	t.Run("large slice result matches sequential maximum", func(t *testing.T) {
		data := generateRandomElements(100_000)
		assert.Equal(t, maximum(data), maxChunks(data))
	})

	t.Run("single element", func(t *testing.T) {
		assert.Equal(t, 55, maxChunks([]int{55}))
	})

	t.Run("known values", func(t *testing.T) {
		data := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
		assert.Equal(t, 9, maxChunks(data))
	})
}
