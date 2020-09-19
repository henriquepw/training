package main

import "fmt"

// Q1281 https://www.urionlinejudge.com.br/judge/pt/problems/view/1281
func Q1281() {
  var N, M, P int
  var productName string
  var productNum float32

  fmt.Scanf("%d", &N)
  for i := 0; i< N; i++ {
    var totalPrice float32
    productsForSale := make(map[string]float32)
    
    fmt.Scanf("%d", &M)
    for j := 0; j < M; j++ {
      fmt.Scanf("%s %f", &productName, &productNum)
      productsForSale[productName] = productNum
    }

    fmt.Scanf("%d", &P)
    for j := 0; j < P; j++ {
      fmt.Scanf("%s %f", &productName, &productNum)
      totalPrice += productsForSale[productName] * productNum
    }

    fmt.Printf("R$ %.2f\n", totalPrice)
  }
}