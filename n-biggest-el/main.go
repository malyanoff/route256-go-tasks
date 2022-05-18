package main

import (
	"fmt"
	"route256tasks/solution"
)

func intScanln(n int) ([]int, error) {
	x := make([]int, n)
	y := make([]interface{}, len(x))
	for i := range x {
		y[i] = &x[i]
	}
	n, err := fmt.Scanln(y...)
	x = x[:n]
	return x, err
}

func main() {
	fmt.Println("Enter M & N numbers:")

	inputData1, err := intScanln(2)
	if err != nil {
		fmt.Println(err)
		return
	}
    fmt.Println("Enter array of numbers with length equal to M")

	inputData2, err := intScanln(inputData1[0])
	fmt.Println(solution.FindBiggestEl(inputData1,inputData2))

}
