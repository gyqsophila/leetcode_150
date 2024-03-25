package main

import "fmt"

var (
	m = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
)

func romanToInt(s string) int {
	num := m[string(s[0])]
	lastNum := num
	for i := 1; i < len(s); i++ {
		n := m[string(s[i])]
		num += n
		if n > lastNum {
			num -= lastNum * 2
		}
		lastNum = n
	}
	return num
}

func main() {
	t := []string{
		"III",
		"LVIII",
		"MCMXCIV",
	}
	for i := range t {

		fmt.Println(romanToInt(t[i]))
	}
}
