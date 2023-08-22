package main

import "fmt"

func insertionSort(entrada []int) []int {
	saida := make([]int, len(entrada))
	copy(saida, entrada)

	for j := 1; j < len(entrada); j++ {
		carta := entrada[j]
		i := j - 1
		for i >= 0 && saida[i] > carta {
			saida[i+1] = saida[i]
			i--
		}
		saida[i+1] = carta
	}

	return saida
}

func main() {
	seq_entrada := []int{10, 25, 8, 42, 16, 70, 32, 5, 18, 99, 54, 22, 76, 12, 88, 47, 3, 60, 30, 67, 91, 6, 38, 80, 15, 50}
	seq_saida := insertionSort(seq_entrada)
	fmt.Println(seq_entrada)
	fmt.Println(seq_saida)
}
