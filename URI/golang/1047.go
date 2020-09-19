package main

import "fmt"

// Q1047 https://www.urionlinejudge.com.br/judge/pt/problems/view/1047
func Q1047() {
  var hourInit,  minInit, hourFinal, minFinal int16

  fmt.Scanf("%d %d %d %d", &hourInit, &minInit, &hourFinal, &minFinal)

  total := (hourFinal * 60 + minFinal) - (hourInit * 60 + minInit)

  if total <= 0 {
    total = 1440 + total // 1440 = 24 hours
	}

  fmt.Printf("O JOGO DUROU %d HORA(S) E %d MINUTO(S)\n", total / 60, total % 60)
}