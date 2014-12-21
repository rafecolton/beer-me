package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var args = os.Args
	if len(args) != 2 {
		usage()
		os.Exit(1)
	}

	var caskfile = os.Args[1]
	contents, err := ioutil.ReadFile(caskfile)
	if err != nil {
		log.Fatalln(err)
	}

	var brew = []string{"install"}
	var cask = []string{"cask", "install"}

	for _, line := range strings.Split(string(contents), "\n") {
		commentPos := strings.Index(line, "#")
		// delete comment
		if commentPos >= 0 {
			line = line[:commentPos]
		}

		words := strings.Split(line, " ")

		// skip blank lines
		if len(words) < 1 {
			continue
		}

		// commands with no args
		if len(words) == 1 {
			continue // this should be error
		}

		switch words[0] {
		case "brew":
			brew = append(brew, words[1])
		case "cask":
			cask = append(cask, words[1])
		default:
			continue // should be error (invalid command)
		}
	}

	for _, args := range [][]string{brew, cask} {
		cmd := exec.Command("brew", args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}

func usage() {
	fmt.Println(`Usage:

beer-me <Caskfile>
`)
}
