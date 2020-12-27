package main

import (
	"fmt"
	"io"
)

type stderrLogger struct {
	stderr io.Writer
}

func (l stderrLogger) Actionf(format string, a ...interface{}) {
	fmt.Fprintln(l.stderr, `►`, fmt.Sprintf(format, a...))
}

func (l stderrLogger) Generatef(format string, a ...interface{}) {
	fmt.Fprintln(l.stderr, `✚`, fmt.Sprintf(format, a...))
}

func (l stderrLogger) Waitingf(format string, a ...interface{}) {
	fmt.Fprintln(l.stderr, `◎`, fmt.Sprintf(format, a...))
}

func (l stderrLogger) Successf(format string, a ...interface{}) {
	fmt.Fprintln(l.stderr, `✔`, fmt.Sprintf(format, a...))
}

func (l stderrLogger) Failuref(format string, a ...interface{}) {
	fmt.Fprintln(l.stderr, `✗`, fmt.Sprintf(format, a...))
}
