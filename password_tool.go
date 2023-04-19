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

	checkPasswordLength(password)
	checkLowerCase(password)
	checkUpperCase(password)
	checkNumeric(password)
	checkSpecialChars(password)
	checkDictionaryWords(password)
	checkRepeatingChars(password)
	checkSequentialChars(password)

	fmt.Println("Password strength check complete")
}

func checkPasswordLength(password string) {
	if len(password) < 8 {
		fmt.Println("Password is too short")
	}
}

func checkLowerCase(password string) {
	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		fmt.Println("Password does not contain any lowercase characters")
	}
}

func checkUpperCase(password string) {
	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		fmt.Println("Password does not contain any uppercase characters")
	}
}

func checkNumeric(password string) {
	if !strings.ContainsAny(password, "0123456789") {
		fmt.Println("Password does not contain any numeric characters")
	}
}

func checkSpecialChars(password string) {
	specialChars := "~`!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?"
	if !strings.ContainsAny(password, specialChars) {
		fmt.Println("Password does not contain any special characters")
	}
}

func checkDictionaryWords(password string) {
	file, err := os.Open("wordlists/rockyou.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var dictionary []string
	for scanner.Scan() {
		dictionary = append(dictionary, scanner.Text())
	}

	for _, word := range dictionary {
		if strings.Contains(strings.ToLower(password), word) {
			fmt.Println("Password contains a dictionary word")
			break
		}
	}
}

func checkRepeatingChars(password string) {
	for i := 0; i < len(password)-2; i++ {
		if password[i] == password[i+1] && password[i+1] == password[i+2] {
			fmt.Println("Password contains repeating characters")
			break
		}
	}
}

func checkSequentialChars(password string) {
	for i := 0; i < len(password)-2; i++ {
		if unicode.IsLetter(rune(password[i])) && unicode.IsLetter(rune(password[i+1])) && unicode.IsLetter(rune(password[i+2])) {
			if password[i]+1 == password[i+1] && password[i+1]+1 == password[i+2] {
				fmt.Println("Password contains sequential characters")
				break
			}
		}
	}
}