package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Enter your password:")
	var password string
	fmt.Scanln(&password)

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

	// Open the rockyou.txt file
	file, err := os.Open("wordlists/rockyou.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the file and split it into lines
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Store each line/word from the file into a slice of strings
	var dictionary []string
	for scanner.Scan() {
		dictionary = append(dictionary, scanner.Text())
	}

	// Check if the password contains a dictionary word
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
