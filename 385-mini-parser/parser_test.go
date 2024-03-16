package parser

import (
	"reflect"
	"testing"
)

func Test_deserialize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *NestedInteger
	}{
		// {
		// 	name: "1",
		// 	args: args{s: "324"},
		// 	want: &NestedInteger{v: 324},
		// },
		{
			name: "2",
			args: args{s: "[123,[456,[789]]]"},
			want: &NestedInteger{
				children: []*NestedInteger{
					{v: 123},
					{
						children: []*NestedInteger{
							{v: 456},
							{children: []*NestedInteger{{v: 789}}},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deserialize(tt.args.s); !reflect.DeepEqual(*got, *tt.want) {
				t.Errorf("deserialize() = %+v, want %+v", *got, *(tt.want))
			}
		})
	}
}
