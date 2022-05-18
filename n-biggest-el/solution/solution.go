/*
Есть массив целых чисел, требуется найти в нём самое большое N-ое число.

Input data format:
На вход подаётся целое число 0 < M ≤ 105, целое число 0 < N ≤ 105, разделённые пробелом на первой строке.
На второй строке массив из M целых чисел, разделённых знаком пробел.

Output data format:
Требуется вывести одно число – N-ое самое большое число или -1, если такого числа нет.

example input:
6 3
1 2 3 -9 3 5

output: 2

*/

package solution

import (
	"sort"
)

func FindBiggestEl(conditionsArr []int, inputArr []int) (biggest int) {


	nPos := conditionsArr[1]

	if nPos > conditionsArr[0]{
		return -1
	}

	sort.Slice(inputArr, func(i, j int) bool { return inputArr[i] > inputArr[j] })

	formatedArr := removeDuplicates(inputArr)

	if len(formatedArr) < nPos {
		return -1
	}

	return formatedArr[nPos-1]
}

func removeDuplicates(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
