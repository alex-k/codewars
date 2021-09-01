package kata

import (
	"fmt"
	"testing"
)

func TestLastDigit(t *testing.T) {
	tests := []struct {
		name string
		as []int
		want int
	}{
		{
			name: "",
			as: []int{ 2, 2, 101, 2 },
			want: 6,
		},
		{
			name: "",
			as: []int{7,6,1},
			want: 9,
		},
		{
			name: "",
			as: []int{2,2,2,0},
			want: 4,
		},
		{
			name: "",
			as: []int{12,30,21},
			want: 6,
		},
		{
			name: "",
			as: []int{2,0,1},
			want: 1,
		},
		{
			name: "",
			as: []int{23,3},
			want: 7,
		},
		{
			name: "",
			as: []int{},
			want: 1,
		},
		{
			name: "",
			as: []int{0,0},
			want: 1,
		},
		{
			name: "",
			as: []int{7,6,21},
			want: 1,
		},
		{
			name: "",
			as: []int{499942, 898102, 846073},
			want: 6,
		},
		{
			name: "",
			as: []int{937640,767456,981242},
			want: 0,
		},
		{
			name: "",
			as: []int{0,0,0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.as), func(t *testing.T) {
			if got := LastDigit(tt.as); got != tt.want {
				t.Errorf("LastDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
