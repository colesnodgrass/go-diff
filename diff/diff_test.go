package diff

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_diff(t *testing.T) {
	tests := []struct {
		name     string
		want     any
		have     any
		expRed   int
		expGreen int
		expClear int
	}{
		{
			name:     "mismatch int",
			want:     1,
			have:     2,
			expRed:   1,
			expGreen: 1,
			expClear: 2,
		},
		{
			name: "match int",
			want: 1,
			have: 1,
		},
		{
			name:     "mismatch string",
			want:     "a string",
			have:     "b string",
			expRed:   1,
			expGreen: 1,
			expClear: 2,
		},
		{
			name: "match string",
			want: "a string",
			have: "a string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			df := diff(tt.want, tt.have)
			if d := cmp.Diff(tt.expRed, strings.Count(df, red)); d != "" {
				t.Error("red count", d)
			}
			if d := cmp.Diff(tt.expGreen, strings.Count(df, green)); d != "" {
				t.Error("red count", d)
			}
			if d := cmp.Diff(tt.expClear, strings.Count(df, reset)); d != "" {
				t.Error("red count", d)
			}
		})
	}
}
