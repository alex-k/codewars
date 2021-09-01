package kata

import (
	"testing"
)

func TestSumOfDivided(t *testing.T) {
	tests := []struct {
		name string
		lst  []int
		want string
	}{
		{
			name: "12, 15",
			lst:  []int{12, 15},
			want: "(2 12)(3 27)(5 15)",
		},
		{
			name: "15,21,24,30,45",
			lst:  []int{15, 21, 24, 30, 45},
			want: "(2 54)(3 135)(5 90)(7 21)",
		},
		{
			name: "[107]",
			lst:  []int{107},
			want: "(107 107)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfDivided(tt.lst); got != tt.want {
				t.Errorf("SumOfDivided() = %v, want %v", got, tt.want)
			}
		})
	}
}
