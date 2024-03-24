package lab2

import (
	"fmt"
	"strings"
	"strconv"
)

// CalculatePostfix
// This function calculate the result of postfix expression
// The program is passed a string as input
// Than, program try to convert string to number and add it to slice
// if program can't convert it, it checks if string is operator
// if true, program get 2 numbers
// and do calculation
// and return the first element of slice as result

func CalculatePostfix(input string) (int, error) {

	trimedString := strings.Trim(input," ")
	arr := strings.Split(trimedString," ")
	lastElem := arr[len(arr)- 1]
	_, err := strconv.Atoi(lastElem)

	if err == nil {
		return -1, fmt.Errorf("Incorrect postfix expression")
	}
	
	nums := make([]int, 0,3)
	var res int

	for _, v := range arr {
		num, err := strconv.Atoi(v)
		if err != nil {

			if len(nums) < 2 {
				return -1, fmt.Errorf("Incorrect postfix expression")
			}
			
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
