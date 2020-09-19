package main

import "fmt"

// Q2452 https://www.urionlinejudge.com.br/judge/pt/problems/view/2452
func Q2452() {
  var rowSize, drops, result, distancy int

  fmt.Scanf("%d %d", &rowSize, &drops)
  fmt.Scanf("%d", &result)
  result--

  positions := make([]int, drops - 1)
  for i := 0; i <= drops; i++ {
    fmt.Scanf("%d", &positions[i])

    distancy = (positions[i] - positions[i - 1]) / 2
    
    if distancy > result {
      result = distancy
    }
  }

  lastDistancy := rowSize - positions[drops - 2]
  if (lastDistancy >  result) {
    result = lastDistancy
  }

  fmt.Printf("%d\n", result)
}