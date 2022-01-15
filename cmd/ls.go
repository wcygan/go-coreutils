/*
Copyright © 2022 William Cygan <wcygan.io@gmail.com>

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
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wcygan/go-coreutils/ls"
	"os"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List directory contents",
	Long:  `The ls program lists information about files (of any type, including directories).`,
	Run:   runLs,
}

const (
	long = "long"
	l    = "l"
	cwd  = "."
)

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.Flags().BoolP(long, l, false, "Display directory contents in a column")
}

func runLs(cmd *cobra.Command, args []string) {
	long, err := cmd.Flags().GetBool(long)
	if err != nil {
		return
	}

	var dir string
	if len(args) > 0 {
		dir = args[0]
	} else {
		dir = cwd
	}

	cfg := ls.Config{
		Directory:  dir,
		LongFormat: long,
	}

	err = ls.Run(cfg, os.Stdout)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
