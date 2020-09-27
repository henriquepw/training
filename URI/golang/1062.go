package main

import (
  "fmt"
  "container/list"
)

// Q1062 https://www.urionlinejudge.com.br/judge/pt/problems/view/1062
func Q1062() {
  var n, next, wagon, value int
  var train, aux, output *list.List
  
  for {
    fmt.Scanf("%d", &n)

    if n == 0 {
      break
    }

    for {
      next = n
      train = list.New()
      aux = list.New()
      output = list.New()

      fmt.Scanf("%d", &wagon)

      if wagon == 0 {
        break
      }

      train.PushBack(wagon)
      for i := 1; i < n; i++ {
        fmt.Scanf("%d", &wagon)
        train.PushBack(wagon)
      }

      for train.Len() > 0 {
        value := train.Remove(train.Back())
        aux.PushBack(value)

        for aux.Back() != nil && aux.Back().Value == next {
          output.PushBack(aux.Back().Value)
          aux.Remove(aux.Back())
          next--
        }
      }

      next = n
      result := true
      for result && output.Len() > 0 {
        value = output.Remove(output.Front()).(int)

        if value != next {
          result = false
        }

        next--
      }

      for result && aux.Len() > 0 {
        value = aux.Remove(aux.Back()).(int)
        
        if value != next {
          result = false
        }

        next--
      }

      if result {
        fmt.Print("Yes\n")
      } else {
        fmt.Print("No\n")
      }
    }

    fmt.Print("\n")
  }
}