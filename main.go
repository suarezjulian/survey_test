package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"io"
	"os"
)

type Config struct {
	stdin  terminal.FileReader
	stdout terminal.FileWriter
	stderr io.Writer
}

func main() {
	c := &Config{
		stdin:  os.Stdin,
		stderr: os.Stderr,
		stdout: os.Stdout,
	}
	AskQuestion(c)
}

func AskQuestion(c *Config) {
	t := terminal.Stdio{In: c.stdin, Err: c.stderr, Out: c.stdout}

	answer := false
	// perform the questions
	err := survey.AskOne(&survey.Confirm{
		Message: "Do you like pizza?",
	}, &answer, survey.WithStdio(t.In, t.Out, t.Err))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Fprintf(c.stdout, "Answer %t", answer)
}
