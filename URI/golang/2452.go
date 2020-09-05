package main

import "fmt"

// Q2452 https://www.urionlinejudge.com.br/judge/pt/problems/view/2452
func Q2452() {
	var rowSize, drops, result, distancy int

	fmt.Scanf("%d %d", &rowSize, &drops)

	positions := make([]int, drops)
	for i := 0; i < drops; i++ {
		fmt.Scanf("%d", &positions[i])

		if i > 0 {
			distancy = (positions[i] - positions[i - 1]) / 2
			
			if distancy > result {
				result = distancy
			}
		} else {
			result = positions[0] - 1
		}
	}

	lastDistancy := rowSize - positions[drops - 1]
	if (lastDistancy >  result) {
		result = lastDistancy
	}

	fmt.Printf("%d\n", result)
}