package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type tools struct {
	Name     string
	Objetivo string
}

type use struct {
	Name      string
	Nota      int
	Essencial bool
}

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

	x = rSoma([]int{1, 3, 5, 7, 9})
}

func rVar(y ...int) int {
	n := 0
	for _, v := range y {
		n += v
	}
	return n
}

func rSoma(x []int) int {
	t := 0
	for _, v := range x {
		t += v
	}
	return t
}

func print(s string) {
	fmt.Println(s)
}
