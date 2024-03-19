package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`\s+`)
	fmt.Println(re.ReplaceAllString("AS -4124GS-TNR", ""))
}
