package main

import "testing"

func Test_shortestManhattanDistance(t *testing.T) {
	type args struct {
		inputs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				inputs: []string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"},
			},
			want: 159,
		},
		{
			name: "Example2",
			args: args{
				inputs: []string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"},
			},
			want: 135,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestManhattanDistance(tt.args.inputs); got != tt.want {
				t.Errorf("shortestManhattanDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shortestSignalDelay(t *testing.T) {
	type args struct {
		inputs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				inputs: []string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"},
			},
			want: 610,
		},
		{
			name: "Example2",
			args: args{
				inputs: []string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"},
			},
			want: 410,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestSignalDelay(tt.args.inputs); got != tt.want {
				t.Errorf("shortestSignalDelay() = %v, want %v", got, tt.want)
			}
		})
	}
}
