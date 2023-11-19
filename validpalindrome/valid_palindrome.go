package validpalindrome

import (
	"strings"
	"unicode"
)

// A phrase is a palindrome if,
// after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters,
// it reads the same forward and backward.
// Alphanumeric characters include letters and numbers.
//
// Given a string s, return true if it is a palindrome, or false otherwise.
//
// Example 1:
//
//	Input: s = "A man, a plan, a canal: Panama"
//	Output: true
//	Explanation: "amanaplanacanalpanama" is a palindrome.
//
// Example 2:
//
//	Input: s = "race a car"
//	Output: false
//	Explanation: "raceacar" is not a palindrome.
//
// Example 3:
//
//	Input: s = " "
//	Output: true
//	Explanation: s is an empty string "" after removing non-alphanumeric characters.
//	Since an empty string reads the same forward and backward, it is a palindrome.
//
// Constraints:
//
//	1 <= s.length <= 2 * 105
//	s consists only of printable ASCII characters.
func isPalindrome(s string) bool {
	// empty string is a palindrome
	if len(s) == 0 {
		return true
	}

	s = strings.ToLower(s)
	isAlphaNum := func(c byte) bool {
		return c >= 'a' && c <= 'z' || c >= '0' && c <= '9'
	}

	for {
		// search for the next first element
		for len(s) > 1 && !isAlphaNum(s[0]) {
			s = s[1:]
		}
		// search for the next valid lat element
		for len(s) > 1 && !isAlphaNum(s[len(s)-1]) {
			s = s[:len(s)-1]
		}

		if len(s) <= 1 {
			return true
		}

		// if the first and last elt are not the same, it is not a palindrome
		if s[0] != s[len(s)-1] {
			return false
		}

		if len(s) == 2 && s[0] == s[len(s)-1] {
			return true
		}

		s = s[1 : len(s)-1]
	}
}

func isPalindrome2(s string) bool {
	clearedStr := strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, s)

	for i := 0; i < len(clearedStr)/2; i++ {
		if clearedStr[i] != clearedStr[len(clearedStr)-1-i] {
			return false
		}
	}

	return true
}

func isPalindrome3(s string) bool {
	clearedStr := strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, s)

	r := []rune(clearedStr)
	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}

	return string(r) == clearedStr
}

func isPalindrome4(s string) bool {
	start, end := 0, len(s)-1
	s = strings.ToLower(s)
	for start < end {
		for start < end && !isAlphaNumeric(s[start]) {
			start++
		}

		for start < end && !isAlphaNumeric(s[end]) {
			end--
		}

		if s[start] != s[end] {
			return false
		}

		start++
		end--
	}

	return true
}

func isAlphaNumeric(c uint8) bool {
	return ('a' <= c && c <= 'z') ||
		('A' <= c && c <= 'Z') ||
		('0' <= c && c <= '9')
}

func isPalindrome5(s string) bool {
	clearedStr := strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, s)

	r := []rune(clearedStr)
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			return false
		}
	}

	return true
}
