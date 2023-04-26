package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Password Checker & Generator")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}
	grid.SetColumnHomogeneous(true)
	grid.SetRowHomogeneous(true)
	win.Add(grid)

	passLabel, err := gtk.LabelNew("Enter password:")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	grid.Attach(passLabel, 0, 0, 1, 1)

	passEntry, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Unable to create entry:", err)
	}
	grid.Attach(passEntry, 1, 0, 1, 1)

	checkButton, err := gtk.ButtonNewWithLabel("Check")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	grid.Attach(checkButton, 0, 1, 1, 1)
	checkButton.Connect("clicked", func() {
		password, _ := passEntry.GetText()
		checkPasswordStrength(password)
	})

	lengthLabel, err := gtk.LabelNew("Password Length:")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	grid.Attach(lengthLabel, 0, 2, 1, 1)

	lengthAdjustment, err := gtk.AdjustmentNew(8, 1, 128, 1, 10, 0)
	if err != nil {
		log.Fatal("Unable to create adjustment:", err)
	}

	lengthSpinButton, err := gtk.SpinButtonNew(lengthAdjustment, 1, 0)
	if err != nil {
		log.Fatal("Unable to create spin button:", err)
	}
	grid.Attach(lengthSpinButton, 1, 2, 1, 1)

	genButton, err := gtk.ButtonNewWithLabel("Generate")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	grid.Attach(genButton, 0, 3, 1, 1)

	genLabel, err := gtk.LabelNew("")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	grid.Attach(genLabel, 1, 3, 1, 1)

	genButton.Connect("clicked", func() {
		length := int(lengthAdjustment.GetValue())
		password := generatePassword(length)
		genLabel.SetText(password)
	})

	win.ShowAll()
	gtk.Main()
}

// ... rest of the code ...

func checkPasswordStrength(password string) {
	score := 0
	score += checkPasswordLength(password)
	score += checkLowerCase(password)
	score += checkUpperCase(password)
	score += checkNumeric(password)
	score += checkSpecialChars(password)
	score += checkDictionaryWords(password)
	score += checkRepeatingChars(password)
	score += checkSequentialChars(password)

	var strength string
	if score >= 6 {
		strength = "strong"
	} else if score >= 4 {
		strength = "medium"
	} else {
		strength = "weak"
	}

	dialog := gtk.MessageDialogNew(nil, gtk.DIALOG_MODAL, gtk.MESSAGE_INFO, gtk.BUTTONS_CLOSE, "Password strength %s", strength)
	dialog.Run()
	dialog.Destroy()
}

func checkPasswordLength(password string) int {
	if len(password) < 8 {
		dialog := gtk.MessageDialogNew(nil, gtk.DIALOG_MODAL, gtk.MESSAGE_WARNING, gtk.BUTTONS_CLOSE, "Password is too short")
		dialog.Run()
		dialog.Destroy()
		return 0
	}
	return 1
}

func checkLowerCase(password string) int {
	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		dialog := gtk.MessageDialogNew(nil, gtk.DIALOG_MODAL, gtk.MESSAGE_WARNING, gtk.BUTTONS_CLOSE, "Password does not contain any lowercase characters")
		dialog.Run()
		dialog.Destroy()
		return 0
	}
	return 1
}

func checkUpperCase(password string) int {
	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		dialog := gtk.MessageDialogNew(nil, gtk.DIALOG_MODAL, gtk.MESSAGE_WARNING, gtk.BUTTONS_CLOSE, "Password does not contain any uppercase characters")
		dialog.Run()
		dialog.Destroy()
		return 0
	}
	return 1
}

func checkNumeric(password string) int {
	if !strings.ContainsAny(password, "0123456789") {
		dialog := gtk.MessageDialogNew(nil, gtk.DIALOG_MODAL, gtk.MESSAGE_WARNING, gtk.BUTTONS_CLOSE, "Password does not contain any numeric characters")
		dialog.Run()
		dialog.Destroy()
		return 0
	}
	return 1
}

func checkSpecialChars(password string) int {
	specialChars := "~`!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?"
	if !strings.ContainsAny(password, specialChars) {
		dialog := gtk.MessageDialogNew(nil, gtk.DIALOG_MODAL, gtk.MESSAGE_WARNING, gtk.BUTTONS_CLOSE, "Password does not contain any special characters")
		dialog.Run()
		dialog.Destroy()
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
			dialog := gtk.MessageDialogNew(nil, gtk.DIALOG_MODAL, gtk.MESSAGE_WARNING, gtk.BUTTONS_CLOSE, "Password contains a dictionary word")
			dialog.Run()
			dialog.Destroy()
			return 0
		}
	}
	return 1
}

func checkRepeatingChars(password string) int {
	for i := 0; i < len(password)-2; i++ {
		if password[i] == password[i+1] && password[i+1] == password[i+2] {
			dialog := gtk.MessageDialogNew(nil, gtk.DIALOG_MODAL, gtk.MESSAGE_WARNING, gtk.BUTTONS_CLOSE, "Password contains repeating characters")
			dialog.Run()
			dialog.Destroy()
			return 0
		}
	}
	return 1
}

func checkSequentialChars(password string) int {
	for i := 0; i < len(password)-2; i++ {
		if unicode.IsLetter(rune(password[i])) && unicode.IsLetter(rune(password[i+1])) && unicode.IsLetter(rune(password[i+2])) {
			if password[i]+1 == password[i+1] && password[i+1]+1 == password[i+2] {
				dialog := gtk.MessageDialogNew(nil, gtk.DIALOG_MODAL, gtk.MESSAGE_WARNING, gtk.BUTTONS_CLOSE, "Password contains sequential characters")
				dialog.Run()
				dialog.Destroy()
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
