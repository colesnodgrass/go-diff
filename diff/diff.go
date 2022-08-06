package diff

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"strings"
	"testing"
)

func Error[T any](t *testing.T, msg string, want, have T, opts ...cmp.Option) {
	t.Helper()
	if d := diff(want, have, opts...); d != "" {
		t.Error(msg, d)
	}
}

func Fatal[T any](t *testing.T, msg string, want, have T, opts ...cmp.Option) {
	t.Helper()
	if d := diff(want, have, opts...); d != "" {
		t.Fatal(msg, d)
	}
}

const (
	red   = "\x1b[31m"
	green = "\x1b[32m"
	clear = "\x1b[0m"
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
