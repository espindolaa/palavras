package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	file, _ := os.Open("br-utf8.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	occurrences := map[string]int{}

	for _, word := range text {
		occurrences[word] = 0
	}

	for _, word := range text {
		size := len(word) - 1
		for i := size; i > 1; i-- {
			for j := 0; j+i <= size; j++ {
				trim := word[j : i+j]

				if val, ok := occurrences[trim]; ok {
					occurrences[trim] = val + 1
				}
			}
		}

	}
	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range occurrences {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for _, o := range ss[:10] {
		fmt.Println(o)
	}
}
