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
					for j := indexLastUnique + 1; j < len(in)-1; j++ {
						tmp := in[j]
						in[j] = in[j+1]
						in[j+1] = tmp
					}
					length--
				}
				inRepeatMode = false
				i = indexLastUnique
			}
		}
	}
	return in[:numUniqueValues]
}
