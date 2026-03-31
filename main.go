package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(capnth([]string{"how", "are", "you", "today"}, 1))
	fmt.Println(ispunc([]string{"hello", ",", "world", "!"}))
	fmt.Println(isbool("="))
	fmt.Println(aorAn("ball"))
	fmt.Println(fixArticles("There it was. A amazing rock. A honest man. A book."))
	result := fixsingle("'hello world'")
	fmt.Printf("%q\n", result)
}
func capnth(word []string, n int) []string {
	for i := len(word) - n; i < len(word); i++ {
		word[i] = strings.ToUpper(word[i])
	}
	return word
}

func ispunc(word []string) string {
	result := strings.Join(word, " ")
	result = strings.ReplaceAll(result, " ,", ",")
	result = strings.ReplaceAll(result, " !", "!")
	return result
}
func isbool(s string) bool {
	return strings.ContainsAny(s, "!,.:")
}

func aorAn(word string) string {
	if word == "" {
		return "a"
	}
	if strings.ContainsAny(strings.ToLower(string(word[0])), "aeiouh") {
		return "an"
	}
	return "a"
}

func fixArticles(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words)-1; i++ {
		firstletter := words[i+1][:1]

		if words[i] == "An" && !strings.ContainsAny(firstletter, "aeiouAEIOUHh") {
			words[i] = "A"
		}

		if words[i] == "A" && strings.ContainsAny(firstletter, "aeiouAEIOUHh") {
			words[i] = "An"

		}

	}
	return strings.Join(words, " ")
}

func fixsingle(text string) string {
	return "'" + strings.TrimSpace(strings.Trim(text, "'")) + "'"
}
