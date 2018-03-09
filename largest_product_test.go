package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestProductNotFound(t *testing.T) {
	t.Parallel()
	found, _, _ := LargestProduct()
	assert.False(t, found)
}

func TestLargestProductNotFound2(t *testing.T) {
	t.Parallel()
	samples := [][]int{{0}, {-3}, {-1, -2}}
	for i, sample := range samples {
		sample := sample
		t.Run(fmt.Sprintf("sample %d", i), func(t *testing.T) {
			t.Parallel()
			found, _, _ := LargestProduct(sample...)
			assert.False(t, found)
		})
	}
}

type largestProductFoundSample struct {
	Numbers []int
	Index0  int
	Index1  int
}

func TestLargestProductFound(t *testing.T) {
	t.Parallel()
	samples := []largestProductFoundSample{
		{Numbers: []int{10, -2, 2, 3, 0, 1}, Index0: 0, Index1: 3},
		{Numbers: []int{1, 3, -2, -10}, Index0: 1, Index1: 0},
		{Numbers: []int{0, -5, 2, 3, 1, -2, 3}, Index0: 3, Index1: 6},
		{Numbers: []int{0, -5, 2, 3, 1, -2, 4}, Index0: 6, Index1: 3},
	}
	for i, sample := range samples {
		sample := sample
		t.Run(fmt.Sprintf("sample %d", i), func(t *testing.T) {
			t.Parallel()
			found, i0, i1 := LargestProduct(sample.Numbers...)
			assert.True(t, found)
			assert.Equal(t, sample.Index0, i0, i1)
			t.Logf("found=%v n[%d]*n[%d] = %d*%d = %d", found, i0, i1,
				sample.Numbers[i0], sample.Numbers[i1], sample.Numbers[i0]*sample.Numbers[i1])
		})
	}
}
