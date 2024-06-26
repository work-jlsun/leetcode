package loader

func mergeAlternately(word1 string, word2 string) string {
	n := len(word1)
	m := len(word2)
	merge := make([]byte, 0)
	for i := 0; i < n || i < m; i++ {
		if i < n {
			merge = append(merge, word1[i])
		}
		if i < m {
			merge = append(merge, word2[i])

		}
	}
	return string(merge)
}
