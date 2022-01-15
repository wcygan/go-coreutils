package ls

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Config struct {
	LongFormat bool
	Directory  string
}

func (c Config) separator() string {
	if c.LongFormat {
		return "\n"
	} else {
		return "  "
	}
}

func Run(config Config, out io.Writer) error {
	dirEntries, err := os.ReadDir(config.Directory)
	if err != nil {
		return err
	}

	var names []string
	for _, ent := range dirEntries {
		names = append(names, ent.Name())
	}

	_, err = fmt.Fprintln(out, strings.Join(names, config.separator()))
	if err != nil {
		return err
	}

	return nil
}
