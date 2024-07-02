func findAndReplacePattern(words []string, pattern string) []string {
	patternKey := ""
	letterMap := make(map[rune]int)
	i := 0
	for _, s := range pattern {
		_, ok := letterMap[s]
		if !ok {
			letterMap[s] = i
			i++
		}
		patternKey += fmt.Sprintf("%d-", letterMap[s])

	}

	result := []string{}
	for _, w := range words {
		clear(letterMap)
		i = 0
		wordPattern := ""
		for _, s := range w {
			_, ok := letterMap[s]
			if !ok {
				letterMap[s] = i
				i++
			}
			wordPattern += fmt.Sprintf("%d-", letterMap[s])
		}
		if wordPattern == patternKey {
			result = append(result, w)
		}
	}
	return result
}