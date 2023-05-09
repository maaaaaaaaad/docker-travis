package main

import "fmt"

func generateParenthesis(n int) []string {
	var result []string
	var generate func(left, right int, str string)
	generate = func(left, right int, str string) {
		if left == 0 && right == 0 {
			result = append(result, str)
			return
		}
		if left > 0 {
			generate(left-1, right+1, str+"(")
		}
		if right > 0 {
			generate(left, right-1, str+")")
		}
	}

	generate(n, 0, "")

	return result
}

func solution(n int) int {
	count := 0
	for _, s := range generateParenthesis(n) {
		if isValid(s) {
			count++
		}
	}
	return count
}

func isValid(s string) bool {
	stack := []rune{}
	for _, c := range s {
		if c == '(' {
			stack = append(stack, c)
		} else if len(stack) > 0 && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func main() {
	n := 3
	result := solution(n)
	fmt.Printf("%v result: %v", n, result)
}
