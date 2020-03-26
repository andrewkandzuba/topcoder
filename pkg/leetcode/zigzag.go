package leetcode

import (
	"bytes"
)

func convert(s string, numRows int) string {
	if numRows == 0 {
		return ""
	}

	var n = len(s)
	var edgeLen = 0
	if numRows > 2 {
		edgeLen = numRows - 2
	}
	var maxZigLength = numRows + edgeLen
	var numsCols = n/numRows + 1

	var b bytes.Buffer
	for i := 0; i < numRows; i++ {
		var currZigLength = maxZigLength - 2*i
		for j := 0; j < numsCols; j++ {
			p1 := i + j*maxZigLength
			if p1 < n {
				b.WriteByte(s[p1])
			}
			if i > 0 && i < numRows-1 {
				p2 := p1 + currZigLength
				if p2 < n {
					b.WriteByte(s[p2])
				}
			}
		}
	}
	return b.String()
}