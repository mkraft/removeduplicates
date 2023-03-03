package main

func RemoveDuplicates(in []int) []int {
	var indexLastUnique int
	var inRepeatMode bool
	numUniqueValues := 1
	length := len(in) - 1
	for i := 0; i < length; i++ {
		a, b := in[i], in[i+1]
		if a == b {
			if !inRepeatMode {
				inRepeatMode = true
				indexLastUnique = i
			}
		} else {
			numUniqueValues++
			if inRepeatMode {
				for k := 0; k < i-indexLastUnique; k++ {
					for j := indexLastUnique + 1; j < length; j++ { // use `length` instead of `len(in)-1` to minimize swaps
						tmp := in[j]
						in[j] = in[j+1]
						in[j+1] = tmp
					}
					length-- // decrement length to avoid re-scanning duplicates
				}
				inRepeatMode = false
				i = indexLastUnique
			}
		}
	}
	return in[:numUniqueValues]
}
