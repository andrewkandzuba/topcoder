package leetcode

func myAtoi(str string) int {
	const MIN_INT int = -2147483648
	const MAX_INT = 2147483647

	n := len(str)

	var i = 0
	for ; i < n && int(str[i])-' ' == 0; i++ {
	}
	if n == i {
		return 0
	}

	ch := int(str[i])
	if (ch-'0' < 0 || ch-'9' > 0) && ch-'-' != 0 && ch-'+' != 0 {
		return 0
	}
	negative := false
	if ch-'-' == 0 {
		negative = true
		i++
	}
	if ch-'+' == 0 {
		i++
	}
	if i == n {
		return 0
	}
	ch = int(str[i])
	if ch-'0' < 0 || ch-'9' > 0 {
		return 0
	}

	res := ch - '0'
	for i++; i < n && int(str[i])-'0' >= 0 && int(str[i])-'9' < 10; i++ {
		res = res*10 + int(str[i]) - '0'
		if negative && -res < MIN_INT {
			return MIN_INT
		}
		if !negative && res > MAX_INT {
			return MAX_INT
		}
	}

	if negative {
		return -res
	}

	return res
}
