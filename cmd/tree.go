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
	"github.com/wcygan/go-coreutils/constants"
	"github.com/wcygan/go-coreutils/tree"
	"os"
)

// treeCmd represents the tree command
var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Display a directory tree",
	Long:  `Displays a recursive directory listing program that produces a depth-indented listing of files.`,
	Run:   runTree,
}

func init() {
	rootCmd.AddCommand(treeCmd)
}

func runTree(cmd *cobra.Command, args []string) {
	var dir string
	if len(args) > 0 {
		dir = args[0]
	} else {
		dir = constants.Dot
	}

	err := tree.Run(dir, os.Stdout)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
