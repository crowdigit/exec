package exec

import "fmt"

type PipelineError struct {
	err   error
	Index int
	Path  string
}

func (e PipelineError) Error() string {
	return fmt.Sprintf("pipeline process %d %s has failed: %s", e.Index, e.Path, e.err)
}

func (e PipelineError) Unwrap() error {
	return e.err
}

type PipelineErrors []PipelineError

func (e PipelineErrors) Error() string {
	return "one or more pipeline process has exited with error"
}

func (e PipelineErrors) Unwrap() []error {
	errs := make([]error, 0, len(e))
	for i := range e {
		errs = append(errs, e[i])
	}
	return errs
}
