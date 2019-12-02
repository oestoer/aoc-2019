package main

import (
	"reflect"
	"testing"
)

func Test_intcodeComputer(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Example1", args: args{[]int{1,0,0,0,99}}, want: []int{2,0,0,0,99}},
		{name: "Example2", args: args{[]int{2,3,0,3,99}}, want: []int{2,3,0,6,99}},
		{name: "Example3", args: args{[]int{2,4,4,5,99,0}}, want: []int{2,4,4,5,99,9801}},
		{name: "Example4", args: args{[]int{1,1,1,4,99,5,6,0,99}}, want: []int{30,1,1,4,2,5,6,0,99}},
		{name: "Example5", args: args{[]int{99,4,4,5,99,0}}, want: []int{99,4,4,5,99,0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := intcodeComputer(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intcodeComputer() = %v, want %v", got, tt.want)
			}
		})
	}
}
