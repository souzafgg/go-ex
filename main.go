package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	i := `just 
			a 
				mod`
	fmt.Println(strings.TrimSpace(i))

	j, k := "dez", 10
	fmt.Println(j, k)
	fmt.Println(j + strings.Repeat("!", k))

	name := "çatz"
	fmt.Println(utf8.RuneCountInString(name))
}
