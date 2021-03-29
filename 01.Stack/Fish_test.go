package stackExam

import (
	"reflect"
	"testing"
)

func TestFish(t *testing.T) {
	type args struct {
		Size []int
		Dir []int
	}

	tests := []struct{
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{
				Size: []int{4, 2, 5, 3, 1},
				Dir:  []int{1, 1, 0, 0, 0},
			},
			want: 3,
		},
		{
			name: "t2",
			args: args{
				Size: []int{4,3,2,1,5},
				Dir:  []int{1,1,1,1,0},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Soluction(tt.args.Size,tt.args.Dir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("soluction() = %v, want %v", got, tt.want)
			}
		})
	}
}