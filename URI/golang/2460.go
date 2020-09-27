package main

import "fmt"

// Q2460 https://www.urionlinejudge.com.br/judge/pt/problems/view/2460
func Q2460() {
  var N, M, read int
  fmt.Scanf("%d", &N)

  queue := make([]int, N)
  set := make(map[int]bool)
  for i := 0; i < N; i++ {
    fmt.Scanf("%d", &queue[i])
    set[queue[i]] = true
  }

  fmt.Scanf("%d", &M)
  for i := 0; i < M; i++ {
    fmt.Scanf("%d", &read)
    set[read] = false
  }

  var result string
  for i := 0; i < N; i++ {
    if set[queue[i]] {
      result += fmt.Sprintf("%d ", queue[i])
    }
  }

  strSize := len(result) - 1
  if result[strSize] == ' ' {
    result = result[:strSize]
  }

  fmt.Printf("%s\n", result)
}