package main

import "testing"

type args struct {
	grid Grid
}

var grid1 = Grid{
	{0, 1, 0, 1, 0},
	{1, 1, 0, 1, 1},
	{0, 1, 1, 1, 0},
	{1, 1, 0, 1, 1},
	{0, 1, 0, 1, 0},
}

var grid2 = Grid{
	{0, 1, 1, 0},
	{1, 1, 1, 0},
	{0, 1, 1, 0},
	{0, 0, 1, 0},
	{0, 0, 1, 0},
}

var grid3 = Grid{
	{1},
}

var grid4 = Grid{}

var tests = []struct {
	name string
	grid Grid
	want int
}{
	{"5x5 grid", grid1, 32},
	{"4x5 grid", grid2, 16},
	{"1x1 grid", grid3, 4},
	{"0x0 grid", grid4, 0},
}

func TestGetPerimeter(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPerimeter(tt.grid); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkGetPerimeter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			GetPerimeter(tt.grid)
		}
	}
}
