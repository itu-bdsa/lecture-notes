package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func printGreeting() {
	// Got the ASCII art from here:
	// http://patorjk.com/software/taag/#p=display&f=Bloody&t=%20%20%20ITU
	fmt.Println("█     █░▓█████  ██▓     ▄████▄   ▒█████   ███▄ ▄███▓▓█████ ")
	fmt.Println("▓█░ █ ░█░▓█   ▀ ▓██▒    ▒██▀ ▀█  ▒██▒  ██▒▓██▒▀█▀ ██▒▓█   ▀ ")
	fmt.Println("▒█░ █ ░█ ▒███   ▒██░    ▒▓█    ▄ ▒██░  ██▒▓██    ▓██░▒███   ")
	fmt.Println("░█░ █ ░█ ▒▓█  ▄ ▒██░    ▒▓▓▄ ▄██▒▒██   ██░▒██    ▒██ ▒▓█  ▄ ")
	fmt.Println("░░██▒██▓ ░▒████▒░██████▒▒ ▓███▀ ░░ ████▓▒░▒██▒   ░██▒░▒████▒")
	fmt.Println("░ ▓░▒ ▒  ░░ ▒░ ░░ ▒░▓  ░░ ░▒ ▒  ░░ ▒░▒░▒░ ░ ▒░   ░  ░░░ ▒░ ░")
	fmt.Println("▒ ░ ░   ░ ░  ░░ ░ ▒  ░  ░  ▒     ░ ▒ ▒░ ░  ░      ░ ░ ░  ░")
	fmt.Println("  ░   ░     ░     ░ ░   ░        ░ ░ ░ ▒  ░      ░      ░  ")
	fmt.Println("    ░       ░  ░    ░  ░░ ░          ░ ░         ░      ░  ░")
	fmt.Println("                        ░                                   ")
	fmt.Println("")
	fmt.Println("                     ▄▄▄█████▓ ▒█████  ")
	fmt.Println("                     ▓  ██▒ ▓▒▒██▒  ██▒")
	fmt.Println("                     ▒ ▓██░ ▒░▒██░  ██▒")
	fmt.Println("                     ░ ▓██▓ ░ ▒██   ██░")
	fmt.Println("                       ▒██▒ ░ ░ ████▓▒░")
	fmt.Println("                       ▒ ░░   ░ ▒░▒░▒░ ")
	fmt.Println("                         ░      ░ ▒ ▒░ ")
	fmt.Println("                       ░      ░ ░ ░ ▒  ")
	fmt.Println("                                  ░ ░  ")
	fmt.Println("")
	fmt.Println("                    ██▓▄▄▄█████▓ █    ██ ")
	fmt.Println("                   ▓██▒▓  ██▒ ▓▒ ██  ▓██▒")
	fmt.Println("                   ▒██▒▒ ▓██░ ▒░▓██  ▒██░")
	fmt.Println("                   ░██░░ ▓██▓ ░ ▓▓█  ░██░")
	fmt.Println("                   ░██░  ▒██▒ ░ ▒▒█████▓ ")
	fmt.Println("                   ░▓    ▒ ░░   ░▒▓▒ ▒ ▒ ")
	fmt.Println("                    ▒ ░    ░    ░░▒░ ░ ░ ")
	fmt.Println("                    ▒ ░  ░       ░░░ ░ ░ ")
	fmt.Println("                    ░              ░     ")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
}

func getName(name string) string {
	fmt.Print("Tell me your name: ")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}

func main() {
	arguments := os.Args[1:]
	var name string
	var prg_name string
	if runtime.GOOS == "windows" {
		prg_name = "hello.exe"
	} else {
		prg_name = "hello"
	}
	if len(arguments) >= 2 {
		fmt.Println("Usage: " + prg_name + " [name]")
		os.Exit(1)
	}

	printGreeting()

	if len(arguments) == 0 {
		name = getName(name)
	} else {
		name = arguments[0]
	}

	fmt.Println("Hej " + strings.TrimSpace(name) + ", welcome at ITU!")
}
