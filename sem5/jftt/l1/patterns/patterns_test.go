package patterns_test

import (
	"patterns/patterns"
	"testing"
)

// func TestFA(t *testing.T) {
// 	pattern := "a*"
// 	file_contents := "aaaa"
// 	found := patterns.FA(pattern, file_contents)
// 	want := 0
// 	if found != want {
// 		t.Errorf("want: %v got: %v", want, found)
// 	}
// }

func TestSuffix(t *testing.T) {
	cases := []struct {
		pattern string
		text    string
		want    int
	}{
		{
			pattern: "ababaca",
			text:    "bbbbabab",
			want:    4,
		},
		{
			pattern: "bbabc",
			text:    "",
			want:    0,
		},
		{
			pattern: "bbabc",
			text:    "abcadebba",
			want:    3,
		},
	}
	for _, tt := range cases {
		got := patterns.Suffix(tt.pattern, tt.text)
		if tt.want != got {
			t.Errorf("pattern: %q text: %q want: %v got: %v", tt.pattern, tt.text, tt.want, got)
		}
	}
}
