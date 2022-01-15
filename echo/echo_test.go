package echo

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	testCases := []struct {
		name  string
		given string
	}{
		{"1", "Hello World"},
		{"2", "This is a test!"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			in := bytes.NewReader([]byte(tc.given))
			out := &bytes.Buffer{}

			err := Run(in, out)
			if err != nil {
				t.Error(err)
			}

			want := fmt.Sprintln(tc.given)
			got := out.String()
			if got != want {
				t.Errorf("got `%s`, want `%s`", got, want)
			}
		})
	}
}
