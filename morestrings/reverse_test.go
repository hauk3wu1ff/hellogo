// Package morestrings implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package morestrings

import "testing"

func TestReverseRunes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Basic", args{"Hello, world"}, "dlrow ,olleH"},
		{"Chinese", args{"Hello, 世界"}, "界世 ,olleH"},
		{"Empty-String", args{""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseRunes(tt.args.s); got != tt.want {
				t.Errorf("ReverseRunes() = %v, want %v", got, tt.want)
			}
		})
	}
}
