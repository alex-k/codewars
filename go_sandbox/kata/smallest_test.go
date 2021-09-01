package kata

import (
	"reflect"
	"testing"
)

func TestSmallest(t *testing.T) {
	tests := []struct {
		name string
		n    int64
		want []int64
	}{
		{
			name: "261235",
			n:    261235,
			want: []int64{126235, 2, 0},
		},
		{
			name: "209917",
			n:    209917,
			want: []int64{29917, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Smallest(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Smallest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swap(t *testing.T) {
	type args struct {
		n int64
		x int64
		y int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "261235",
			args: args{
				n: 261235,
				x: 2,
				y: 0,
			},
			want: 126235,
		},
		{
			name: "209917",
			args: args{
				n: 209917,
				x: 0,
				y: 1,
			},
			want: 29917,
		},
		{
			name: "342567",
			args: args{
				n: 342567,
				x: 1,
				y: 3,
			},
			want: 325467,
		},
		{
			name: "10000",
			args: args{
				n: 10000,
				x: 0,
				y: 4,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swap(tt.args.n, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("swap() = %v, want %v", got, tt.want)
			}
		})
	}
}
