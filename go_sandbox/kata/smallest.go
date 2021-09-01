package kata

import (
	"fmt"
	"strconv"
)

func Smallest(n int64) []int64 {
	smallest := []int64{n, 0, 0}
	s := int64(len(fmt.Sprint(n)))

	var i, j int64
	for i = 0; i < s; i++ {
		for j = 0; j < s; j++ {
			x := swap(n, i, j)
			if x < smallest[0] {
				smallest = []int64{x, i, j}
			}
		}
	}
	return smallest
}

func swap(n, x, y int64) int64 {
	s := fmt.Sprint(n)
	r := ""
	switch {
	case x > y:
		r1 := s[:y]
		r2 := s[y:x]
		t := s[x : x+1]
		r3 := s[x+1:]
		r = r1 + t + r2 + r3
	case y > x:
		r1 := s[:x]
		t := s[x : x+1]
		r2 := s[x+1 : y+1]
		r3 := s[y+1:]
		r = r1 + r2 + t + r3
	default:
		return n
	}
	d, _ := strconv.ParseInt(r, 10, 64)
	return d
}
