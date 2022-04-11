package _111

import (
	"fmt"
	"strconv"
	"testing"
)


func Test_II_036(t *testing.T) {
	var stack []int
	// tokens := []string{"2", "1", "+", "3", "*"}
	tokens := []string{"4","13","5","/","+"}
	if len(tokens) < 1{
		return // tokens
	}
	// tokens := ["4","13","5","/","+"]
	// tokens := ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err != nil {
			// 两数出栈
			n := len(stack)
			if n < 2{
				continue
			}
			stack = append(stack[:n-2],calculate(token,stack[n-1],stack[n-2]))
			continue
		}
		stack = append(stack, num)

	}
	fmt.Println("stack:", stack[0])
}

func calculate(v string,num1,num2 int) int{
	switch v {
	case "+":
		return  num1+num2
	case "-":
		return  num1-num2
	case "*":
		return  num1*num2
	case "/":
		return  num2/num1
	default:
		return 0
	}
}
