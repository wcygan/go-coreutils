package whoami

import (
	"fmt"
	"io"
	"os/user"
)

func Run(out io.Writer) error {
	u, err := user.Current()
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(out, u.Username)
	if err != nil {
		return err
	}

	return nil
}
