package leetcode

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	return "say " + strconv.Itoa(n-1) + " " + countAndSay(n-1)
}
