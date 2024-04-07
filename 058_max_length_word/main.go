package main

import (
	"fmt"
	"strings"
)

func main() {
	data := []string{
		"Hello World",
		"   fly me   to   the moon  ",
		"luffy is still joyboy",
	}

	for _, str := range data {
		fmt.Println(lengthOfLastWordByRear(str))
	}
}

func lengthOfLastWord(s string) int {
	parts := strings.Split(strings.TrimSpace(s), " ")
	if l := len(parts); l > 0 {
		return len(parts[l-1])
	}
	return 0
}

func lengthOfLastWordByRear(s string) int {
	isword := false
	wordLength := 0
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " && isword {
			return wordLength
		} else if string(s[i]) != " " {
			isword = true
			wordLength++
			if i == 0 {
				return wordLength
			}
		}
	}
	return 0
}
