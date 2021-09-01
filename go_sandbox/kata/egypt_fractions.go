package kata

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Fraction struct {
	m float64
	n float64
}

func ToFract(s string) (*Fraction, error) {
	x := strings.Split(s, "/")
	switch len(x) {
	case 1:
		y := strings.Split(x[0], ".")
		switch len(y) {
		case 1:
			n, err := strconv.ParseInt(y[0], 10, 64)
			if err != nil {
				return nil, err
			}
			if n == 0 {
				return nil, errors.New("Null")
			}
			return &Fraction{float64(n), 1}, nil
		case 2:
			m, err := strconv.ParseInt(y[0]+y[1], 10, 64)
			if err != nil {
				return nil, err
			}
			p := len(y[1])
			n := int64(math.Pow10(p))
			return &Fraction{float64(m), float64(n)}, nil

		default:
			return nil, errors.New("invalid format")
		}
	case 2:
		m, err := strconv.ParseInt(x[0], 10, 64)
		if err != nil {
			return nil, err
		}
		n, err := strconv.ParseInt(x[1], 10, 64)
		if err != nil {
			return nil, err
		}
		if n == 0 {
			return nil, errors.New("Division by zero")
		}
		return &Fraction{float64(m), float64(n)}, nil
	default:
		return nil, errors.New("Invalid format")
	}

}

func Decompose(s string) (ret []string) {
	println(s)
	ret = []string{}

	var f *Fraction
	f, err := ToFract(s)
	if err != nil {
		return []string{}
	}

	rem := int64(f.m / f.n)
	if rem > 0 {
		ret = append(ret, fmt.Sprint(int64(rem)))
		f.m -= float64(rem) * f.n
	}

	if f.m==0 {
		return
	}

	for f.m != 0 {
		x := math.Ceil(f.n / f.m)
		ret = append(ret, "1/"+fmt.Sprint(int64(x)))

		m := (math.Ceil(f.n/f.m)*f.m-f.n)
		n := f.n * x
		f = &Fraction{m, n}
	}
	return
}