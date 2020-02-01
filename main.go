package main

import (
	"algorithm_questions/algos"
	"fmt"
)

func main() {
	a := []int{3, 5, 6, 3, 3, 5, 6, 7, 8, 7, 7, 6, 3, 3, 4, 5, 6, 7, 8, 5, 4, 3, 4, 8}
	steps := algos.Parity("0011100")
	count := algos.IdenticalIndices(a)
	fmt.Println(steps)
	fmt.Println(count)
}
