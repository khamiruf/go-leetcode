func isAnagram(s string, t string) bool {
	// if different length, not anagram
	if len(s) != len(t) {
		return false
	}

	// init bucket of 26
	bucket := make([]int, 26)
	for index, _ := range bucket {
		bucket[index] = 0
	}

	// increment and decrement the bucket counter
	for i := 0; i < len(s); i++ {
		bucket[int(s[i])-int('a')]++
		bucket[int(t[i])-int('a')]--
	}

	for _, val := range bucket {
		if val != 0 {
			return false
		}
	}

	return true
}