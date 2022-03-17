package main

import "fmt"

func main() {

	var myarr []int
	myarr = append(myarr, 23171)
	myarr = append(myarr, 21011)
	myarr = append(myarr, 21123)
	myarr = append(myarr, 21366)
	myarr = append(myarr, 21013)
	myarr = append(myarr, 21367)

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
