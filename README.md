# crowdigit/exec

Wrapper library for `os/exec` with testable interface and command piping utilities.

## Usage

Please [see package example](https://pkg.go.dev/github.com/crowdigit/exec#example-package)

### `CommandProvider`/`Command`

Interface/wrapper type for `exec.Command`, `exec.CommandContext`.

#### 1. Initializing System Command

```go
cp := exec.NewCommandProvider()
// CommandOpts's value is assigned to *exec.Cmd exported fields.
opt := exec.CommandOpts{
    Path: "ls",
    Args: []string{"."},
    // ...
}
cmd := cp.Command(opt)
// or with kill (cancel)
// cmd := cp.CommandContext(ctx, opt)
```

#### 2. Executing `Command`

`Command` interface provides the same receivers as `*exec.Cmd`, you can use `Output`, `Start`, etc., anything you want, the same way as you are using `*exec.Cmd`.

```golang
result, err := cmd.Output()
if err != nil {
    // oh no
}
```

#### 3. Accessing `Command` fields

Plus, it add some receivers to access exported field from `*exec.Cmd`, (complete list is not added yet).

```golang
cmd := cp.Command(opt)
// *exec.Path field value is not accessible via [Command] interface
path := cmd.Path()
```

### `PipeSpec`, `Pipeline`

These are utilities for commonly used system command piping operation.

#### 1. Initialization

```go
// 1. cancel func is required to handle error situations
ctx, kill := context.WithCancel(context.Background())
defer kill()

// 2. Define command specs to build command pipeline
// Following results in the same thing as "ls -1 | tail" in bash
pipespecs := []exec.PipeSpec{
    {
        CmdOpt: exec.CommandOpts{Path: "ls -1", Args: []string{}},
        Next:   exec.Stdout,
    },
    {
        CmdOpt: exec.CommandOpts{Path: "tail", Args: []string{}},
        Next:   exec.Stdout,
    },
}

// 3. Initialize pipeline, it does not spawn processes yet.
pipeline, err := exec.NewPipeline(ctx, cp, pipespecs)
if err != nil {
    return fmt.Errorf("failed to initialize pipeline: %w", err)
}
```

 You can also pipe stderr instead of stdout, or read stderr while piping stdout to next process. [See package example](https://pkg.go.dev/github.com/crowdigit/exec#example-package).

#### 2. Start Pipeline

```go
// 1. Start pipeline commands, it actually spawns process.
// There may be partial success state if some command fail.
// Client is responsible to call Pipeline.Cancel regardless to Pipeline.Start result.
defer pipeline.Cancel(kill)
if err = pipeline.Start(); err != nil {
    return fmt.Errorf("failed to start pipeline: %w", err)
}

// 2. Client must consume Pipeline.Output stream (last process's output stream) on separate goroutine,
// (otherwise it blocks) if it was not specified to pipe to Null.
// In this example, it prints the stdout to parent's stdout.
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
    return fmt.Errorf("failed to operate on pipe: %w", err)
}
```

#### 3. Wait & Getting results

```go
// Pipeline.Any blocks until any or all of command in pipeline has exited.
if errs := pipeline.Any(kill); errs != nil {
    fmt.Fprintln(os.Stderr, errs)
    for i, err := range errs {
        fmt.Println(err)
    }
    return errs
}
```

If `Pipeline.Any` returned `nil` slice, all commands has exited normally. Otherwise it's slice of `PipelineError` which implements `error`, `Unwrap` interface and some context values to provide information which command failed.

```go
for _, err := range errors {
    _ = err.Index // command index in pipeline
    _ = err.Path // command Path e.g. /bin/ls
    _ = err.Unwrap() // original error that *exec.Cmd.Wait returned
}
```

After `Pipeline.Any` has returned, optional output stream writers are filled with data. [See package example](https://pkg.go.dev/github.com/crowdigit/exec#example-package) for how to read stdout/stderr.

## Why?

I write program whose main purpose is to invoke system command a lot. (sure, writing bash script is more efficient but Golang is fun). I often find myself writing interface type for `*exec.Cmd` and its mocking interface to write tests. So I made it reusable.

Mock interface is included in `/mocks` directory ftw, functions accepting `CommandProvider` (rather than calling `exec.Command`) is more testable I think.[^1]

[^1]: Some may argue that we shouldn't test *that* code. That may be correct, but I think if its role is like shell script (e.g. execute command and parse standard output), I think it is good to test that code.