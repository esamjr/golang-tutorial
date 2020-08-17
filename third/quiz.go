package main

import "fmt"

func main() {
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	// QUIZZ 1
	var total int

	for _, score := range scores {
		total = total + score

	}
	length := len(scores)
	average := float64(total) / float64(length)

	fmt.Println(average)

	// quizz 2
	var goodScores []int

	for _, gScore := range scores {
		if gScore >= 78 {
			goodScores = append(goodScores, gScore)
		}
	}
	fmt.Println(goodScores)
}
