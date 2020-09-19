package main

import (
  "fmt"
  "sort"
)

// Q1171 https://www.urionlinejudge.com.br/judge/pt/problems/view/1171
func Q1171() {
  var N int

  fmt.Scanf("%d", &N)

  values := make(map[int]int)
  var keys []int

  for i := 0; i < N; i++ {
    var value int
    fmt.Scanf("%d", &value)

    _, key := values[value]

    if !key {
      keys = append(keys, value)
    }
    
    values[value] = values[value] + 1
  }

  sort.Ints(keys)

  for _, key := range keys {
    fmt.Printf("%d aparece %d vez(es)\n", key, values[key])
  }
}