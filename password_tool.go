package main

// Import required packages for the program
import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

// Main function, entry point of the program
func main() {

	// Initialize GTK
	gtk.Init(nil)

	// Apply custom CSS styles from the "styles.css" file
	applyCustomCSS("styles.css")

	// Create a new top-level window
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	// Set window title and make it resizable
	win.SetTitle("Password Checker & Generator")
	win.SetResizable(true)
	// Connect the window's "destroy" event to the gtk.MainQuit function
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Create a new grid layout and set its properties
	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}
	grid.SetColumnHomogeneous(true)
	grid.SetRowHomogeneous(true)
	win.Add(grid)

	// Create a label for the password input field and add it to the grid
	passLabel, err := gtk.LabelNew("Enter password:")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	grid.Attach(passLabel, 0, 0, 1, 1)

	// Create the password input field and add it to the grid
	passEntry, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Unable to create entry:", err)
	}
	grid.Attach(passEntry, 1, 0, 1, 1)

	// Create a button to check the password strength and add it to the grid
	checkButton, err := gtk.ButtonNewWithLabel("Check")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	grid.Attach(checkButton, 0, 1, 1, 1)

	// Connect the button's "clicked" event to the checkPasswordStrength function
	checkButton.Connect("clicked", func() {
		password, _ := passEntry.GetText()
		checkPasswordStrength(password)
	})

	// Create a label for the password length spinner and add it to the grid
	lengthLabel, err := gtk.LabelNew("Password Length:")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	grid.Attach(lengthLabel, 0, 2, 1, 1)

	// Create an adjustment for the password length spinner
	lengthAdjustment, err := gtk.AdjustmentNew(8, 1, 64, 1, 10, 0)
	if err != nil {
		log.Fatal("Unable to create adjustment:", err)
	}

	// Create a spinner to adjust password length and add it to the grid
	lengthSpinButton, err := gtk.SpinButtonNew(lengthAdjustment, 1, 0)
	if err != nil {
		log.Fatal("Unable to create spin button:", err)
	}
	grid.Attach(lengthSpinButton, 1, 2, 1, 1)

	// Create toggle switches for character types (lowercase, uppercase, numbers, symbols) and add them to the grid
	// Also create corresponding labels for each toggle switch

	// Lowercase switch and label
	lowercaseSwitch, err := gtk.SwitchNew()
	if err != nil {
		log.Fatal("Unable to create switch:", err)
	}
	grid.Attach(lowercaseSwitch, 1, 3, 1, 1)
	lowercaseLabel, err := gtk.LabelNew("Include lowercase:")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	grid.Attach(lowercaseLabel, 0, 3, 1, 1)

	// Uppercase switch and label
	uppercaseSwitch, err := gtk.SwitchNew()
	if err != nil {
		log.Fatal("Unable to create switch:", err)
	}
	grid.Attach(uppercaseSwitch, 1, 4, 1, 1)
	uppercaseLabel, err := gtk.LabelNew("Include uppercase:")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	grid.Attach(uppercaseLabel, 0, 4, 1, 1)

	// Numbers switch and label
	numbersSwitch, err := gtk.SwitchNew()
	if err != nil {
		log.Fatal("Unable to create switch:", err)
	}
	grid.Attach(numbersSwitch, 1, 5, 1, 1)
	numbersLabel, err := gtk.LabelNew("Include numbers:")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	grid.Attach(numbersLabel, 0, 5, 1, 1)

	// Symbols switch and label
	symbolsSwitch, err := gtk.SwitchNew()
	if err != nil {
		log.Fatal("Unable to create switch:", err)
	}
	grid.Attach(symbolsSwitch, 1, 6, 1, 1)
	symbolsLabel, err := gtk.LabelNew("Include symbols:")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	grid.Attach(symbolsLabel, 0, 6, 1, 1)

	// Create a button to generate a password and add it to the grid
	genButton, err := gtk.ButtonNewWithLabel("Generate")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	grid.Attach(genButton, 0, 7, 1, 1)

	// Create a label to display the generated password and add it to the grid
	genLabel, err := gtk.LabelNew("")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	grid.Attach(genLabel, 1, 7, 1, 1)

	// Connect the generate button's "clicked" event to a function that generates a password with the specified options
	genButton.Connect("clicked", func() {
		length := int(lengthAdjustment.GetValue())
		includeLowercase := lowercaseSwitch.GetActive()
		includeUppercase := uppercaseSwitch.GetActive()
		includeNumbers := numbersSwitch.GetActive()
		includeSymbols := symbolsSwitch.GetActive()

		password := generatePassword(length, includeLowercase, includeUppercase, includeNumbers, includeSymbols)
		genLabel.SetText(password)
	})

	// Display all widgets in the window and start the GTK main loop
	win.ShowAll()
	gtk.Main()
}

// checkPasswordStrength calculates the strength of the input password based on various criteria
// and displays a message dialog with the strength rating
func checkPasswordStrength(password string) {
	// Initialize the password score
	score := 0
	// Calculate the score by checking various criteria
	score += checkPasswordLength(password)
	score += checkLowerCase(password)
	score += checkUpperCase(password)
	score += checkNumeric(password)
	score += checkSpecialChars(password)
	score += checkDictionaryWords(password)
	score += checkRepeatingChars(password)
	score += checkSequentialChars(password)

	// Determine the password strength based on the score
	var strength string
	if score >= 8 {
		strength = "strong"
	} else if score >= 6 {
		strength = "medium"
	} else {
		strength = "weak"
	}

	// Display a message dialog with the password strength
	dialog := gtk.MessageDialogNew(nil, gtk.DIALOG_MODAL, gtk.MESSAGE_INFO, gtk.BUTTONS_CLOSE, "Password strength %s", strength)
	dialog.Run()
	dialog.Destroy()
}

// checkPasswordLength returns 1 if the password length is at least 8 characters, otherwise returns 0 and displays an error
func checkPasswordLength(password string) int {
	if len(password) < 8 {
		dialog := gtk.MessageDialogNew(nil, gtk.DIALOG_MODAL, gtk.MESSAGE_WARNING, gtk.BUTTONS_CLOSE, "Password is too short")
		dialog.Run()
		dialog.Destroy()
		return 0
	}
	return 1
}

// checkLowerCase returns 1 if the password contains at least one lowercase letter, otherwise returns 0 and displays a warning message
func checkLowerCase(password string) int {
	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		dialog := gtk.MessageDialogNew(nil, gtk.DIALOG_MODAL, gtk.MESSAGE_WARNING, gtk.BUTTONS_CLOSE, "Password does not contain any lowercase characters")
		dialog.Run()
		dialog.Destroy()
		return 0
	}
	return 1
}

// checkUpperCase returns 1 if the password contains at least one lowercase letter, otherwise returns 0 and displays a warning message
func checkUpperCase(password string) int {
	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		dialog := gtk.MessageDialogNew(nil, gtk.DIALOG_MODAL, gtk.MESSAGE_WARNING, gtk.BUTTONS_CLOSE, "Password does not contain any uppercase characters")
		dialog.Run()
		dialog.Destroy()
		return 0
	}
	return 1
}

// checkNumeric returns 1 if the password contains at least one numeric character, otherwise returns 0 and displays a warning message
func checkNumeric(password string) int {
	if !strings.ContainsAny(password, "0123456789") {
		dialog := gtk.MessageDialogNew(nil, gtk.DIALOG_MODAL, gtk.MESSAGE_WARNING, gtk.BUTTONS_CLOSE, "Password does not contain any numeric characters")
		dialog.Run()
		dialog.Destroy()
		return 0
	}
	return 1
}

// checkSpecialChars returns 1 if the password contains at least one special character, otherwise returns 0 and displays a warning message
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

// checkDictionaryWords returns 1 if the password does not contain any dictionary words, otherwise returns 0 and displays a warning message
func checkDictionaryWords(password string) int {
	// Open the dictionary file and handle any errors
	file, err := os.Open("wordlists/rockyou.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the dictionary words line by line
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Store the dictionary words in a slice
	var dictionary []string
	for scanner.Scan() {
		dictionary = append(dictionary, scanner.Text())
	}

	// Check if the password contains any dictionary words
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

// checkRepeatingChars returns 1 if the password does not contain repeating characters, otherwise returns 0 and displays a warning message
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

// checkSequentialChars returns 1 if the password does not contain sequential characters, otherwise returns 0 and displays a warning message
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

// generatePassword creates a random password based on the provided criteria and returns it as a string
func generatePassword(length int, includeLowercase, includeUppercase, includeNumbers, includeSymbols bool) string {
	const (
		lowercaseChars = "abcdefghijklmnopqrstuvwxyz"
		uppercaseChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		numberChars    = "0123456789"
		symbolChars    = "~`!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?"
	)

	// Build the allowed character set based on the provided criteria
	var allowedChars string

	if includeLowercase {
		allowedChars += lowercaseChars
	}
	if includeUppercase {
		allowedChars += uppercaseChars
	}
	if includeNumbers {
		allowedChars += numberChars
	}
	if includeSymbols {
		allowedChars += symbolChars
	}

	// Return an empty string if no character types are allowed
	if len(allowedChars) == 0 {
		return ""
	}

	// Generate the random password
	rand.Seed(time.Now().UnixNano())
	var password strings.Builder
	for i := 0; i < length; i++ {
		char := allowedChars[rand.Intn(len(allowedChars))]
		password.WriteByte(char)
	}
	return password.String()
}

// applyCustomCSS applies a custom CSS style from a given CSS file to the GTK application
func applyCustomCSS(cssFile string) {
	// Create a new CSS provider
	cssProvider, err := gtk.CssProviderNew()
	if err != nil {
		log.Fatal("Unable to create CSS provider:", err)
	}

	// Load the custom CSS from the provided file
	err = cssProvider.LoadFromPath(cssFile)
	if err != nil {
		log.Fatal("Unable to load CSS from file:", err)
	}

	// Get the default screen for the application
	screen, err := gdk.ScreenGetDefault()
	if err != nil {
		log.Fatal("Unable to get default screen:", err)
	}

	// Add the custom CSS provider to the application
	gtk.AddProviderForScreen(screen, cssProvider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
}
