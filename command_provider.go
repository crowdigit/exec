package exec

import (
	"context"
	"io"
	"os/exec"
)

// CommandOpts specifies values that will be assigned to [[os/exec.Cmd]] after
// its creation.
type CommandOpts struct {
	Path   string
	Args   []string
	Env    []string
	Dir    string
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

// CommandProvider wraps [[os/exec.Command]] and [[os/exec.CommandContext]].
type CommandProvider interface {
	// See [[os/exec.Command]]
	Command(opt CommandOpts) Command
	// See [[os/exec.CommandContext]]
	CommandContext(ctx context.Context, opt CommandOpts) Command
}

type commandProviderImpl struct{}

func (i commandProviderImpl) Command(opts CommandOpts) Command {
	cmd := exec.Command(opts.Path, opts.Args...)
	cmd.Env = opts.Env
	cmd.Dir = opts.Dir
	cmd.Stdin = opts.Stdin
	cmd.Stdout = opts.Stdout
	cmd.Stderr = opts.Stderr
	return commandImpl{cmd}
}

func (i commandProviderImpl) CommandContext(ctx context.Context, opt CommandOpts) Command {
	cmd := exec.CommandContext(ctx, opt.Path, opt.Args...)
	cmd.Env = opt.Env
	cmd.Dir = opt.Dir
	cmd.Stdin = opt.Stdin
	cmd.Stdout = opt.Stdout
	cmd.Stderr = opt.Stderr
	return commandImpl{cmd}
}

// NewCommandProvider returns default implementation of [CommandProvider].
func NewCommandProvider() commandProviderImpl {
	return commandProviderImpl{}
}
