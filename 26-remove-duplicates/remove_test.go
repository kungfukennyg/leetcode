package removedupes

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name     string
		args     args
		want     int
		wantNums []int
	}{
		{
			name:     "1",
			args:     args{nums: []int{1, 1, 2}},
			want:     2,
			wantNums: []int{1, 2, -1},
		},
		{
			name:     "2",
			args:     args{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
			want:     5,
			wantNums: []int{0, 1, 2, 3, 4, -1, -1, -1, -1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			} else if !reflect.DeepEqual(tt.wantNums, tt.args.nums) {
				t.Errorf("removeDuplicates() nums = %v, want %v", tt.args.nums, tt.wantNums)

			}
		})
	}
}
