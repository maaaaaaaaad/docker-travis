func solution(e int, starts []int) []int {
	freq := make([]int, e+1)
	for i := 1; i <= e; i++ {
		for j := i; j <= e; j += i {
			freq[j]++
		}
	}
	maxNum := make([]int, e+1)
	maxFreq := freq[1]
	maxNum[1] = 1
	for i := 2; i <= e; i++ {
		if freq[i] >= maxFreq {
			maxFreq = freq[i]
			maxNum[i] = i
		} else {
			maxNum[i] = maxNum[i-1]
		}
	}
	result := make([]int, len(starts))
	for i, start := range starts {
		maxFreqInRange := freq[start]
		maxNumInRange := start
		for j := start; j <= e; j++ {
			if freq[j] > maxFreqInRange {
				maxFreqInRange = freq[j]
				maxNumInRange = j
			}
		}
		result[i] = maxNumInRange
	}

	return result
}
