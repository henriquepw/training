package main

import "fmt"

// Q1895 https://www.urionlinejudge.com.br/judge/pt/problems/view/1895
func Q1895() {
  var N, cardOnTable, cardPulled, limit, i, alicePoints, BobPoints int8
  var aliceTurn bool

  fmt.Scanf("%d %d %d", &N, &cardOnTable, &limit)

  for i = 1; i < N; i++ {
    aliceTurn = !aliceTurn
    fmt.Scanf("%d", &cardPulled)

    currentPoints := cardOnTable - cardPulled
    
    if currentPoints < 0 {
      currentPoints = -currentPoints
    }
    
    if currentPoints > limit {
      continue
    }

    if aliceTurn {
      alicePoints = alicePoints + currentPoints
    } else {
      BobPoints = BobPoints + currentPoints
    }

    cardOnTable = cardPulled
  }

  fmt.Printf("%d %d\n", alicePoints, BobPoints)
}