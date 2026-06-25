import "log"

type Solution struct{
	encodedString string
	decodedString []string
}

func (s *Solution) Encode(strs []string) string {
	var str strings.Builder

	for _, word := range strs {
		delimiter := fmt.Sprintf("%d#", len(word))
		str.WriteString(delimiter)
		str.WriteString(word)
	}
	s.encodedString = str.String()
	return str.String()
}

func (s *Solution) Decode(encoded string) []string {
	var output []string
	var word strings.Builder

	for {
		index := strings.Index(encoded, "#")
		if index == -1 {
			break
		}
		before, after, _ := strings.Cut(encoded, "#")
		numberOfWords, err := strconv.Atoi(before)
		if err != nil {
			log.Fatal(err)
		}
		for i := range numberOfWords {
			word.WriteString(string(after[i]))
		}
		encoded = encoded[index+numberOfWords+1:]
		output = append(output, word.String())
		word.Reset()

	}
	return output
}
