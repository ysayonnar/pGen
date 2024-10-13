package main

import (
	"flag"
	"fmt"
	"os"
	"pgen/pkg/generator"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("type help to get list of commands")
		return
	}
	command := os.Args[1]
	//test
	switch command {
	case "help":
		{
			fmt.Println("Commands: list - get list of password, delete - delete file with passwords, generate - generate passworrd(flags: special, length), help - list of commands")
		}
	case "generate":
		{
			special := flag.NewFlagSet("generate", flag.ExitOnError)
			useSpecial := special.Bool("special", false, "Use special symbols in password generating")
			length := special.Int("length", 8, "length of the password")

			err := special.Parse(os.Args[2:])

			if err != nil {
				fmt.Println("Error parsing flags: ", err)
				return
			}

			password := generator.GeneratePassword(*useSpecial, *length)

			fmt.Println("~~password~~ ", password)

		}
	case "list":
		{
			fmt.Println("~~~List of passwords~~~")
			generator.PrintList()
		}
	case "delete":
		{
			generator.Delete()
		}
	default:
		{
			fmt.Println("type help to get list of commands")
		}
	}
}
