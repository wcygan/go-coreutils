package echo

import (
	"bufio"
	"io"
)

// Run redirects the bytes from "in" to "out"
func Run(in io.Reader, out io.Writer) error {
	s := bufio.NewScanner(in)

	for s.Scan() {
		_, err := out.Write(s.Bytes())
		if err != nil {
			return err
		}
	}

	return nil
}
