package exec

import (
	"context"
	"errors"
	"fmt"
	"io"
	"sync"
)

type StreamType int

const (
	// Stdout specifies stdout stream should pipe to next process's stdin
	Stdout StreamType = iota
	// Stderr specifies stdout stream should pipe to next process's stdin
	Stderr StreamType = iota
	// Null specifies both output streams are not piped. Thus it is only allowed
	// on the last PipeSpec element.
	Null StreamType = iota
)

func (t StreamType) String() string {
	switch t {
	case Stdout:
		return "stdout"
	case Stderr:
		return "stderr"
	}
	return "invalid stream type"
}

// PipeSpec is used to build Pipeline. It specifies each process's command
// option and pipe operation.
type PipeSpec struct {
	// Stream options (Stdin, Stdout, Stderr) are ignored and override by
	// PipeSpec.Next and PipeSpec.Other.
	CmdOpt CommandOpts

	// Next specifies which output stream (one of Stdout, Stderr) is piped to next
	// process (or Null for last process in pipeline and you decided to ignore).
	Next StreamType

	// Other specifies where output stream goes which is not specified by Next.
	//
	// e.g. Stderr if Next is Stdout.
	Other io.Writer
}

// Pipeline represents complete system command pipeline operation.
type Pipeline struct {
	cmds   []Command
	output io.Reader

	chErr       chan PipelineError
	startWaiter sync.Once
}

func nextStdout(
	ctx context.Context,
	cp CommandProvider,
	index int,
	pipeSpec *PipeSpec,
) (Command, io.ReadCloser, error) {
	pipeSpec.CmdOpt.Stderr = pipeSpec.Other
	cmd := cp.CommandContext(ctx, pipeSpec.CmdOpt)
	prev, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get standard stream %d: %w", index, err)
	}
	return cmd, prev, nil
}

func nextStderr(
	ctx context.Context,
	cp CommandProvider,
	index int,
	pipeSpec *PipeSpec,
) (Command, io.ReadCloser, error) {
	pipeSpec.CmdOpt.Stdout = pipeSpec.Other
	cmd := cp.CommandContext(ctx, pipeSpec.CmdOpt)
	prev, err := cmd.StderrPipe()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get standard stream %d: %w", index, err)
	}
	return cmd, prev, nil
}

func nextNull(
	ctx context.Context,
	cp CommandProvider,
	pipeSpecsLen int,
	index int,
	pipeSpec *PipeSpec,
) (Command, io.ReadCloser, error) {
	cmd := cp.CommandContext(ctx, pipeSpec.CmdOpt)
	if pipeSpecsLen-1 != index {
		return nil, nil, errors.New("attempted to pipe to null for non-terminal pipe")
	}
	return cmd, nil, nil
}

// NewPipeline creates pipeline from pipeline specs. ctx should have cancel
// func to handle errorneous situations.
func NewPipeline(
	ctx context.Context,
	cp CommandProvider,
	pipeSpecs []PipeSpec,
) (*Pipeline, error) {
	if len(pipeSpecs) == 0 {
		return nil, errors.New("len(pipeSpecs) == 0")
	}

	prev := pipeSpecs[0].CmdOpt.Stdin
	var err error
	cmds := make([]Command, 0, len(pipeSpecs))
	for i, pipeSpec := range pipeSpecs {
		pipeSpec.CmdOpt.Stdin = prev
		var cmd Command
		switch pipeSpec.Next {
		case Stdout:
			cmd, prev, err = nextStdout(ctx, cp, i, &pipeSpec)
		case Stderr:
			cmd, prev, err = nextStderr(ctx, cp, i, &pipeSpec)
		case Null:
			cmd, prev, err = nextNull(ctx, cp, len(pipeSpecs), i, &pipeSpec)
		}
		if err != nil {
			return nil, err
		}
		cmds = append(cmds, cmd)
	}
	return &Pipeline{
		cmds:        cmds,
		output:      prev,
		chErr:       make(chan PipelineError),
		startWaiter: sync.Once{},
	}, nil
}

// Start starts each pipeline command (which is [[exec.Cmd]]). It does not abort
// if any error has occured so client is responsible to call [Pipeline.Cancel]
// regardless with result.
func (p *Pipeline) Start() error {
	var errs []error
	for _, cmd := range p.cmds {
		if err := cmd.Start(); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func (p *Pipeline) wait() {
	chErr := make(chan PipelineError)
	for index, cmd := range p.cmds {
		go func(index int, cmd Command) {
			chErr <- PipelineError{cmd.Wait(), index, cmd.Path()}
		}(index, cmd)
	}
	for i := 0; i < len(p.cmds); i += 1 {
		if err := <-chErr; err.err != nil {
			p.chErr <- err
		}
	}
	close(p.chErr)
}

// Wait returns channel which will send errors caused by command. It closes
// if every process in pipeline has exited. Closing without sending any error
// means all processes have exited normally. It is safe to call Wait multiple
// times.
func (p *Pipeline) Wait() <-chan PipelineError {
	p.startWaiter.Do(func() { go p.wait() })
	return p.chErr
}

// Cancel calls cancel and wait for all pipeline processes to exit ignoring any
// error. cancel must be cancel func of ctx which passed to [NewPipeline] to
// initialize p.
func (p *Pipeline) Cancel(cancel context.CancelFunc) {
	cancel()
	for range p.Wait() {
	}
}

// Any waits for pipeline processes and returns if any of process has exited
// with error or all processes have exited normally. If one or more process
// have exited with errors, it calls cancel.
func (p *Pipeline) Any(cancel context.CancelFunc) PipelineErrors {
	var errs []PipelineError
	for err := range p.Wait() {
		cancel()
		errs = append(errs, err)
	}
	return errs
}

// Output returns last process's output stream in pipeline. If last
// PipeSpec.Next was Null, it returns nil.
func (p *Pipeline) Output() io.Reader {
	return p.output
}
