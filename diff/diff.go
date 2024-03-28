package diff

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// Diff compares wand and have, returns a colored diff of their differences or an empty string if they are equal.
func Diff[T any](want, have T, opts ...cmp.Option) string {
	return diff(want, have, opts...)
}

// Error is a test helper function for comparing two objects of the same type T.
//
// If want and have are not equal, msg and a colored-diff will be sent to t.Error.
// If want is of type error, cmpOpts.EquateErrors will automatically be called.
func Error[T any](t *testing.T, msg string, want, have T, opts ...cmp.Option) {
	t.Helper()
	if d := diff(want, have, opts...); d != "" {
		t.Error(msg, d)
	}
}

// Fatal is a test helper function for comparing two objects of the same type T.
//
// If want and have are not equal, msg and a colored-diff will be sent to t.Fatal.
// If want is of type error, cmpOpts.EquateErrors will automatically be called.
func Fatal[T any](t *testing.T, msg string, want, have T, opts ...cmp.Option) {
	t.Helper()
	if d := diff(want, have, opts...); d != "" {
		t.Fatal(msg, d)
	}
}

const (
	red   = "\x1b[31m" // set the terminal red color
	green = "\x1b[32m" // set the terminal green color
	clear = "\x1b[0m"  // clears the terminal color
)

func diff[T any](want, have T, opts ...cmp.Option) string {
	switch any(want).(type) {
	case error:
		if len(opts) == 0 {
			opts = append(opts, cmpopts.EquateErrors())
		}
	}

	d := cmp.Diff(want, have, opts...)
	if d == "" {
		return ""
	}

	splits := strings.Split(d, "\n")
	for i, s := range splits {
		switch {
		case strings.HasPrefix(s, "-"):
			splits[i] = red + s + clear
		case strings.HasPrefix(s, "+"):
			splits[i] = green + s + clear
		}
	}

	return strings.Join(splits, "\n")
}
