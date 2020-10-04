package main

import (
  "io"
  "fmt"
)

// Q2506 https://www.urionlinejudge.com.br/judge/pt/problems/view/2506
func Q2506() {
  var n, h, m, c, incomingTime, limitTime int

  for {
    _, err := fmt.Scanf("%d", &n)
    if err == io.EOF {
      break
    }

    var currentTime, result int
    for ; n > 0; n-- {
      fmt.Scanf("%d %d %d", &h, &m, &c)

      incomingTime = h * 60 + m
      limitTime = incomingTime + c

      if currentTime == 0 || incomingTime > currentTime {
        currentTime = incomingTime / 30

        if incomingTime % 30 > 0 {
          currentTime++
        } 

        currentTime *= 30
      } else {
        currentTime += 30
      }

      if currentTime > limitTime {
        result++
      }
    }

    fmt.Printf("%d\n", result)
  }
}