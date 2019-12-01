package main

import "testing"

func Test_fuel(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{12}, 2},
		{"Example 2", args{14}, 2},
		{"Example 3", args{1969}, 654},
		{"Example 4", args{100756}, 33583},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fuel(tt.args.i); got != tt.want {
				t.Errorf("fuel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fuelForFuel(t *testing.T) {
	type args struct {
		f int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{12}, 2},
		{"Example 2", args{1969}, 966},
		{"Example 3", args{100756}, 50346},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fuelWithAddedFuel(tt.args.f); got != tt.want {
				t.Errorf("fuelWithAddedFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}
