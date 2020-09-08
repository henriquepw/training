package main

import (
	"io"
	"fmt"
)

// Q1222 https://www.urionlinejudge.com.br/judge/pt/problems/view/1222
func Q1222() {
	var n, l, c, chars, lines, pages int
	var word string
	
	for {
		_, err := fmt.Scanf("%d %d %d", &n, &l, &c)
		
		if err == io.EOF {
			break
		}
		
		lines, chars = 0, 0;

		for i := 0; i < n; i++ {
			fmt.Scanf("%s", &word)
			chars += len(word)

			if chars < c {
				chars++
			}

			if chars == c {
				chars = 0
				lines++
			}

			if chars > c {	
				chars = len(word) + 1
				lines++
			}
		}

		if (chars > 0) {
			lines++
		}

		pages = lines / l
		if (lines % l > 0) {
			pages++
		}

		fmt.Printf("%d\n", pages)
	}
}