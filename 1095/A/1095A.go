package main

import (
	"fmt"
)

func decode(in string) string {
	var i int

	for i < len(in) {
		in = in[:i+1] + in[i+i+1:]
		i++
	}

	return in
}

func main() {
	var lenin int
	var text string
	fmt.Scanf("%d\n", &lenin)
	fmt.Scanf("%s\n", &text)

	fmt.Print(decode(text))
}
