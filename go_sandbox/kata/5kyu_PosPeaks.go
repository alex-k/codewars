package kata

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(n []int) (ret PosPeaks) {
	ret = PosPeaks{[]int{},[]int{}}
	for i := 1; i < len(n)-1; i++ {
		switch {
		case n[i-1] >= n[i]:
			continue
		case n[i+1] > n[i]:
			continue
		case n[i+1] < n[i]:
			ret.Pos = append(ret.Pos, i)
			ret.Peaks = append(ret.Peaks, n[i])
		case n[i+1] == n[i]:
			loop:
			for j:=i+1;j<len(n);j++ {
				switch  {
				case n[j]>n[i]:
					break loop
				case n[j]<n[i]:
					ret.Pos = append(ret.Pos, i)
					ret.Peaks = append(ret.Peaks, n[i])
					break loop
				}
			}
		}
	}
	return
}
