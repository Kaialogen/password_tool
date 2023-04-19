package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

func main() {
	fmt.Println("Select an option:")
	fmt.Println("1. Check password strength")
	fmt.Println("2. Generate a password")

	var choice int
	fmt.Scanln(&choice)

	if choice == 1 {
		fmt.Println("Enter your password:")
		var password string
		fmt.Scanln(&password)

		score := 0
		score += checkPasswordLength(password)
		score += checkLowerCase(password)
		score += checkUpperCase(password)
		score += checkNumeric(password)
		score += checkSpecialChars(password)
		score += checkDictionaryWords(password)
		score += checkRepeatingChars(password)
		score += checkSequentialChars(password)

		if score >= 6 {
			fmt.Println("Password is strong")
		} else if score >= 4 {
			fmt.Println("Password is medium")
		} else {
			fmt.Println("Password is weak")
		}
	} else if choice == 2 {
		fmt.Println("Enter password length:")
		var length int
		fmt.Scanln(&length)

		password := generatePassword(length)
		fmt.Println("Generated password:", password)
	} else {
		fmt.Println("Invalid choice")
	}
}

func checkPasswordLength(password string) int {
	if len(password) < 8 {
		fmt.Println("Password is too short")
		return 0
	}
	return 1
}

func checkLowerCase(password string) int {
	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		fmt.Println("Password does not contain any lowercase characters")
		return 0
	}
	return 1
}

func checkUpperCase(password string) int {
	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		fmt.Println("Password does not contain any uppercase characters")
		return 0
	}
	return 1
}

func checkNumeric(password string) int {
	if !strings.ContainsAny(password, "0123456789") {
		fmt.Println("Password does not contain any numeric characters")
		return 0
	}
	return 1
}

func checkSpecialChars(password string) int {
	specialChars := "~`!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?"
	if !strings.ContainsAny(password, specialChars) {
		fmt.Println("Password does not contain any special characters")
		return 0
	}
	return 1
}

func checkDictionaryWords(password string) int {
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
			return 0
		}
	}
	return 1
}

func checkRepeatingChars(password string) int {
	for i := 0; i < len(password)-2; i++ {
		if password[i] == password[i+1] && password[i+1] == password[i+2] {
			fmt.Println("Password contains repeating characters")
			return 0
		}
	}
	return 1
}

func checkSequentialChars(password string) int {
	for i := 0; i < len(password)-2; i++ {
		if unicode.IsLetter(rune(password[i])) && unicode.IsLetter(rune(password[i+1])) && unicode.IsLetter(rune(password[i+2])) {
			if password[i]+1 == password[i+1] && password[i+1]+1 == password[i+2] {
				fmt.Println("Password contains sequential characters")
				return 0
			}
		}
	}
	return 1
}

func generatePassword(length int) string {
	rand.Seed(time.Now().UnixNano())
	var password strings.Builder
	for i := 0; i < length; i++ {
		char := byte(rand.Intn(94) + 33)
		password.WriteByte(char)
	}
	return password.String()
}
