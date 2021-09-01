package kata

import (
	"reflect"
	"testing"
)

func TestDecompose(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "0",
			want: []string{},
		},
		{
			name: "1",
			want: []string{"1"},
		},
		{
			name: "12/3",
			want: []string{"4"},
		},
		{
			name: "2/3",
			want: []string{"1/2", "1/6"},
		},
		{
			name: "7/15",
			want: []string{"1/3", "1/8", "1/120"},
		},
		{
			name: "21/23",
			want: []string{"1/2", "1/3", "1/13", "1/359", "1/644046"},
		},
		{
			name: "50/4187",
			want: []string{"1/84", "1/27055", "1/1359351420"},
		},
		{
			name: "16/6",
			want: []string{"2", "1/2", "1/6"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decompose(tt.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decompose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFract(t *testing.T) {
	tests := []struct {
		name    string
		want    *Fraction
		wantErr bool
	}{
		{
			"1/2",
			&Fraction{1,2},
			false,
		},
		{
			"3/2",
			&Fraction{3,2},
			false,
		},
		{
			"0.666",
			&Fraction{666,1000},
			false,
		},
		{
			"1.666",
			&Fraction{1666,1000},
			false,
		},
		{
			"0",
			nil,
			true,
		},
		{
			"",
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToFract(tt.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToFract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToFract() got = %v, want %v", got, tt.want)
			}
		})
	}
}