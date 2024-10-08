# Сonsole utility for generating passwords

## Features

- **Generate Passwords**: Create random passwords with options to include special characters and specify the length.
- **List Passwords**: View a list of all previously generated passwords stored in a text file.
- **Help Command**: Get a list of available commands and their descriptions.

## Installation

To run this utility, you need to have Go installed on your machine. Clone this repository and navigate to the project directory. Then, you can build and run the application using the following commands:

```bash
cd cmd/generate
go build main.go
```

## Usage

- **Generate password**: Generate has 2 flags `--special` for using special symbols, `--length={value}`
  Example:

```bash
./main.exe generate --special --length=14

```

- **List of password**: Writes list of password `(saved in cmd/generate/passwords.txt)`. You can delete `passwords.txt` also.
  Example:

```bash
./main.exe list
./main.exe delete
```

- **Help**

```bash
./main.exe help

```
