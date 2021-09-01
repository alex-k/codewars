package kata

import (
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "S mkx bocod",
			want: "T poc iwlyo",
		},
		{
			name: "I should have known that you would have a perfect answer for me!!!",
			want: "J vltasl rlhr zdfog odxr ypw atasl rlhr p gwkzzyq zntyhv lvz wp!!!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encode(tt.name, 1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovingShift() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestDecodeShift(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "T poc iwlyo",
			want: "S mkx bocod",
		},
		{
			name: "J vltasl rlhr zdfog odxr ypw atasl rlhr p gwkzzyq zntyhv lvz wp!!!",
			want: "I should have known that you would have a perfect answer for me!!!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decode(tt.name, 1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovingShift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_split(t *testing.T) {
	tests := []struct {
		name    string
		wantRet []string
	}{
		{
			name:    "aaabbbcccdd",
			wantRet: []string{"aaa", "bbb", "ccc", "dd", ""},
		},
		{
			name:    "aaaabbbbccccdddde",
			wantRet: []string{"aaaa", "bbbb", "cccc", "dddd", "e"},
		},
		{
			name:    "abcd",
			wantRet: []string{"a", "b", "c", "d", ""},
		},
		{
			name:    "",
			wantRet: []string{"", "", "", "", ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := split(tt.name, 5); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("split() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func TestMovingShift(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "S mkx bocod",
			want: []string{"T p", "oc ", "iwl", "yo", ""},
		},
		{
			name: "I should have known that you would have a perfect answer for me!!!",
			want: []string{"J vltasl rlhr ", "zdfog odxr ypw", " atasl rlhr p ", "gwkzzyq zntyhv", " lvz wp!!!"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MovingShift(tt.name, 1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MovingShift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDemovingShift(t *testing.T) {
	tests := []struct {
		name string
		arg  []string
	}{
		{
			name: "I should have known that you would have a perfect answer for me!!!",
			arg:  []string{"J vltasl rlhr ", "zdfog odxr ypw", " atasl rlhr p ", "gwkzzyq zntyhv", " lvz wp!!!"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DemovingShift(tt.arg, 1); got != tt.name {
				t.Errorf("DemovingShift() = %v, want %v", got, tt.name)
			}
		})
	}
}
