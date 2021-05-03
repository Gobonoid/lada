package lada

import (
	"io"
)

import "os"

type Cli struct {
	Version string
	Name string
	output io.Writer
	errorOutput io.Writer
	input io.Reader
}

func NewCli(name string, version string) *Cli {
	return &Cli{
		Name: name,
		Version: version,
		input: os.Stdin,
		output: os.Stdout,
		errorOutput: os.Stderr,
	}
}

func (c *Cli) Output() io.Writer {
	return c.output
}

func (c *Cli) ErrorOutput() io.Writer {
	return c.errorOutput
}

func (c *Cli) Input() io.Reader {
	return c.input
}