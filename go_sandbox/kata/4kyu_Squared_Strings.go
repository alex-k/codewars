package kata

import (
	"math"
	"strings"
)
const FILL_CHAR uint8 = 11

func Code(s string) string {
	l:=square_size(len(s))
	m:=array(l)
	for i:=range(m){
		for j:=range(m) {
			k:=i*l+j
			if k<len(s) {
				m[j][l-1-i]=s[k]
			} else {
				m[j][l-1-i] = FILL_CHAR
			}
		}
	}
	return format(m, "\n")
}
func Decode(s string) string {
	t:=strings.Split(s,"\n");
	l:=len(t)
	m:=array(l)
	for i:=range(m) {
		for j:=range (m) {
			m[l-1-j][i]=t[i][j]
		}
	}
	ret:=strings.TrimRight(format(m,""),string(FILL_CHAR))
	return ret
}

func array(l int) [][]uint8 {
	m:=make([][]uint8,l,l)
	for i:=range(m){
		m[i]=make([]uint8,l,l)
	}
	return m
}

func format(m [][]uint8, sep string) string{
	l:=len(m)
	r:=make([]string,l,l)
	for i:=range(r){
		r[i] = string(m[i])
	}

	_= r;
	return strings.Join(r,sep);
}

func square_size(n int) int {
	return int(math.Ceil(math.Sqrt(float64(n))))
}
