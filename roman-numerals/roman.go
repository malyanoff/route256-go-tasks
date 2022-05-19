/*
Римская система счисления содержит символы: I, V, X, L, C, D и M.

Следующие цифры римской системы счисления соответствуют символам в десятичной системе: I = 1; V = 5; X = 10; L = 50; C = 100; D = 500; M = 1000.

Числа в римской системе обычно записываются от большего к меньшему, слева направо, например, семь можно записать как V + II = VII, а двадцать семь как XX + V + II = XXVII.

Однако, число 4 записывается не как IIII, а как IV. Т.е. Для получения некоторых цифр в римской системе счисления используется принцип вычитания.

Следующие правила задают условия, при которых следует использовать принцип вычитания:

1. I может предшествовать V (5) и X (10) для получения 4-х и 9-и соответственно;
2. X может предшествовать L (50) и С (100) для получения 40-а и 90-а соответственно;
3. С может предшествовать D (500) и M (1000) для получения 400-а и 900-а соответственно;

Но не всё так просто. Согласно классическим римским правилам, максимальное число, которое может быть представлено в этой системе счисления, это 3999.
Однако, есть способы, которые позволяют обойти это ограничение. Мы будем использовать модификацию vinculum.
В базовой версии это позволяет разделить число на два блока цифр, первый из которых умножается на 1000, т.е. IV|DCXXVII = 4 627, при этом второй блок не больше тысячи, чтобы не было коллизий при записи.
Мы же пойдём чуть дальше и введём следующие правила:

1. Может быть сколько угодно много блоков, которые разделены символом |
2. Каждый из блоков, кроме первого, содержит в себе число меньше 1000
3. Первый блок может быть каким угодно римским числом.
4. Предполагается, что все числа до 3999 включительно мы пишем как раньше, чтобы не было коллизий.
5. Пустой разряд заменяется символом N (nihil, ничто). N используется только для записи пустого разряда. XXN не валидная запись.
Примеры записей:

1. X|N|X = 10 000 010
2. X|N = 10 000
3. DCXXVII|X|N|X = 627 010 000 010

*/

package roman

import (
	"math"
	"strings"
)

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
	"N": 0,
}

func romanToInt(romanNumeral string) int {
	if romanNumeral == "" {
		return 0
	}
	total, lastnum, num := 0, 0, 0
	for i := 0; i < len(romanNumeral); i++ {
		char := romanNumeral[len(romanNumeral)-(i+1) : len(romanNumeral)-i]
		num = roman[char]

		if num < lastnum {
			total = total - num
		} else {
			total = total + num
		}

		lastnum = num

	}

	return total
}

//Roman vinculum system realization
func convertRomanBlocks(romanNumeral string) int {

	blocks := strings.Split(romanNumeral, "|")

	//coefficient for multiplying
	q := len(blocks) - 1

	total := 0

	for _, block := range blocks {
		num := romanToInt(block)
		total += num * int(math.Pow(float64(1000), float64(q)))
		q--
	}

	return total
}

func CompareTwoRomans(roman1, roman2 string) int {

	if convertRomanBlocks(roman1) == convertRomanBlocks(roman2) {
		return 1
	}

	if convertRomanBlocks(roman1) > convertRomanBlocks(roman2) {
		return 1
	} else {
		return -1
	}

}
