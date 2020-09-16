package main

import "fmt"


// Dada uma lista de N valores inteiros, faça um algoritmo que determine se é possível encontrar 2 elementos cuja soma seja X.

// 􏰓 Ex: lista = [1,6,7,5,2]. N=5 e X=8 → Resposta: SIM
// 􏰓 Ex2: lista = [1,6,7,5,2]. N = 5 e X = 10 → Resposta: NAO
func main() {
	var N, X, i, aux, diff int
	result := "NAO"

	fmt.Scanf("%d %d", &N, &X)

	numbers := make(map[int]int, N)
	for i = 0; i < N; i++ {
		fmt.Scanf("%d", &aux)
		numbers[aux]++

		diff = X - aux

		if (diff == aux && numbers[aux] > 1) || (diff != aux && numbers[diff] > 0) {
			result = "SIM"
		}
	}

	fmt.Printf("%s \n", result)
}