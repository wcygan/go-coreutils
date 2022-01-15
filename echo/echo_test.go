package echo

import (
	"bytes"
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

			got := out.String()
			if got != tc.given {
				t.Errorf("got `%s`, want `%s`", got, tc.given)
			}
		})
	}
}
