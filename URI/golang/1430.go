package main

import "fmt"

// Q1430 https://www.urionlinejudge.com.br/judge/pt/problems/view/1430
func Q1430() {
  notes := map[string]float32 {
    "W": 1,
    "H": 0.5,
    "Q": 0.25,
    "E": 0.125,
    "S": 0.0625,
    "T": 0.03125,
    "X": 0.015625,
  }
  
  var composition string
  for {
    fmt.Scanf("%s", &composition)

    if composition == "*" {
      break
    }

    strSize := len(composition)
    var sum float32
    var right int
    for i := 1; i < strSize; i++ {
      letter := string(composition[i])

      if (letter == "/") {
        if (sum == 1) {
          right++
        }

        sum = 0
        continue
      }

      sum += notes[letter]
    }

    fmt.Printf("%d\n", right)
  }
}