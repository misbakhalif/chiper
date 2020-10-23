package main

import "fmt"
import "unicode"
import "strings"

func main() {
	fmt.Println(caesarChiper("mobil", 3)) 
	fmt.Println(caesarChiper("semantic", 1))
	fmt.Println(caesarChiper("drop table", 5)) 
	fmt.Println(caesarChiper("database", 9)) 
	fmt.Println(caesarChiper("solar", 2))
	fmt.Println(caesarChiper("alter", 3)) 
}

func caesarChiper(text string, move int) string {
	var (
		ALFABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		alfabet = "abcdefghijklmnopqrstuvwxyz"
		result = ""
	)

  	for _, j := range text {
		if unicode.IsUpper(j) {
			index := strings.IndexAny(ALFABET, string(j)) + move
			if index >= 26 {
				index %= 26
			} else if index < 0 {
				index = (26 + (index % 26))
			}
			result += string(ALFABET[index])
		} else if unicode.IsLower(j) {
			index := strings.IndexAny(alfabet, string(j)) + move
			if index >= 26 {
				index %= 26
			} else if index < 0 {
				index = (26 + (index % 26))
			}
			result += string(alfabet[index])
		} else {
			result += string(j)
		}
	}
	return result
}