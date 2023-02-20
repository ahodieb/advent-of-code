package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {

	tests := []struct {
		name   string
		packet string
		want   []string
	}{
		{name: "empty list", packet: "[]", want: []string{"[", "]"}},
		{name: "nested empty lists", packet: "[[], []]", want: []string{"[", "[", "]", "[", "]", "]"}},
		{name: "numbers list", packet: "[1,23,444,56]", want: []string{"[", "1", "23", "444", "56", "]"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.packet); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v (%d), want %v (%d)", got, len(got), tt.want, len(tt.want))
			}
		})
	}
}

func TestInOrder(t *testing.T) {
	tests := []struct {
		name  string
		left  []string
		right []string
		want  bool
	}{
		{
			name:  "compare lists of numbers",
			left:  Parse("[1,1,3,1,1]"),
			right: Parse("[1,1,5,1,1]"),
			want:  true,
		},
		{
			name:  "compare lists of numbers not in order",
			left:  Parse("[1,1,5,1,1]"),
			right: Parse("[1,1,3,1,1]"),
			want:  false,
		},
		{
			name:  "list vs number",
			left:  Parse("[[2,3,4]]"),
			right: Parse("[4]"),
			want:  true,
		},
		{
			name:  "number vs list ",
			left:  Parse("[9]"),
			right: Parse("[[8,7,6]]"),
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InOrder(tt.left, tt.right); got != tt.want {
				t.Errorf("InOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
