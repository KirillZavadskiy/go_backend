package main

import "fmt"

func MostPopularWord(words []string) string {
	wordsCount := make(map[string]int, 0)
	mostPopWord := ""
	max := 0

	for i := len(words) - 1; i >= 0; i -= 1 {
		word := words[i]
		wordsCount[word] += 1
		if wordsCount[word] >= max {
			max = wordsCount[word]
			mostPopWord = word
		}
	}
	return mostPopWord
}

func main() {
	words := []string{"s", "d", "d"}
	fmt.Println(MostPopularWord(words))
}
