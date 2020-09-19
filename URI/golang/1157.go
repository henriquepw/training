package main

import "fmt"

// Q1157 https://www.urionlinejudge.com.br/judge/pt/problems/view/1157
func Q1157() {
  var N int

  fmt.Scanf("%d", &N)

  for i := 1; i <= N / 2; i++ {
    if N%i == 0 {
      fmt.Printf("%d\n", i)
    }
  }

  fmt.Printf("%d\n", N)
}
