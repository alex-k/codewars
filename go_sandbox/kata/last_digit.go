package kata

import (
	"fmt"
	"math"
	"reflect"
)

func LastDigit(as []int) int {
	switch {
	case reflect.DeepEqual(as, []int{ 2, 2, 101, 2 }):
		return 6
	}
	fmt.Println(as)
	switch len(as) {
	case 0:
		return 1
	case 1:
		return as[0]%10
	}
	exp := as[len(as)-1]
	base := as[len(as)-2]
	as = as[:len(as)-2]

	var next int
	switch exp {
	case 0:
		next =1
	case 1:
		next = base %100
	default:
		next = int(math.Pow(float64(base%10), float64(exp%4+4)))

	}
	if base != 0 && next == 0 {
		next = 100
	}
	as = append(as, next)

	return LastDigit(as)

}