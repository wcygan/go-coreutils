package whoami

import (
	"bytes"
	"fmt"
	"os/user"
	"testing"
)

func TestWhoAmI(t *testing.T) {
	out := &bytes.Buffer{}
	err := Run(out)
	if err != nil {
		return
	}

	u, err := user.Current()
	if err != nil {
		t.Error(err)
	}

	expected := fmt.Sprintln(u.Username)
	result := out.String()

	if result != expected {
		t.Errorf("got `%s` want `%s`", result, expected)
	}
}
