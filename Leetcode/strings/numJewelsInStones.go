package strings

import "strings"

func NumJewelsInStones(J string, S string) int {
	js := strings.Split(J, "")
	sum := 0
	for _, j := range js {
		c := strings.Count(S, j)
		sum += c
	}
	return sum
}
