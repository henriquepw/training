package main

import (
  "io"
  "fmt"
)

// Q1068 https://www.urionlinejudge.com.br/judge/pt/problems/view/1068
func Q1068() {
  var expresion string
  
  for {
    var openBrackets int

    _, err := fmt.Scanf("%s", &expresion)
    if err == io.EOF {
      break
    }

    n := len(expresion)
    for i := 0; i < n; i++ {
      if expresion[i] == '(' {
        openBrackets++
      } else if expresion[i] == ')'{
        openBrackets--

        if openBrackets < 0 {
          break
        }
      }
    }

    if openBrackets == 0 {
      fmt.Print("correct\n")
    } else {
      fmt.Print("incorrect\n")
    }
  } 
}