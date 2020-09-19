package main

import "fmt"

// Time limit exceeded, this solution works on c++
// Q1566 https://www.urionlinejudge.com.br/judge/pt/problems/view/1566
func Q1566() {
  var N, NC, i, j, h, value int

  fmt.Scanf("%d", &N)

  for i = 0; i < N; i++ {
    fmt.Scanf("%d", &NC)

    heights := make([]int, 231)
    for j = 0; j < NC; j++ {
      fmt.Scanf("%d", &value)
      heights[value]++
    }

    var notFirst bool
    for j = 20; j < 231; j++ {
      if heights[j] > 0 {
        for h = 0; h < heights[j]; h++ {
          if notFirst {
            fmt.Print(" ")
          }
          notFirst = true
        
          fmt.Printf("%d", j)
        }
      }
    }
    
    fmt.Print("\n")
  }
}
