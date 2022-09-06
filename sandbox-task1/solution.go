package main

// В магазине акция: «купи три одинаковых товара и заплати только за два». Конечно, каждый купленный товар может участвовать лишь в одной акции. Акцию можно использовать многократно.

// Например, если будут куплены 7 товаров одного вида по цене 2 за штуку и 5 товаров другого вида по цене 3 за штуку, то вместо 7⋅2+5⋅3 надо будет оплатить 5⋅2+4⋅3=22.

// Считая, что одинаковые цены имеют только одинаковые товары, найдите сумму к оплате.

// Неполные решения этой задачи (например, недостаточно эффективные) могут быть оценены частичным баллом.

// Входные данные
// В первой строке записано целое число t (1≤t≤104) — количество наборов входных данных.

// Далее записаны наборы входных данных. Каждый начинается строкой, которая содержит n (1≤n≤2⋅105) — количество купленных товаров. Следующая строка содержит их цены p1,p2,…,pn (1≤pi≤104). Если цены двух товаров одинаковые, то надо считать, что это один и тот товар.

// Гарантируется, что сумма значений n по всем тестам не превосходит 2⋅105.

// Выходные данные
// Выведите t целых чисел — суммы к оплате для каждого из наборов входных данных.


import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		goods := make([]int, 0)
		var goodsQuantity int
		fmt.Fscan(in, &goodsQuantity)

		for i := 0; i < goodsQuantity; i++ {
			var goodPrice int
			fmt.Fscan(in, &goodPrice)
			goods = append(goods, goodPrice)
		}
		fmt.Fprintln(out, countTotalSum(goods))

	}

}

func countTotalSum(inp []int) int {
	var decreaseSum, totalSum int

	duplicateMap := dupCount(inp)

	for k, v := range duplicateMap {
		decreaseSum += int(math.Floor(float64(v/3))) * k
	}

	for _, price := range inp {
		totalSum += price
	}

	return totalSum-decreaseSum
}

func dupCount(list []int) map[int]int {

	duplicateFrequency := make(map[int]int)

	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map

		_, exist := duplicateFrequency[item]

		if exist {
			duplicateFrequency[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicateFrequency[item] = 1 // else start counting from 1
		}
	}
	return duplicateFrequency
}
