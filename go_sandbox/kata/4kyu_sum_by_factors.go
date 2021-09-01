package kata

import (
	"fmt"
)

func primes(n int) (primes []int) {
	b := make([]bool, n)
	for i := 2; i < n; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for k := i; k < n; k += i {
			b[k] = true
		}
	}
	return
}

func SumOfDivided(lst []int) (ret string) {
	max := 0
	for _, d := range lst {
		if d > max {
			max = d
		} else if -d > max {
			max = -d
		}
	}
	primesD := primes(max+1)
	for _, d := range primesD {
		found := 0
		sum := 0
		for _, n := range lst {
			if n%d == 0 {
				found++
				sum += n
			}
		}
		if found > 0 {
			ret += fmt.Sprintf("(%d %d)", d, sum)
		}
	}
	return
}
