package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")
	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	switch args[0] {
	case "cd":

		if len(args) < 2 {
			return errors.New("path required")
		}

		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	return cmd.Run()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	Hostname, err := os.Hostname()
	Mydir, err := os.Getwd()
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	Hostname = Hostname[:5]
	for {
		fmt.Print("hoost:", Hostname, "user:", user, Mydir, "> ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = execInput(input); err != nil {
			fmt.Fprint(os.Stderr, err)
		}
	}
}
