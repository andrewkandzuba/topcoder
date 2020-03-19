package leetcode

func longestPalindrome(s string) string {

	n := len(s)
	if n <= 1 {
		return s
	}

	var table = make([][]bool, n)

	for i := 0; i < n-1; i++ {
		table[i] = make([]bool,n)
		table[i][i] = true
	}
	var maxLength = 1

	var start = 0
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			table[i][i+1] = true
			maxLength = 2
			start = i
		}
	}

	for k := 3; k <= n; k++ {
		for i := 0; i < n-k+1; i++ {
			j := i + k - 1
			if table[i+1][j-1] && s[i] == s[j] {
				table[i][j] = true
				if k > maxLength {
					start = i
					maxLength = k
				}
			}
		}
	}
	return s[start : start+maxLength]
}

func isPalindrome(s string) bool {
	var n = len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}
