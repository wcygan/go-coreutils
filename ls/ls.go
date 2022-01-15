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

func Run(config Config, out io.Writer) error {
	dirEntries, err := os.ReadDir(config.Directory)
	if err != nil {
		return err
	}

	var names []string
	for _, ent := range dirEntries {
		names = append(names, ent.Name())
	}

	sep := getSeparator(config.LongFormat)
	_, err = fmt.Fprintln(out, strings.Join(names, sep))
	if err != nil {
		return err
	}

	return nil
}

func getSeparator(isLongFormat bool) string {
	if isLongFormat {
		return "\n"
	} else {
		return "  "
	}
}
