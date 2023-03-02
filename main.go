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

	x := rVar([]int{2, 4, 6, 7, 10}...)
	fmt.Println(x)
}

func rVar(y ...int) int {
	n := 0
	for _, v := range y {
		n += v
	}
	return n
}
