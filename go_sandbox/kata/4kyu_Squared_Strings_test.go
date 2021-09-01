package kata

import "testing"

func Test_square_size(t *testing.T) {
	tests := []struct {
		name string
		n int
		want int
	}{
		{
			name: "1",
			n:    1,
			want: 1,
		},
		{
			name: "2",
			n:    2,
			want: 2,
		},
		{
			name: "3",
			n:    3,
			want: 2,
		},
		{
			name: "4",
			n:    4,
			want: 2,
		},
		{
			name: "99",
			n:    99,
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := square_size(tt.n); got != tt.want {
				t.Errorf("square_size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeCode(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "I.was.going.fishing.that.morning.at.ten.o'clock",
			want: "c.nhsoI\nltiahi.\noentinw\ncng.nga\nk..mg.s\n\voao.f.\n\v'trtig",
		},
		{
			name: "Process terminated with status 0 (0 minute(s), 6 second(s))",
			want: "s t setP\n)se(tder\n)e(0a ro\n\vcs twmc\n\vo)muiie\n\vn,istns\n\vd n has\n\v(6u0 t ",
		},
	}
	for _, tt := range tests {
		t.Run("Code "+tt.name, func(t *testing.T) {
			if got := Code(tt.name); got != tt.want {
				t.Errorf("Code() = %v, want %v", got, tt.want)
			}
		})
		t.Run("Decode "+tt.name, func(t *testing.T) {
			if got := Decode(tt.want); got != tt.name {
				t.Errorf("Code() = %v, want %v", got, tt.name)
			}
		})
	}
}

