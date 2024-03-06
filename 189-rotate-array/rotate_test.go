package rotate

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		nums  []int
		steps int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums:  []int{1, 2, 3, 4, 5, 6, 7},
				steps: 3,
			},
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name: "2",
			args: args{
				nums:  []int{-1, -100, 3, 99},
				steps: 2,
			},
			want: []int{3, 99, -1, -100},
		},
		{
			name: "3",
			args: args{
				nums:  []int{-1},
				steps: 2,
			},
			want: []int{-1},
		},
		{
			name: "4",
			args: args{
				nums:  []int{1, 2},
				steps: 3,
			},
			want: []int{2, 1},
		},
		{
			name: "5",
			args: args{
				nums:  []int{1, 2},
				steps: 1,
			},
			want: []int{2, 1},
		},
		{
			name: "6",
			args: args{
				nums:  []int{1, 2, 3},
				steps: 4,
			},
			want: []int{3, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.steps)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("removeElement() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
