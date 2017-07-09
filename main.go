package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var args = os.Args
	var caskfile = "Caskfile"
	if len(args) > 1 {
		caskfile = args[1]
	}

	contents, err := ioutil.ReadFile(caskfile)
	if err != nil {
		exitUsage(err)
	}

	var brew = []string{"install"}
	var cask = []string{"cask", "install"}
	var tap = []string{"tap", "caskroom/cask"}

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
		case "tap":
			tap = append(tap, words[1])
		default:
			continue // should be error (invalid command)
		}
	}

	for _, args := range [][]string{tap, brew, cask} {
		cmd := exec.Command("brew", args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}

func exitUsage(errs ...error) {
	for _, err := range errs {
		fmt.Println(err.Error())
	}

	fmt.Println(`Usage:

beer-me <Caskfile>
`)
	os.Exit(1)
}
