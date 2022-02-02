package yes

import (
	"fmt"
	"io"
)

func Run(text string, out io.Writer) error {
	for {
		_, err := fmt.Fprintln(out, text)
		if err != nil {
			return err
		}
	}
}
