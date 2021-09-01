package kata

import (
	"reflect"
	"testing"
)

const (
	ANN int = iota
	JOHN
)

func TestAnn(t *testing.T) {
	type args struct {
		person int
		n      int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "JOHN 11",
			args: args{
				person: JOHN,
				n:      11,
			},
			want: []int{0, 0, 1, 2, 2, 3, 4, 4, 5, 6, 6},
		},
		{
			name: "ANN 6",
			args: args{
				person: ANN,
				n:      6,
			},
			want: []int{1, 1, 2, 2, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int
			switch tt.args.person {
			case ANN:
				got = Ann(tt.args.n)
			case JOHN:
				got = John(tt.args.n)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ann() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args struct {
		person int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "John 75",
			args: args{
				person: JOHN,
				n:      75,
			},
			want: 1720,
		},
		{
			name: "Ann 115",
			args: args{
				person: ANN,
				n:      115,
			},
			want: 4070,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got int
			switch tt.args.person {
			case ANN:
				got = SumAnn(tt.args.n);
			case JOHN:
				got = SumJohn(tt.args.n);
			}
			if  got != tt.want {
				t.Errorf("SumJohn() = %v, want %v", got, tt.want)
			}
		})
	}
}