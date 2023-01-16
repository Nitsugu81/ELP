package main

import (
	"sync"
)

func main() {
	m1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	m2 := [][]int{{9, 8, 7}, {6, 5, 4}, {3, 2, 1}}
	m3 := mat_cal(m1, m2)
	for i := 0; i < len(m3); i++ {
		for j := 0; j < len(m3[i]); j++ {
			print(m3[i][j], " ")
		}
		println()

	}

}

func elem_cal(col []int, lin []int) int {
	sum := 0
	for i := 0; i < len(col); i++ {
		sum += col[i] * lin[i]
	}
	return sum
}

func mat_cal(mat1 [][]int, mat2 [][]int) [][]int {
	m3 := make([][]int, len(mat1))
	for i := range m3 {
		m3[i] = make([]int, len(mat2[0]))
	}
	var wg sync.WaitGroup
	for i := 0; i < len(m3); i++ {
		for j := 0; j < len(m3[i]); j++ {
			wg.Add(1)
			go func(i, j int) {
				row := mat1[i]
				var column []int
				for l := 0; l < len(mat2); l++ {
					column = append(column, mat2[l][j])
				}
				m3[i][j] = elem_cal(column, row)
				wg.Done()
			}(i, j)
		}
	}
	wg.Wait()
	return m3
}
