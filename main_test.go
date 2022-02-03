package main

import (
	"fmt"
	"github.com/Netflix/go-expect"
	pseudotty "github.com/creack/pty"
	"github.com/hinshun/vt10x"
	"os"
	"strings"
	"testing"
	"time"
)

func TestAskQuestion(t *testing.T) {

	pty, tty, err := pseudotty.Open()
	if err != nil {
		t.Fatalf("failed to open pseudotty: %v", err)
	}
	var donec chan struct{}
	term := vt10x.New(vt10x.WithWriter(tty))
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

	console, err := expect.NewConsole(
		expect.WithStdin(pty),
		expect.WithStdout(term),
		expect.WithCloser(pty, tty),
		expectNoError(t),
		expect.WithDefaultTimeout(time.Second),
	)

	if err != nil {
		t.Fatalf("failed to create console: %v", err)
	}

	defer console.Close()

	donec = make(chan struct{})
	go func() {
		defer close(donec)
		console.ExpectString("Do you like pizza?")
		console.SendLine("y")
		console.ExpectString("Answer true")
		console.ExpectEOF()
	}()

	c := &Config{
		stdin:  os.Stdin,
		stderr: os.Stderr,
		stdout: os.Stdout,
	}
	c.stdout = console.Tty()
	c.stderr = console.Tty()
	c.stdin = console.Tty()

	AskQuestion(c)

	if err := console.Tty().Close(); err != nil {
		t.Errorf("error closing Tty: %v", err)
	}
	<-donec

	output := strings.TrimSpace(expect.StripTrailingEmptyLines(term.String()))
	expectedOutput := `? Do you like pizza? Yes                                                        
Answer true`
	if output != expectedOutput {
		t.Fatalf("Unexpected output.\nExpected: \n%s ; \nFound: \n%s", expectedOutput, output)
	}
}

func expectNoError(t *testing.T) expect.ConsoleOpt {
	return expect.WithExpectObserver(
		func(matchers []expect.Matcher, buf string, err error) {
			if err == nil {
				return
			}
			if len(matchers) == 0 {
				t.Fatalf("Error occurred while matching %q: %s\n", buf, err)
			} else {
				var criteria []string
				for _, matcher := range matchers {
					criteria = append(criteria, fmt.Sprintf("%q", matcher.Criteria()))
				}
				t.Fatalf("Unexpected output; expected: %s ; got %q: \nError: %s\n", strings.Join(criteria, ", "), buf, err)
			}
		},
	)
}
