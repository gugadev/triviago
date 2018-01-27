package helpers

import "strconv"

/*
MapAtoi map a slice of strings to ints
this function is used to convert the
query params (strings) to int
*/
func MapAtoi(vs []string, f func(string) int) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

/*
ApplyMapAtoi apply the map function
*/
func ApplyMapAtoi(v string) int {
	val, _ := strconv.Atoi(v) // nolint: gas,errcheck
	return val
}
