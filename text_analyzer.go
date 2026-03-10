package main

import (
	"fmt"
	"strings"
)

func getTopWords(wordMap map[string]int, n int) []string {
	res := make([]string, 0)

	for i := 0; i < n; i++ {
		maxWord := ""
		maxCount := 0

		for word, count := range wordMap {
			if count > maxCount {
				maxCount = count
				maxWord = word
			}
		}

		res = append(res, maxWord)
		delete(wordMap, maxWord)
	}

	return res
}

func AnalyzeText(text string) {
	text = strings.ToLower(text)
	r := strings.NewReplacer(
		".", " ",
		",", " ",
		"!", " ",
		"?", " ",
	)
	text = r.Replace(text)
	words := strings.Fields(text)

	m := make(map[string]int)

	amountWords := len(words)

	for _, word := range words {
		m[word]++
	}

	unique := len(m)

	max := 0
	maxWord := ""
	for word, count := range m {
		if count > max {
			max = count
			maxWord = word
		}
	}

	copyMap := make(map[string]int)
	for k, v := range m {
		copyMap[k] = v
	}

	fmt.Printf("Количество слов: %d\n", amountWords)
	fmt.Printf("Количество уникальных слов: %d\n", unique)
	fmt.Printf("Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n", maxWord, m[maxWord])
	fmt.Println("Топ-5 самых часто встречающихся слов:")
	for _, word := range getTopWords(copyMap, 5) {
		fmt.Printf("\"%s\": %d раз\n", word, m[word])
	}
}
