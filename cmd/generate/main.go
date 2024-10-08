package main

import (
	"flag"
	"fmt"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const specialChars = "!@#$%^&*()-_=+[]{}|;:,.<>?/"

func main() {
	//сделать чтобы было две команды list и generate и флаги запрашивались только в команде generate
	
	special := flag.Bool("special", false, "Use special symbols in password generating")
	length := flag.Int("length", 8, "length of the password")
	flag.Parse()	

	fmt.Println("Special:", *special, "| Length: ", *length)
}