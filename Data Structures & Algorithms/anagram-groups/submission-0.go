import "slices"

func groupAnagrams(strs []string) [][]string {
	words := make(map[string][]string)
	result := [][]string{}
	for _, str := range strs {
		sortedString := sortString(str)
		words[sortedString] = append(words[sortedString], str)
	}
	for _, val := range words {
		result = append(result, val)
	}
	return result
}

func sortString(s string) string {
	r := []rune(s)
	slices.Sort(r)
	return string(r)
}
