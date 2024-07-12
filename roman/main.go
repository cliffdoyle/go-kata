package main

import (
	// "fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
	// "golang.org/x/text/number"
)

func Rm(number int) (string, string) {
	var result, ans string

	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	calculations := []string{"M", "M-C", "D", "D-C", "C", "C-X", "L", "L-X", "X", "X-I", "V", "V-I", "I"}

	for i, val := range values {
		for number >= val {
			number -= val
			result += roman[i]
			if len(ans) > 0 {
				ans += "+"
			}
			ans += calculations[i]
		}
	}
	return result, ans
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil || number < 0 || number >= 4000 {
		z01.PrintRune('0')
		return
	}

	result, ans := Rm(number)

	PrintStr(ans)
	z01.PrintRune('\n')
	PrintStr(result)
	z01.PrintRune('\n')

	// fmt.Println(ans)
	// fmt.Println(result)
}
func PrintStr(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
}
