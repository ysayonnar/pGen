package generator

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const specialChars = "!@#$%^&*()-_=+[]{}|;:,.<>?/"

func GeneratePassword(useSpecial bool, length int) string {
	characters := letters
	if useSpecial {
		characters += specialChars
	}
	password := []string{}

	for i := 0; i < length; i++ {
		rndIndex := rand.Intn(len(characters))
		password = append(password, string(characters[rndIndex]))
	}

	concatedPassword := strings.Join(password, "")

	file, err := os.OpenFile("passwords.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	
	if err != nil {
		fmt.Println("Unable to load file: ", err)
		os.Exit(1)
	}

	defer file.Close()
	file.WriteString(concatedPassword + "\n")
	fmt.Println("Password writen into passwords.txt")

	return strings.Join(password, "")
}

func PrintList(){
	data, err := os.ReadFile("passwords.txt")

	if err != nil{
		fmt.Println("Error reading file: ", err)
	}

	fmt.Println(string(data))
}

func Delete(){
	err := os.Remove("passwords.txt") // Удаляем файл passwords.txt
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
	fmt.Println("File successfully deleted")
}