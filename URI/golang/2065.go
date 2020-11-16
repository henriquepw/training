package main

import (
  "fmt"
  "sort"
)
// Q2065 Incomplete
// https://www.urionlinejudge.com.br/judge/pt/problems/view/2065
func Q2065() {
  var N, M int
  
  fmt.Scanf("%d %d", &N, &M)

  employers := make([]int, N)
  for i := 0; i < N; i++ {
    fmt.Scanf("%d", &employers[i])
  }

  sort.Ints(employers)

  clients := make([]int, M)
  for i := 0; i < M; i++ {
    fmt.Scanf("%d", &clients[i])
  }

  
}