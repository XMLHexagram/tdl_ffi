package main

import (
	"bufio"
	"context"
	"errors"
	surveyterm "github.com/AlecAivazis/survey/v2/terminal"
	"github.com/fatih/color"
	"github.com/iyear/tdl/cmd"
	"go.etcd.io/bbolt"
	"os"
	"os/signal"
	"strings"
)

import "C"

func init() {
	color.NoColor = true
	s.ctx, s.cancel = signal.NotifyContext(context.Background(), os.Interrupt)
	redirectStdout()
}

func main() {}

type Service struct {
	legacyStdout *os.File
	w            *os.File
	r            *os.File
	scanner      *bufio.Scanner
	cancel       context.CancelFunc
	ctx          context.Context
}

var s = new(Service)

func redirectStdout() {
	var err error
	s.r, s.w, err = os.Pipe()
	if err != nil {
		panic(err)
	}

	s.legacyStdout = os.Stdout
	os.Stdout = s.w

	s.scanner = bufio.NewScanner(s.r)
}

//export ReadLine
func ReadLine() string {
	s.scanner.Scan()
	if err := s.scanner.Err(); err != nil {
		return err.Error()
	}
	return s.scanner.Text()
}

//export Cancel
func Cancel() {
	s.cancel()
	os.Exit(0)
}

// RunScript example: "tdl dl xxx"
//
//export RunScript
func RunScript(script string) {
	os.Args = strings.Split(script, " ")

	defer s.cancel()

	humanizeErrors := map[error]string{
		bbolt.ErrTimeout:        "Current database is used by another process, please terminate it first",
		surveyterm.InterruptErr: "Interrupted",
	}

	if err := cmd.New().ExecuteContext(s.ctx); err != nil {
		for e, m := range humanizeErrors {
			if errors.Is(err, e) {
				color.Red("%s", m)
				os.Exit(1)
			}
		}

		color.Red("Error: %+v", err)
		os.Exit(1)
	}
}
