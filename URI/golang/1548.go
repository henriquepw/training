package main

import (
  "fmt"
  "sort"
)

// Q1548 https://www.urionlinejudge.com.br/judge/pt/problems/view/1548
func Q1548() {
  var M, N int

  fmt.Scanf("%d", &N)
  for i := 0; i < N; i++ {
    fmt.Scanf("%d", &M)

    students := make([]int, M)
    ordainedStudents := make([]int, M)
    for j := 0; j < M; j++ {
      fmt.Scanf("%d", &students[j])
    }

    copy(ordainedStudents, students)
    sort.Ints(ordainedStudents)

    var result int
    for j := 0; j < M; j++ {
      if students[j] == ordainedStudents[M - j - 1]  {
        result++
      } 
    }

    fmt.Printf("%d\n", result)
  }
}