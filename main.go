package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Circle struct {
	x int
	y int
}

//готова
func opr_x_i_y(text string, text1 string, text2 string) (string, string) {

	rs := []rune(text)
	p := len(text)
	o := 0
	for i := 0; i < p; i++ {
		switch {
		case (rs[i]) == ('*'):
			{
				o = i
				i = p
				break
			}
		case (rs[i]) == ('/'):
			{
				o = i
				i = p
				break
			}
		case (rs[i]) == ('+'):
			{
				o = i
				i = p
				break
			}
		case (rs[i]) == ('-'):
			{
				o = i
				i = p
				break
			}

		}

	}
	text1 = text[0:o]
	text2 = text[o+1 : p]
	return text1, text2
}

//готова
func proverka_na_rim(text1, text2 string) bool {
	rs := []rune(text1)
	rs2 := []rune(text2)
	var b, b1 bool
	for i := range rs {
		if unicode.IsLetter(rs[i]) == true {
			b = true
			break
		}
	}

	for i := range rs2 {
		if unicode.IsLetter(rs2[i]) == true {
			b1 = true
			break
		}
	}
	if ((b == true) && (b1 != true)) || ((b != true) && (b1 == true)) {
		panic("Вывод ошибки, так как используются одновременно разные системы счисления.\n")
	}

	if b == true && b1 == true {
		return true
	} else {

		return false
	}
}

//готова
func preobr_iz_arab_v_rim(number int) string {

	table_namber := [9]int{
		(1), (4), (5), (9),
		(10), (40), (50), (90),
		(100)}
	table_symvol := [9]string{
		("I"), ("IV"), ("V"), ("IX"),
		("X"), ("XL"), ("L"), ("XC"),
		("C")}
	N := number
	result := ""
	for i := 8; N > 0; i-- {
		for table_namber[i] > N {
			i--
		}
		result += table_symvol[i]
		N -= table_namber[i]
		i = i + 1
	}

	return result

}

//готова
func preobr_iz_rim_v_arab(rim string) int {
	table_namber := [9]int{
		(1), (5), (4), (10), (9),
		(50), (40), (100), (90)}
	table_symvol := [9]string{
		("I"), ("V"), ("IV"),
		("X"), ("IX"), ("L"), ("XL"),
		("C"), ("XC")}

	res := 0

	for i := len(rim) - 1; i >= 0; i-- {
		for j := 8; j >= 0; {
			if strings.HasSuffix(rim, table_symvol[j]) == false {
				j--
			} else {
				rim = strings.TrimSuffix(rim, table_symvol[j])
				res += table_namber[j]
				break
			}

		}

	}
	return res
}

//готова
func itog(text string, x, y int) int {
	var itog int
	if strings.Contains(text, "*") == true {
		itog = x * y
	}
	if strings.Contains(text, "/") == true {
		itog = x / y
	}
	if strings.Contains(text, "+") == true {
		itog = x + y
	}
	if strings.Contains(text, "-") == true {
		itog = x - y
	}
	if (strings.Contains(text, "*") == false) && (strings.Contains(text, "/") == false) && (strings.Contains(text, "+") == false) && (strings.Contains(text, "-") == false) {
		panic(" Ошибка,строка не соответствует ни одной  арифметической операции")
	}
	if (strings.Count(text, "*") > 1) || (strings.Count(text, "/") > 1) || (strings.Count(text, "-") > 1) || (strings.Count(text, "+") > 1) {
		panic("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
	return itog
}

func main() {
	//var x, y int
	var text1, text2, results string
	var text1int, text2int, resultin int
	fmt.Println("введите неравенство")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSpace(text)

	//fmt.Println(preobr_v_arab(39))
	//fmt.Println(preobr_iz_rim_v_arab("LXV"))
	//itog(text, &Circle{x, y})
	//	fmt.Println(opr_x_i_y("VI+I", text1, text2))
	//text1, text2 = opr_x_i_y("VI+I", text1, text2)
	//fmt.Println(proverka_na_rim(text1, text2))
	text1, text2 = opr_x_i_y(text, text1, text2)
	if proverka_na_rim(text1, text2) == true {
		text1int = preobr_iz_rim_v_arab(text1)
		text2int = preobr_iz_rim_v_arab(text2)
		if (text1int > 10) || (text2int > 10) {
			panic("Ошибка,Число >10")
		}
		resultin = itog(text, text1int, text2int)
		if resultin < 1 {
			fmt.Println("Ошибка,результатом работы калькулятора с римскими числами могут быть только положительные числа,ваше число меньше 1")
		} else {
			results = preobr_iz_arab_v_rim(resultin)
			fmt.Println("Ответ:", results)
		}
	} else {
		text1int, _ = strconv.Atoi(text1)
		text2int, _ = strconv.Atoi(text2)
		if (text1int > 10) || (text2int > 10) {
			panic("Ошибка,Число >10")
		}
		resultin = itog(text, text1int, text2int)
		fmt.Println("Ответ:", resultin)
	}
}
