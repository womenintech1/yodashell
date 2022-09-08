# yodashell

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("YoDaShell ------>")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

var ErrNoPath = errors.New("Pfad erforderlich")

func execInput(input string) error {

	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	switch args[0] {

	case "cd":

		if len(args) < 2 {
			return ErrNoPath
		}

		return os.Chdir(args[1])
	case "exit":

		os.Exit(0)

	case "mkdir":
		err := os.Mkdir("testdir", 0750)
		if err != nil && !os.IsExist(err) {
			log.Fatal(err)
		}
		err = os.WriteFile("testdir/testfile.txt", []byte(" eine Texdatei erstellen!"), 0660)
		if err != nil {
			log.Fatal(err)

		}

	case "ls":
		files, err := ioutil.ReadDir("")
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			fmt.Println(f.Name())
		}
	case "ls-l":
		out, err := exec.Command("ls", "-l").Output()
		if err != nil {
			fmt.Printf("%s", err)
		}
		output := string(out[:])
		fmt.Println(output)
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

