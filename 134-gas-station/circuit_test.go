package circuit

import (
	"fmt"
	"testing"
)

func Test_canCompleteCircuit(t *testing.T) {
	type args struct {
		gas  []int
		cost []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				gas:  []int{1, 2, 3, 4, 5},
				cost: []int{3, 4, 5, 1, 2},
			},
			want: 3,
		},
		{
			args: args{
				gas:  []int{2, 3, 4},
				cost: []int{3, 4, 3},
			},
			want: -1,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := canCompleteCircuit(tt.args.gas, tt.args.cost); got != tt.want {
				t.Errorf("canCompleteCircuit() = %v, want %v", got, tt.want)
			}
		})
	}
}
