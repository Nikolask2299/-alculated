package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RimRedact(rimRed map[int]string, val int) string {
	res := ""
	for val > 0 {
		switch {
			case val >= 40 && val < 50:
				val -= 40
				res += rimRed[40]
			case val < 40 && val >= 10:
				val -= 10
				res += rimRed[10]
			case val >= 50 && val < 90:
				val -= 50
				res += rimRed[50]
			case val >= 90 && val < 100:
				val -= 90
				res += rimRed[90]
			case val >= 100:
				val -= 100
				res += rimRed[100]
			case val < 10:
				res += rimRed[val]
				val -= val
		}
	}
	return res
}






func Calulates(expr string) string {
	mapRim := map[string]int{
		"I": 1,
		"II": 2,
		"III": 3,
		"IV": 4,
		"V": 5,
		"VI": 6,
		"VII": 7,
		"VIII": 8,
		"IX": 9,
		"X": 10,
		"XL": 40,
		"L": 50,
		"XC": 90,
		"C": 100,
	}

	mapRimRev := map[int]string{
		1: "I",
		2: "II",
		3: "III",
		4: "IV",
		5: "V",
		6: "VI",
		7: "VIII",
		8: "IX",
		10: "X",
		40: "XL",
		50: "L",
		90: "XC",
		100: "C",
	}

	mapArb := map[string]int {
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"10": 10,
	}
	
	expmas := strings.Split(expr, " ")

	if len(expmas) < 2 || len(expmas) > 3 {
		panic("Expected expresson not in the correct format")
	}

	if val1, ok := mapArb[expmas[0]]; ok {
		if val2, ok := mapArb[expmas[2]]; ok {
			switch expmas[1] {
				case "+":
					return strconv.Itoa(val1 + val2)
				case "-":
					return strconv.Itoa(val1 - val2)
				case "*":
					return strconv.Itoa(val1 * val2)
				case "/":
					return strconv.Itoa(val1 / val2)
				default:
					panic("Expected expression not in the correct format")
			}
		}
		panic("Expected expresson not in the correct format")
	}
	
	if val1, ok := mapRim[expmas[0]]; ok {
		if val2, ok := mapRim[expmas[2]]; ok {
			res := 0
			switch expmas[1] {
			case "+":
				res = val1 + val2
			case "-":
				res = val1 - val2
			case "*":
				res = val1 * val2
			case "/":
				res = val1 / val2
			default:
				panic("Expected expression not in the correct format")
		}
		if res < 0 {
			panic("Expected expression not in the correct format")
		}

		if val, ok := mapRimRev[res]; ok {
			return val
		}
		resst := RimRedact(mapRimRev, res)
		return resst
		}
		panic("Expected expresson not in the correct format")
	}
	panic("Expected expression not in the correct format")
}

func main() {
	var expresson string
	read := bufio.NewReader(os.Stdin)
	strk, _ := read.ReadString('\n')
	expresson = strings.TrimSpace(string(strk))
	fmt.Println(Calulates(expresson))
}