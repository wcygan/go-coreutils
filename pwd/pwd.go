package pwd

import (
	"fmt"
	"io"
	"os"
)

// Run prints the current directory to "out"
func Run(out io.Writer) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(out, wd)
	if err != nil {
		return err
	}

	return nil
}
