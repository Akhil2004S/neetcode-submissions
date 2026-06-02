import "slices"

func isAnagram(s string, t string) bool {
    sortedS := sortString(s)
    sortedT := sortString(t)
    if sortedS == sortedT {
        return true
    }
    return false
}

func sortString(s string) string {
	r := []rune(s)
	slices.Sort(r)
	return string(r)
}

