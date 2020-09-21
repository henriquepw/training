package main

import (
  "fmt"
  "container/list"
)

// Q1110 https://www.urionlinejudge.com.br/judge/pt/problems/view/1110
func main() {
	var n int
	
  for {
    fmt.Scanf("%d", &n)

    if n == 0 {
      break
    }

    cards := list.New()
    for i := 1; i <= n; i++ {
      cards.PushBack(i)
    }

    fmt.Print("Discarded cards: ")
    for cards.Len() > 1 {
      card := cards.Front()

      if cards.Len() < n {
        fmt.Printf(", %d", card.Value)
      } else {
        fmt.Printf("%d", card.Value)
      }

      cards.Remove(card)
      cards.MoveToBack(cards.Front())
    }

    fmt.Printf("\nRemaining card: %d\n", cards.Front().Value)
  }
}