package ransomnote

// Given two strings ransomNote and magazine,
// return true if ransomNote can be constructed by using the letters from magazine
// and false otherwise.
//
// Each letter in magazine can only be used once in ransomNote.
//
// Time complexity: O(n)
// Space complexity: O(1)
func canConstruct(ransomNote string, magazine string) bool {
	letters := make(map[rune]int)
	for _, r := range ransomNote {
		letters[r] += 1
	}

	availableLetters := make(map[rune]int)
	for _, r := range magazine {
		availableLetters[r] += 1
	}

	for r, n := range letters {
		if availableLetters[r] < n {
			return false
		}
	}

	return true
}

// Time complexity: O(n)
// Space complexity: O(1)
func canConstruct_array(ransomNote string, magazine string) bool {
	var letterCount [26]int

	for _, r := range magazine {
		letterCount[r-'a'] += 1
	}

	for _, r := range ransomNote {
		letterCount[r-'a'] -= 1
		if letterCount[r-'a'] < 0 {
			return false
		}
	}

	return true
}
