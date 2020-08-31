package main

import (
	"reflect"
	"testing"
)

type args struct {
	a string
	b string
}

var tests = []struct {
	args args
	want float64
}{
	{args{"i love your hair", "i love your long mustache"}, 1.0 / 4.0},
	{args{"sorry i gotta go now", "gotta go now"}, 1.0 / 3.0},
	{args{"too cute", "too short but cute"}, 0.0},
}

func TestGetTrigramSimilarity(t *testing.T) {
	t.Parallel()
	for _, tt := range tests {
		if got := GetTrigramSimilarity(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}

func BenchmarkGetTrigramSimilarity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			GetTrigramSimilarity(tt.args.a, tt.args.b)
		}
	}
}
