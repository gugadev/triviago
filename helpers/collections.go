package helpers

import "strconv"

func MapAtoi(vs []string, f func(string) int) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func ApplyMapAtoi(v string) int {
	val, _ := strconv.Atoi(v)
	return val
}
