package main

import "fmt"

// Q1018 https://www.urionlinejudge.com.br/judge/pt/problems/view/1018
func Q1018() {
  var N, amount int
  notes := [7]int{100, 50, 20, 10, 5, 2, 1}

  fmt.Scanf("%d", &N)
  fmt.Printf("%d\n", N)

  for i := 0; i < 7; i++ {
    amount = N / notes[i]
    fmt.Printf("%d nota(s) de R$ %d,00\n", amount, notes[i])
    N = N % notes[i]
  }
}