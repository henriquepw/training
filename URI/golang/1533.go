package main

import "fmt"

// Q1533 https://www.urionlinejudge.com.br/judge/pt/problems/view/1533
func Q1533() {
	var N, biggest, secondyBiggest, biggestIndex, guilty, i uint16

	for {
		fmt.Scanf("%d", &N)
		suspects := make([]uint16, N)
	
		if N == 0 {
			break
		}

		for i = 0; i < N; i++ {
			fmt.Scanf("%d", &suspects[i])
		}

		if suspects[0] > suspects[1] {
			biggest, secondyBiggest, biggestIndex, guilty = suspects[0], suspects[1], 0, 1
		} else {
			biggest, secondyBiggest, biggestIndex, guilty = suspects[1], suspects[0], 1, 0
		}

		for i = 2; i < N; i++ {
			if suspects[i] > biggest {
				if biggest > secondyBiggest {
					secondyBiggest = biggest
					guilty = biggestIndex;
				}

				biggest = suspects[i]
				biggestIndex = i
			}

			if suspects[i] > secondyBiggest && biggest > suspects[i] {
				secondyBiggest = suspects[i]
				guilty = i;
			}
		}

		fmt.Printf("%d\n", guilty + 1)
	}
}