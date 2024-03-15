package lab2

import (
	"fmt"
	"unicode"
	"strconv"
	"strings"
)

// TODO: document this function.
// PrefixToPostfix converts
func CalculatePostfix(input string) (int, error) {

	arr := strings.Split(input," ")
	nums := make([]int, 0,3)
	var res int

	for _, v := range arr {
		if unicode.IsNumber(v) {
			num, _ := strconv.Atoi(v)
			nums = append(nums, num)
		}
		n1 := nums[0]
		n2 := nums[1]
		switch v {
		case "+":
			res = n1 + n2
		case "-":
			res = n1 - n2
		case "*":
			res = n1 * n2
		case "/":
			res = n1 / n2
		case "^":
		default:
			//error
		}
	}
	return res, fmt.Errorf("TODO")
}
