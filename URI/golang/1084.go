package main

import "fmt"

// Q1084 https://www.urionlinejudge.com.br/judge/pt/problems/view/1084
func Q1084() {
  var n, d int

  for {
    fmt.Scanf("%d %d", &n, &d)

    if n == d && d == 0 {
      break
    }

    var str string
    fmt.Scanf("%s", &str)

    size := len(str) + 1
    list := make([]int, size)
    listAux := make([]int, size)

    list[size - 1], listAux[size - 1] = -1, -1
    for i := 0; i < len(str); i++ {
      list[i] = int(str[i])
      listAux[i] = int(str[i])
    }

    j := 1
    count := 0
    for i := 0; i < len(str) && count != d; i++ {
      for i >= 0 && list[j] > list[i] && count != d {
        if listAux[i] != -1 {
          listAux[i] = -1
          count++
        }

        i--
      }
      
      i = j - 1
      j++
    }

    count = 0
    for i := 0; i < len(str) && count != n - d; i++ {
      if listAux[i] != -1 {
        fmt.Printf("%c", listAux[i])
        count++
      }
    }
    fmt.Print("\n")
  }
}