package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	"github.com/codegangsta/cli"
	"golang.org/x/crypto/ssh/terminal"
)

var vStdInScanner *bufio.Scanner

func stdInScanner() *bufio.Scanner {
	if vStdInScanner == nil {
		vStdInScanner = bufio.NewScanner(os.Stdin)
	}
	return vStdInScanner
}

func getUserName(ctx *cli.Context) (username string) {
	username = ctx.String("username")

	if username == "" {
		scanner := stdInScanner()
		fmt.Fprint(os.Stderr, "Github Username: ")
		scanner.Scan()
		username = scanner.Text()
	}

	return
}

func getPassword(ctx *cli.Context) (password string) {
	password = ctx.String("password")

	if password == "" {
		fmt.Fprint(os.Stderr, "Github Password: ")
		pwBytes, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			panic(err)
		}
		fmt.Fprint(os.Stderr, "\n")
		password = string(pwBytes)
	}

	return
}

func getOTP() (otp string) {
	scanner := stdInScanner()
	fmt.Fprint(os.Stderr, "Github OTP: ")
	scanner.Scan()
	otp = scanner.Text()

	return
}
