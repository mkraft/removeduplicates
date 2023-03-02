package main

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicatesExplicitInputs(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{input: []int{0, 0}, expect: []int{0}},
		{input: []int{0, 0, 0}, expect: []int{0}},
		{input: []int{1, 2, 2, 2, 3}, expect: []int{1, 2, 3}},
		{input: []int{1, 2, 2, 3}, expect: []int{1, 2, 3}},
		{input: []int{0, 1, 1, 2, 2, 3}, expect: []int{0, 1, 2, 3}},
		{input: []int{0, 0, 0, 1, 1, 1, 2, 2, 2, 3, 3, 3}, expect: []int{0, 1, 2, 3}},
		{input: []int{0, 0, 0, 1}, expect: []int{0, 1}},
		{input: []int{0, 0, 0, 1}, expect: []int{0, 1}},
		{input: []int{0, 1, 1, 1}, expect: []int{0, 1}},
		{input: []int{0, 0, 1, 1, 1}, expect: []int{0, 1}},
		{input: []int{0, 0, 0, 1, 1, 1, 2, 2, 2}, expect: []int{0, 1, 2}},
	}
	for _, tc := range tests {
		actual := RemoveDuplicates(tc.input)
		require.Equal(t, tc.expect, actual)
	}
}

func TestRemoveDuplicatesRandomInputs(t *testing.T) {
	random := make([]int, 100)
	for i := range random {
		random[i] = rand.Int() % 40
	}
	inputValsMap := map[int]bool{}
	var inputDupCount int
	for _, val := range random {
		if _, hasKey := inputValsMap[val]; hasKey {
			inputDupCount++
		}
		inputValsMap[val] = true
	}
	sort.Ints(random)
	actual := RemoveDuplicates(random)

	t.Run("results are a length of input - duplicate count", func(t *testing.T) {
		if len(actual) != len(random)-inputDupCount {
			t.Fail()
		}
	})

	t.Run("results have no duplicates", func(t *testing.T) {
		actualMap := make(map[int]bool, len(actual))
		for _, item := range actual {
			if _, hasKey := actualMap[item]; hasKey {
				t.Fail()
			}
			actualMap[item] = true
		}
	})

	t.Run("results are sorted", func(t *testing.T) {
		copyOfActual := make([]int, len(actual))
		copy(copyOfActual, actual)
		sort.Ints(actual)
		for i := 0; i < len(copyOfActual); i++ {
			if copyOfActual[i] != actual[i] {
				t.Fail()
			}
		}
	})
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RemoveDuplicates([]int{0, 0, 0, 1, 1, 1, 2, 2, 2})
	}
}
