package x

import (
	"bytes"
	"fmt"
	"os/exec"
)

// Run is similar to (*exec.Cmd).Run() but conveniently wraps exec.ExitError with stderr:
//
//	err1 := exec.Command("go", "run").Run() // standard way
//	err2 := Run(exec.Command("go", "run"))  // with this function
//
//	err1.Error() == "exit status 1"
//	err2.Error() == "exit status 1: go: no go files listed"
//
// The underlying *exec.ExitError remains recoverable:
//
//	if err := new(exec.ExitError); errors.As(err2, &err) {
//		fmt.Println(err.UserTime())
//	}
func Run(c *exec.Cmd) error {
	var stderr bytes.Buffer
	c.Stderr = &stderr
	err := c.Run()
	if err, ok := err.(*exec.ExitError); ok {
		err.Stderr = bytes.TrimSpace(stderr.Bytes())
		if len(err.Stderr) > 0 {
			return fmt.Errorf("%w: %s", err, err.Stderr)
		}
	}
	return err
}

// Output is similar to (*exec.Cmd).Output() but conveniently wraps exec.ExitError with stderr:
//
//	_, err1 := exec.Command("ls", "").Output() // standard way
//	_, err2 := Output(exec.Command("ls", ""))  // with this function
//
//	err1.Error() == "exit status 2"
//	err2.Error() == "exit status 2: ls: cannot access '': No such file or directory"
//
// The underlying *exec.ExitError remains recoverable:
//
//	if err := new(exec.ExitError); errors.As(err2, &err) {
//		fmt.Println(err.UserTime())
//	}
func Output(c *exec.Cmd) ([]byte, error) {
	var stdout, stderr bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stderr
	if err := c.Run(); err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			err.Stderr = bytes.TrimSpace(stderr.Bytes())
			if len(err.Stderr) > 0 {
				return nil, fmt.Errorf("%w: %s", err, err.Stderr)
			}
		}
		return nil, err
	}
	return stdout.Bytes(), nil
}
