func isPalindrome(s string) bool {
	var reverseString strings.Builder
	// noSpaceString := strings.ToLower(strings.ReplaceAll(s, " ", ""))
	var finalString strings.Builder

	for i := range s {
		if (s[i] >= 65 && s[i] <= 90) || (s[i] >= 97 && s[i] <= 122) || (s[i] >= 48 && s[i] <= 57) {
			finalString.WriteString(strings.ToLower(string(s[i])))
		} else {
			continue
		}
	}

	for i := len(finalString.String()) - 1; i >= 0; i-- {
		reverseString.WriteString(string(finalString.String()[i]))
	}
	return reverseString.String() == finalString.String()
}
