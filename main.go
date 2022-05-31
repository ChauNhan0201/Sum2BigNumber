package main

import (
	"fmt"
	"os"
)

func Swap(str1, str2 *string) {
    *str1, *str2 = *str2, *str1
}

func Sum(num1, num2 string) string {
	var carry uint8
	var res string

	carry = 0
	res = ""
	
	// Make sure num1 longer than num2
	if(len(num1) < len(num2)) {
		Swap(&num1, &num2)
	}

	len1 := len(num1)
	len2 := len(num2)

	for i := 0; i < len2; i++ {
		sum := (num1[len1 - i - 1] - '0') + (num2[len2 - i - 1] - '0') + carry	// return uint8 type
		res = string((sum % 10 + '0')) + res
		carry = sum / 10
	}

	// Add remaining digits of larger number
	for i := len2; i < len1; i++ {
		sum := (num1[len1 - i - 1] - '0') + carry
		res = string((sum % 10 + '0')) + res
		carry = sum / 10
	}

	if (carry == 1) {
		res = string((carry + '0')) + res
	}

	return res
}

func main() {
	num1 := string(os.Args[1])
	num2 := string(os.Args[2])
	fmt.Println(Sum(num1, num2))
}