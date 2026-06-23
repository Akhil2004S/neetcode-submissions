func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		for left < right && !(unicode.IsLetter(rune(s[left])) || unicode.IsDigit(rune(s[left]))) {
			left++
		}
		for left < right && !(unicode.IsLetter(rune(s[right])) || unicode.IsDigit(rune(s[right]))) {
			right--
		}

		if strings.EqualFold(string(s[left]), string(s[right])) {
			left++
			right--
			continue
		} else {
			return false
		}
	}

	return true
}