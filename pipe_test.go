package exec_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"

	"github.com/crowdigit/exec"
)

// It does the samething as running following command in bash
// ls -1 / | sed 's/$/ oh no/'

func Example() {
	// cancel func is required
	ctx, kill := context.WithCancel(context.Background())
	defer kill()

	// optionally you can add Ctrl-C interrupt handler
	ctx, unregister := signal.NotifyContext(ctx, os.Interrupt)
	defer unregister()

	// optionally you can have io.Writer buffers that will be written with stream
	// that's not piped to next process. e.g. read stderr while piping stdout
	buffers := []*bytes.Buffer{{}, {}}

	// exec.PipeSpec specifies how a command is ran and how pipe operation should
	// work for each command in pipeline
	pipespecs := []exec.PipeSpec{
		{
			// ls -1 /
			CmdOpt: exec.CommandOpts{Path: "ls", Args: []string{"-1", "/"}},
			// pipe Stdout to next process
			Next: exec.Stdout,
			// write other stream (stderr) to io.Writer (buffers[0])
			Other: buffers[0],
		},
		{
			// sed 's/$/ oh no/'
			CmdOpt: exec.CommandOpts{Path: "sed", Args: []string{"s/$/ oh no/"}},
			// last process's output stream is accessible with Pipeline.Output
			Next:  exec.Stdout,
			Other: buffers[1],
		},
	}

	cp := exec.NewCommandProvider()

	pipeline, err := exec.NewPipeline(ctx, cp, pipespecs)
	if err != nil {
		panic(err)
	}

	err = pipeline.Start()
	// Pipeline.Start may partially succeed. Client should call Pipeline.Cancel
	// regardless of the result.
	defer pipeline.Cancel(kill)
	if err != nil {
		panic(err)
	}

	// Client is responsible for consuming Pipeline.Output reader.
	chPipeErr := make(chan error)
	go func() {
		for {
			subBuffer := make([]byte, 1024)
			read, err := pipeline.Output().Read(subBuffer)
			if read > 0 {
				fmt.Println(string(subBuffer[:read]))
			}
			if errors.Is(err, io.EOF) {
				close(chPipeErr)
				return
			} else if err != nil {
				chPipeErr <- err
				return
			}
		}
	}()
	if err := <-chPipeErr; err != nil {
		panic(err)
	}

	if errs := pipeline.Any(kill); errs != nil {
		fmt.Fprintln(os.Stderr, errs)
		for i, err := range errs {
			fmt.Println(err)
			if buffers[errs[i].Index].Len() > 0 {
				fmt.Println(buffers[i].String())
			}
		}
		panic(errs)
	}
}
