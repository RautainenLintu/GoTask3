package main

import (
	"errors"
	"fmt"
)

func askForInput() (string, error) {
	var input string
	fmt.Println("Enter string: ")
	fmt.Scan(&input)
	if len(input) > 100 {
		return "", errors.New("Input is too long")
	}
	return input, nil
}

func main() {
	input, err := askForInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	// rune_slice := []rune(input)
	char_slice := make([]string, 100)
	for _, ch := range input {
		char_slice = append(char_slice, string(ch))
	}
	fmt.Println(char_slice)
	return
}
