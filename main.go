func solution(e int, starts []int) []int {
	freq := make([]int, e+1)
	for i := 1; i <= e; i++ {
		for j := i; j <= e; j += i {
			freq[j]++
		}
	}

	result := make([]int, len(starts))
	for i, start := range starts {
		maxFreq := freq[start]
		maxNum := start
		for j := start; j <= e; j++ {
			if freq[j] > maxFreq {
				maxFreq = freq[j]
				maxNum = j
			}
		}
		result[i] = maxNum
	}

	return result
}
