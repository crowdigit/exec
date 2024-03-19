package exec

import (
	"io"
	"os/exec"
)

// Command wraps initialized [[os/exec.Cmd]].
type Command interface {
	Output() ([]byte, error)

	Start() error
	Wait() error
	Run() error

	StdinPipe() (io.WriteCloser, error)
	StdoutPipe() (io.ReadCloser, error)
	StderrPipe() (io.ReadCloser, error)

	Path() string
}

type commandImpl struct {
	*exec.Cmd
}

func (c commandImpl) Path() string {
	return c.Cmd.Path
}
