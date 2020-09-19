package main

import "fmt"

// Q2415 https://www.urionlinejudge.com.br/judge/pt/problems/view/2415
func Q2415() {
  var N, points, last, currentPoints, current int

  fmt.Scanf("%d", &N)

  for i := 0; i < N; i++ {
    fmt.Scanf("%d", &current)
    
    if (last == current) || (i == 0) {
      currentPoints = currentPoints + 1
    } else {
      if currentPoints > points {
        points = currentPoints
      }

      currentPoints = 1
    }

    last = current
  }

  if currentPoints > points {
    points = currentPoints
  }
  
  fmt.Printf("%d\n", points)
}