package kata

import (
	"reflect"
	"testing"
)

func TestPickPeaks(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want PosPeaks
	}{
		{
			name: "1236412321",
			args: []int{1, 2, 3, 6, 4, 1, 2, 3, 2, 1},
			want: PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}},
		},
		{
			name: "should support finding peaks, but should ignore peaks on the edge of the array",
			args: []int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3},
			want: PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}},
		},
		{
			name: "should support finding peaks; if the peak is a plateau, it should only return the position of the first element of the plateau",
			args: []int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1},
			want: PosPeaks{Pos: []int{3, 7, 10}, Peaks: []int{6, 3, 2}},
		},
		{
			name: "should support finding peaks after the plateau",
			args: []int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1, 2, 3, 2},
			want: PosPeaks{Pos: []int{3, 7, 10, 15}, Peaks: []int{6, 3, 2, 3}},
		},
		{
			name: "should support finding peaks, but should ignore peaks on the edge of the array",
			args: []int{2, 1, 3, 1, 2, 2, 2, 2},
			want: PosPeaks{Pos: []int{2}, Peaks: []int{3}},
		},
		{
			name: "[13 9 -2 -5 8 8 14 -2 -3]",
			args: []int{13, 9, -2, -5, 8, 8, 14, -2, -3},
			want: PosPeaks{Pos: []int{6}, Peaks: []int{14}},
		},
		{
			name: "[-3 12 12 9 -3 14 5 1 -4]",
			args: []int{-3, 12, 12, 9, -3, 14, 5, 1, -4},
			want: PosPeaks{Pos: []int{1,5}, Peaks: []int{12,14}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PickPeaks(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PickPeaks() = %v, want %v", got, tt.want)
			}
		})
	}
}
