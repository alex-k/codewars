package kata

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func MovingShift(s string, shift int) []string {
	return split(encode(s, shift),5)
}

func DemovingShift(arr []string, shift int) string {
	return decode(strings.Join(arr,""),shift)
}

func encode(s string, shift int) (ret string) {
	for i, a := range s {
		if unicode.IsLetter(a) {
			switch {
			case a >= 'A' && a <= 'Z':
				a += int32((i + shift) % 26)
				if a > 'Z' {
					a -= 26
				}
			case a >= 'a' && a <= 'z':
				a += int32((i + shift) % 26)
				if a > 'z' {
					a -= 26
				}
			}
		}
		ret += fmt.Sprintf("%c", a)
	}
	return
}

func decode(s string, shift int) (ret string) {
	for i, a := range s {
		if unicode.IsLetter(a) {
			switch {
			case a >= 'A' && a <= 'Z':
				a -= int32((i + shift) % 26)
				if a < 'A' {
					a += 26
				}
			case a >= 'a' && a <= 'z':
				a -= int32((i + shift) % 26)
				if a < 'a' {
					a += 26
				}
			}
		}
		ret += fmt.Sprintf("%c", a)
	}
	return
}

func split(s string, p int) (ret []string) {
	ret = make([]string,p)
	d:= int(math.Ceil(float64(len(s))/float64(p)))
	for i:=0; i<p; i++ {
		x:=int64(math.Min(float64(i*d),float64(len(s))))
		y:=int64(math.Min(float64((i+1)*d),float64(len(s))))
		ret[i] = s[x:y]
	}
	return
}

