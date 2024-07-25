package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

const (
	lowercaseChars = "abcdefghijklmnopqrstuvwxyz"
	uppercaseChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars    = "0123456789"
	symbolChars    = "!@#$%^&*()-+=?;.,:|"
)

func generatePassword(length int, includeUppercase, includeLowercase, includeNumbers, includeSymbols bool) (string, error) {
	var passwordChars string
	if includeUppercase {
		passwordChars += uppercaseChars
	}
	if includeLowercase {
		passwordChars += lowercaseChars
	}
	if includeNumbers {
		passwordChars += numberChars
	}
	if includeSymbols {
		passwordChars += symbolChars
	}

	if passwordChars == "" {
		return "", fmt.Errorf("no character sets selected")
	}

	password := make([]byte, length)
	for i := range password {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(passwordChars))))
		if err != nil {
			return "", err
		}
		password[i] = passwordChars[num.Int64()]
	}
	return string(password), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	passwordGenerator := `
 ____  _  __             ___  _           _                               
/ ___|(_)/ _|_ __ ___   / _ \| |_   _ ___| |_ _   _ _ __ _   _  ___ _   _ 
\___ \| | |_| '__/ _ \ | | | | | | | / __| __| | | | '__| | | |/ __| | | |
 ___) | |  _| | |  __/ | |_| | | |_| \__ \ |_| |_| | |  | |_| | (__| |_| |
|____/|_|_| |_|  \___|  \___/|_|\__,_|___/\__|\__,_|_|   \__,_|\___|\__,_|
   )__)                                )_)                                
	
 `

	fmt.Println(passwordGenerator)
	fmt.Print("Şifre uzunluğunu girin: ")
	lengthInput, _ := reader.ReadString('\n')
	lengthInput = strings.TrimSpace(lengthInput)
	length, err := strconv.Atoi(lengthInput)
	if err != nil || length < 4 {
		fmt.Println("Geçersiz şifre uzunluğu.")
		os.Exit(1)
	}

	fmt.Print("Büyük harfler (A-Z) içersin mi? (E/H): ")
	includeUppercase := getYesOrNo(reader)

	fmt.Print("Küçük harfler (a-z) içersin mi? (E/H): ")
	includeLowercase := getYesOrNo(reader)

	fmt.Print("Rakamlar (0-9) içersin mi? (E/H): ")
	includeNumbers := getYesOrNo(reader)

	fmt.Print("Semboller (!@#$%^&*+-) içersin mi? (E/H): ")
	includeSymbols := getYesOrNo(reader)

	password, err := generatePassword(length, includeUppercase, includeLowercase, includeNumbers, includeSymbols)
	if err != nil {
		fmt.Println("Şifre oluşturulamadı:", err)
		os.Exit(1)
	}

	fmt.Println("\n\nOluşturulan Şifre:", password)
	fmt.Printf("\n\n")
}

func getYesOrNo(reader *bufio.Reader) bool {
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(strings.ToUpper(response))
	return response == "E"
}