package echo

import (
	"bufio"
	"fmt"
	"io"
)

// Run redirects the bytes from "in" to "out"
func Run(in io.Reader, out io.Writer) error {
	s := bufio.NewScanner(in)

	for s.Scan() {
		_, err := fmt.Fprintln(out, s.Text())
		if err != nil {
			return err
		}
	}

	return nil
}
