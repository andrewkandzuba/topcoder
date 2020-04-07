package leetcode

func letterCombinations(digits string) []string {
	var r = make([]string, 0)
	letterCombinationsImp(digits, "", &r)
	return r
}

var matrix = [][]byte{
	{},
	{},
	{'a', 'b', 'c'},
	{'d', 'e', 'f'},
	{'g', 'h', 'i'},
	{'j', 'k', 'l'},
	{'m', 'n', 'o'},
	{'p', 'q', 'r', 's'},
	{'t', 'u', 'v'},
	{'w', 'x', 'y', 'z'},
}

func letterCombinationsImp(digits string, c string, r *[]string) {
	if len(digits) == 0 {
		if len(c) > 0 {
			*r = append(*r, c)
		}
		return
	}
	i := digits[0] - '0'
	if len(matrix[i]) > 0 {
		for _, b := range matrix[i] {
			letterCombinationsImp(digits[1:], c+string(b), r)
		}
	}
}
