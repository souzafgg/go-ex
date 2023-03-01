package main

import (
	"fmt"
	"strings"
)

func main() {

	i := `just 
			a 
				mod`
	fmt.Println(i)

	j, k := "dez", 10

	fmt.Println(j + strings.Repeat("!", k))

}
