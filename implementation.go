package lab2

import (
	"fmt"
	"unicode"
	"strconv"
)

// TODO: document this function.
// PrefixToPostfix converts
func CalculatePostfix(input string) (int, error) {

	arr := strings.Split(input," ")
	nums := make([]int, 0,3)
	var res int

	for _, v := range arr {
		num, err := strconv.Atoi(v)
		if err != nil {
			var n1,n2 int
			n2, nums = pop(nums)
		    n1, nums = pop(nums)

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
				res = power(n1, n2)
			default:
				return -1, fmt.Errorf("Unsuitable symbol")
			}
			nums = append(nums, res)
		}
		if err == nil {
			nums = append(nums, num)
		}
	}
	return nums[0], nil
}

func pop(stack []int) (int, []int) {
    if len(stack) == 0 {
        return 0, stack
    }
    poppedElement := stack[len(stack)-1]
    stack = stack[:len(stack)-1]
    return poppedElement, stack
}

func power(a, b int) (int) {
	x := 1
	for i := 0; i < b ;i++ {
		x *= a
	}
	return x
}
