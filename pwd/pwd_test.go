package pwd

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestPwd(t *testing.T) {
	expected, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	expected = fmt.Sprintln(expected)

	out := &bytes.Buffer{}
	err = Run(out)
	if err != nil {
		return
	}

	result := out.String()

	if result != expected {
		t.Errorf("got `%s` want `%s`", result, expected)
	}
}
