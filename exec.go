package finder

import (
	"fmt"
)

// executor defines the type for execution functions.
// Parameters:
// - path: A string representing the path for execution.
// Returns:
// - error: An error encountered during execution, if any.
type executor func(string) error

// Execute represents a named execution operation.
type Execute struct {
	id   string
	exec executor
}

// Exec executes the associated function for the given path.
// Parameters:
// - path: A string representing the path to execute the function on.
// Returns:
// - error: An error returned by the execution function, if any.
func (exec Execute) Exec(path string) error {
	return exec.exec(path)
}

// ID returns the unique identifier for the execution.
// Returns:
// - string: The identifier for the execution.
func (exec Execute) ID() string {
	return exec.id
}

// WithExec creates a execution operation and adds it to the Finder configuration.
// Parameters:
// - id: A unique identifier for the execution.
// - exec: An execution function of type executor.
// Returns:
// - Option: A function that modifies the Finder's executes list.
func WithExec(id string, exec executor) Option {
	return func(f *Finder) {
		f.executes = append(f.executes, Execute{
			id:   id,
			exec: exec,
		})
	}
}

// WithPrintExec creates a predefined execution operation that prints the path to stdout.
// Returns:
// - Option: A function that modifies the Finder's executes list with the print execution.
func WithPrintExec() Option {
	return WithExec("PrintExec", func(path string) error {
		_, err := fmt.Println(path)
		return err
	})
}

// ExecRuntimeError represents an error occurring during execution at runtime.
type ExecRuntimeError struct {
	value string
}

// NewExecRuntimeError creates a new ExecRuntimeError.
// Parameters:
// - value: A string representing the erroneous value or path.
// Returns:
// - ExecRuntimeError: A new instance of ExecRuntimeError.
func NewExecRuntimeError(value string) ExecRuntimeError {
	return ExecRuntimeError{
		value: value,
	}
}

// Error implements the error interface for ExecRuntimeError.
// Returns:
// - string: A formatted error message describing the runtime error.
func (e ExecRuntimeError) Error() string {
	return fmt.Sprintf("unexpected exec runtime error occures for '%s'", e.value)
}
