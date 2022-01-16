package tree

import (
	"fmt"
	"github.com/wcygan/go-coreutils/constants"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	vertical         = "│"
	singleHorizontal = "─"
	doubleHorizontal = singleHorizontal + singleHorizontal
	branch           = "├"
	ending           = "└─"
	singleSpace      = " "
	tripleSpace      = singleSpace + singleSpace + singleSpace
)

func Run(root string, out io.Writer) error {
	_, err := fmt.Fprintln(out, root)
	if err != nil {
		return err
	}

	walkDir(root, root, out, make([]bool, 0), 0)
	return nil
}

func walkDir(root, current string, out io.Writer, bars []bool, depth int) {
	dirEntries, err := os.ReadDir(current)
	if err != nil {
		return
	}

	for idx, ent := range dirEntries {
		if strings.HasPrefix(ent.Name(), constants.Dot) {
			continue
		}

		isFinalEntry := idx == len(dirEntries)-1
		prefix := getPrefix(isFinalEntry, bars)
		output := fmt.Sprintf("%s%s", prefix, ent.Name())

		_, err := fmt.Fprintln(out, output)
		if err != nil {
			return
		}

		// recurse
		if ent.IsDir() {
			nextDir := filepath.Join(current, ent.Name())
			nextBars := append(bars, !isFinalEntry)
			nextDepth := depth + 1
			walkDir(root, nextDir, out, nextBars, nextDepth)
		}
	}
}

func getPrefix(lastEntry bool, bars []bool) string {
	var sb strings.Builder

	for _, active := range bars {
		if active {
			sb.WriteString(vertical + tripleSpace)
		} else {
			sb.WriteString(singleSpace + tripleSpace)
		}
	}

	if lastEntry {
		sb.WriteString(ending + singleHorizontal + singleSpace)
	} else {
		sb.WriteString(branch + doubleHorizontal + singleSpace)
	}

	return sb.String()
}
