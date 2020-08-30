package main

import "fmt"

// Q1019 https://www.urionlinejudge.com.br/judge/pt/problems/view/1019
func Q1019() {
	var N int

	fmt.Scanf("%d", &N)

	fmt.Printf("%d:%d:%d\n", N / 3600, (N % 3600) / 60, N % 60)
}