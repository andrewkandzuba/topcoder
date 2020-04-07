package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TwoSums(t *testing.T) {
	assert.EqualValues(t, []int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
	assert.EqualValues(t, []int{0, 6}, twoSum([]int{2, 7, 11, 15, 0, 0, 1}, 3))
}

func Test_MedianOfTwoSortedArrays(t *testing.T) {
	assert.Equal(t, 1.5, findMedianSortedArrays([]int{1}, []int{2}))

	assert.Equal(t, 2.0, findMedianSortedArrays([]int{2}, []int{1, 3}))
	assert.Equal(t, 3.0, findMedianSortedArrays([]int{5}, []int{1, 3}))
	assert.Equal(t, 7.0, findMedianSortedArrays([]int{7}, []int{4, 8}))

	assert.Equal(t, 4.5, findMedianSortedArrays([]int{1}, []int{2, 7, 12}))
	assert.Equal(t, 5.5, findMedianSortedArrays([]int{4}, []int{2, 7, 12}))
	assert.Equal(t, 6.0, findMedianSortedArrays([]int{5}, []int{2, 7, 12}))
	assert.Equal(t, 10.0, findMedianSortedArrays([]int{19}, []int{2, 7, 13}))

	assert.Equal(t, 7.0, findMedianSortedArrays([]int{1}, []int{2, 7, 12, 15}))
	assert.Equal(t, 7.0, findMedianSortedArrays([]int{4}, []int{2, 7, 12, 15}))
	assert.Equal(t, 9.0, findMedianSortedArrays([]int{9}, []int{2, 7, 12, 15}))
	assert.Equal(t, 12.0, findMedianSortedArrays([]int{13}, []int{2, 7, 12, 15}))
	assert.Equal(t, 12.0, findMedianSortedArrays([]int{18}, []int{2, 7, 12, 15}))

	assert.Equal(t, 5.0, findMedianSortedArrays([]int{1}, []int{2, 3, 7, 12, 15}))
	assert.Equal(t, 5.5, findMedianSortedArrays([]int{4}, []int{2, 3, 7, 12, 15}))
	assert.Equal(t, 10.5, findMedianSortedArrays([]int{9}, []int{2, 7, 12, 11, 15}))
	assert.Equal(t, 12.5, findMedianSortedArrays([]int{13}, []int{2, 7, 12, 14, 15}))
	assert.Equal(t, 13.5, findMedianSortedArrays([]int{18}, []int{2, 7, 12, 15, 17}))

	assert.Equal(t, 2.0, findMedianSortedArrays([]int{1, 3}, []int{2}))
	assert.Equal(t, 3.0, findMedianSortedArrays([]int{1, 3}, []int{5}))
	assert.Equal(t, 7.0, findMedianSortedArrays([]int{4, 8}, []int{7}))

	assert.Equal(t, 6.0, findMedianSortedArrays([]int{4, 8}, []int{5, 7}))
	assert.Equal(t, 4.0, findMedianSortedArrays([]int{1, 2}, []int{6, 7}))
	assert.Equal(t, 3.0, findMedianSortedArrays([]int{1, 4}, []int{2, 7}))
	assert.Equal(t, 4.5, findMedianSortedArrays([]int{3, 6}, []int{2, 7}))
	assert.Equal(t, 8.0, findMedianSortedArrays([]int{9, 10}, []int{2, 7}))

	assert.Equal(t, 7.0, findMedianSortedArrays([]int{4, 8}, []int{5, 7, 11}))

	assert.Equal(t, 5.5, findMedianSortedArrays([]int{3, 4, 5, 6, 7, 8, 9}, []int{1, 2, 3, 9, 10}))

	assert.Equal(t, 3.0, findMedianSortedArrays([]int{}, []int{1, 2, 3, 9, 10}))
	assert.Equal(t, 1.0, findMedianSortedArrays([]int{}, []int{1}))
	assert.Equal(t, 1.0, findMedianSortedArrays([]int{1}, []int{}))
}

func Test_Palindrome(t *testing.T) {
	assert.True(t, isPalindrome(""))
	assert.True(t, isPalindrome("a"))
	assert.True(t, isPalindrome("aa"))
	assert.False(t, isPalindrome("abc"))
	assert.True(t, isPalindrome("aba"))
	assert.True(t, isPalindrome("ababa"))
	assert.False(t, isPalindrome("ababc"))
	assert.False(t, isPalindrome("abaabc"))
	assert.True(t, isPalindrome("abaaba"))
}

func Test_LongestPalindromeSubstring(t *testing.T) {
	assert.Equal(t, "bab", longestPalindrome("babad"))
	assert.Equal(t, "a", longestPalindrome("abcd"))
	assert.Equal(t, "aa", longestPalindrome("aabcd"))
	assert.Equal(t, "bb", longestPalindrome("abbcd"))
	assert.Equal(t, "bcb", longestPalindrome("aabcbbzd"))
	assert.Equal(t, "bcb", longestPalindrome("aazzbcb"))
	assert.Equal(t, "aaaaaaaaa", longestPalindrome("aaaaaaaaac"))
	assert.Equal(t, "ddtattarrattatdd", longestPalindrome("babaddtattarrattatddetartrateedredividerb"))
}

func Test_ZigZagConvention(t *testing.T) {
	assert.Equal(t, "", convert("PAYPALISHIRING", 0))
	assert.Equal(t, "PAYPALISHIRING", convert("PAYPALISHIRING", 1))
	assert.Equal(t, "PAHNAPLSIIGYIR", convert("PAYPALISHIRING", 3))
	assert.Equal(t, "PINALSIGYAHRPI", convert("PAYPALISHIRING", 4))
	assert.Equal(t, "PAYPALISHIRING", convert("PAYPALISHIRING", 14))
	assert.Equal(t, "ACB", convert("ABC", 2))
	assert.Equal(t, "ACBD", convert("ABCD", 2))
}

func Test_Atoi(t *testing.T) {
	assert.Equal(t, 0, myAtoi(""))
	assert.Equal(t, 0, myAtoi("    "))
	assert.Equal(t, 0, myAtoi("  $$ && () - "))
	assert.Equal(t, 0, myAtoi("  $$23 && () - "))
	assert.Equal(t, 0, myAtoi("  $$-23 && () - "))
	assert.Equal(t, 0, myAtoi("  $$- 23 && () - "))
	assert.Equal(t, 23, myAtoi("  23 && () - "))
	assert.Equal(t, -23, myAtoi(" -23 && () - "))
	assert.Equal(t, 23, myAtoi("  +23 && () - "))
	assert.Equal(t, 23, myAtoi("23"))
	assert.Equal(t, -23, myAtoi("-23"))
	assert.Equal(t, 0, myAtoi("words and 987"))
	assert.Equal(t, 2147483647, myAtoi(" 2147483647 dfsgdfg"))
	assert.Equal(t, 2147483647, myAtoi(" 2147483647000 dfsgdfg"))
	assert.Equal(t, -2147483648, myAtoi("  -2147483648 "))
	assert.Equal(t, -2147483648, myAtoi("  -2147483640008 "))
	assert.Equal(t, -2147483648, myAtoi("-91283472332"))
	assert.Equal(t, 0, myAtoi("+-2"))
}

func Test_MaxArea(t *testing.T) {
	assert.Equal(t, 0, maxArea([]int{0, 0}))
	assert.Equal(t, 1, maxArea([]int{1, 1}))
	assert.Equal(t, 0, maxArea([]int{0, 1}))
	assert.Equal(t, 1, maxArea([]int{0, 1, 2}))
	assert.Equal(t, 2, maxArea([]int{2, 1, 1}))
	assert.Equal(t, 2, maxArea([]int{1, 1, 1}))
	assert.Equal(t, 2, maxArea([]int{1, 2, 1}))
	assert.Equal(t, 2, maxArea([]int{1, 0, 1}))
	assert.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	assert.Equal(t, 6, maxArea([]int{1, 0, 0, 0, 0, 0, 8, 0, 0}))
	assert.Equal(t, 2, maxArea([]int{0, 0, 0, 2, 0, 1, 0, 0, 0}))
	assert.Equal(t, 0, maxArea([]int{0, 0, 0, 0, 0, 0, 0, 0, 0}))
	assert.Equal(t, 0, maxArea([]int{0, 0, 0, 0, 1, 0, 0, 0, 0}))
	assert.Equal(t, 17, maxArea([]int{2, 3, 4, 5, 18, 17, 6}))
	assert.Equal(t, 24, maxArea([]int{1, 3, 2, 5, 25, 24, 5}))
	assert.Equal(t, 1, maxArea([]int{0, 0, 0, 0, 0, 8, 1}))
	assert.Equal(t, 2, maxArea([]int{0, 0, 0, 0, 8, 0, 1}))
	assert.Equal(t, 1, maxArea([]int{1, 8, 0, 0, 0, 0, 0}))
	assert.Equal(t, 2, maxArea([]int{1, 0, 8, 0, 0, 0, 0}))
	assert.Equal(t, 48, maxArea([]int{8, 0, 1, 0, 0, 0, 8}))
}

func Test_TreeSum(t *testing.T) {
	assert.Equal(t, [][]int{}, threeSum([]int{0}))
	assert.Equal(t, [][]int{}, threeSum([]int{0, 0}))
	assert.Equal(t, [][]int{{0, 0, 0}}, threeSum([]int{0, 0, 0}))
	assert.Equal(t, [][]int{{-1, 0, 1}}, threeSum([]int{1, 0, -1}))
	assert.Equal(t, [][]int{{-1, 0, 1}}, threeSum([]int{0, 1, 0, -1}))
	assert.Equal(t, [][]int{{-100, 1, 99}, {-1, 0, 1}, {0, 0, 0}}, threeSum([]int{0, 1, 0, -1, -1, 0, 0, 1, 1, -100, 0, 99, 1}))
	assert.Equal(t, [][]int{{-1, 0, 1}}, threeSum([]int{0, 1, -1}))
	assert.Equal(t, [][]int{{-1, 0, 1}, {0, 0, 0}}, threeSum([]int{0, 1, 0, -1, 0}))
	assert.Equal(t, [][]int{{-100, 0, 100}, {-1, 0, 1}, {0, 0, 0}}, threeSum([]int{0, 1, 0, -1, 0, 100, -100}))
	assert.Equal(t, [][]int{{-1,-1,2},{-1,0,1}}, threeSum([]int{-1,0,1,2,-1,-4}))
	assert.Equal(t, [][]int{{-2,0,2},{-2,1,1}}, threeSum([]int{-2,0,1,1,2}))
	assert.Equal(t, [][]int{{-4,-2,6},{-4,0,4},{-4,1,3},{-4,2,2},{-2,-2,4},{-2,0,2}}, threeSum([]int{-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6}))
}

func TestLetterCombinations(t *testing.T) {
	assert.Equal(t, []string{}, letterCombinations(""))
	assert.Equal(t, []string{}, letterCombinations("0"))
	assert.Equal(t, []string{}, letterCombinations("1"))
	assert.Equal(t, []string{}, letterCombinations("00"))
	assert.Equal(t, []string{}, letterCombinations("01"))
	assert.Equal(t, []string{}, letterCombinations("10"))
	assert.Equal(t, []string{}, letterCombinations("11"))
	assert.Equal(t, []string{"a", "b", "c"}, letterCombinations("2"))
	assert.Equal(t, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, letterCombinations("23"))
	assert.Equal(t, []string{}, letterCombinations("123"))
}
