package cmd

import (
	"bytes"
	"io"
	"os"
	"strings"
)

// combineArgs joins all args into a string with each arg separated by a space
func combineArgs(args []string) string {
	return strings.Join(args[:], " ")
}

// argsOrStdin if args are present, use these args. Otherwise, get stdin
func argsOrStdin(args []string) io.Reader {
	if len(args) > 0 {
		return bytes.NewReader([]byte(combineArgs(args)))
	} else {
		return os.Stdin
	}
}
