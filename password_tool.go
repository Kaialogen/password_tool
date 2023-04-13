package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	password := "examplepassword123"

	// Check password length
	if len(password) < 8 {
		fmt.Println("Password is too short")
	}

	// Check for lowercase characters
	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		fmt.Println("Password does not contain any lowercase characters")
	}

	// Check for uppercase characters
	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		fmt.Println("Password does not contain any uppercase characters")
	}

	// Check for numeric characters
	if !strings.ContainsAny(password, "0123456789") {
		fmt.Println("Password does not contain any numeric characters")
	}

	// Check for special characters
	specialChars := "~`!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?"
	if !strings.ContainsAny(password, specialChars) {
		fmt.Println("Password does not contain any special characters")
	}

	// Check for dictionary words
	dictionary := []string{"password", "123456", "qwerty", "letmein", "monkey", "football"}
	for _, word := range dictionary {
		if strings.Contains(strings.ToLower(password), word) {
			fmt.Println("Password contains a dictionary word")
			break
		}
	}

	// Check for repeating characters
	for i := 0; i < len(password)-2; i++ {
		if password[i] == password[i+1] && password[i+1] == password[i+2] {
			fmt.Println("Password contains repeating characters")
			break
		}
	}

	// Check for sequential characters
	for i := 0; i < len(password)-2; i++ {
		if unicode.IsLetter(rune(password[i])) && unicode.IsLetter(rune(password[i+1])) && unicode.IsLetter(rune(password[i+2])) {
			if password[i]+1 == password[i+1] && password[i+1]+1 == password[i+2] {
				fmt.Println("Password contains sequential characters")
				break
			}
		}
	}

	fmt.Println("Password strength check complete")
}
