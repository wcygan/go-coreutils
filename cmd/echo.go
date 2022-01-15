/*
Copyright © 2022 Will Cygan <wcygan.io@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bytes"
	"fmt"
	"github.com/wcygan/go-coreutils/echo"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "writes each given string to standard output",
	Long:  `echo writes each given string to standard output, with a space between each and a newline after the last one.`,
	Run:   setup,
}

func init() {
	rootCmd.AddCommand(echoCmd)
}

func setup(cmd *cobra.Command, args []string) {
	var in io.Reader
	out := os.Stdout

	if len(args) > 0 {
		var sb strings.Builder
		for idx, str := range args {
			var salt string
			if idx == len(args)-1 {
				salt = "\n"
			} else {
				salt = " "
			}

			sb.WriteString(str + salt)
		}
		in = bytes.NewReader([]byte(sb.String()))
	} else {
		in = os.Stdin
	}

	err := echo.Run(in, out)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
