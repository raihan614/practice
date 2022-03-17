package main

import "fmt"

func main() {

	var myarr = []int{23171, 21011, 21123, 21366, 21013, 21367}

	result := solution(myarr)
	fmt.Println(result)
}

func solution(A []int) int {

	var result, pointer int

	for i := 0; i < len(A); i++ {

		for j := i; j < len(A); j++ {
			pointer = A[j] - A[i]
			if pointer > result {
				result = pointer
			}
		}

	}

	return result
}
