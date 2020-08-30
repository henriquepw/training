package main

import "fmt"

// Q1089 https://www.urionlinejudge.com.br/judge/pt/problems/view/1089
func Q1089() {
	var N int

	for {
		fmt.Scanf("%d", &N)

		if N == 0 {
			break
		}

		var points []int
		var peaks int

		for i := 0; i < N; i++ {
			var point int
			fmt.Scanf("%d", &point)
			points = append(points, point)
		}

		points = append(points, points[0:2]...)
		
		for i := 1; i <= N; i++ {
			if ((points[i] > points[i - 1]) && (points[i] > points[i + 1])) || ((points[i] < points[i - 1]) && (points[i] < points[i + 1])) {
				peaks = peaks + 1
			}
		}
		
		fmt.Printf("%d\n", peaks)
	}
}