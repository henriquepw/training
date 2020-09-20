package main

import (
  "fmt"
  "sort"
)

func binarySearch(haystack []int, needle int) int {
  low := 0
  high := len(haystack) - 1

  for low <= high {
    median := (low + high) / 2

    if haystack[median] < needle {
      low = median + 1
    } else {
      high = median - 1
    }
  }

  if low == len(haystack) || haystack[low] != needle {
    return -1
  }

  return low
}

// Q1025 https://www.urionlinejudge.com.br/judge/pt/problems/view/1025
func Q1025() {
  var N, Q, i, count, search int

  for {
    fmt.Scanf("%d %d", &N, &Q)

    if N == 0 && Q == 0 {
      break
    }

    count++
    fmt.Printf("CASE# %d:\n", count)

    marbles := make([]int, N)
    for i = 0; i < N; i++ {
      fmt.Scanf("%d", &marbles[i])
    }

    sort.Ints(marbles)

    for i = 0; i < Q; i++ {
      fmt.Scanf("%d", &search)

      index := binarySearch(marbles, search)

      if index == -1 {
        fmt.Printf("%d not found\n", search)
      } else {
        fmt.Printf("%d found at %d\n", search, index + 1)
      }
    }
  }
}
