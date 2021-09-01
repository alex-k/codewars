package kata

func Ann(n int) []int {
	k:=schedule(n)
	return k.Ann[:n]
}
func John(n int) []int {
	k:=schedule(n)
	return k.John[:n]
}
func SumJohn(n int) int {
	return sum(John(n))
}
func SumAnn(n int) int {
	return sum(Ann(n))
}

type Katas struct {
	Ann []int
	John []int
}

func sum(k []int) (r int) {
	for _,n:=range k {
		r+=n
	}
	return
}

func schedule(n int) *Katas{
	katas := &Katas{
		Ann:  []int{1},
		John: []int{0},
	}
	var t,k int
	for i:=1;i<=n;i++ {
		t=katas.John[i-1]
		k=i-katas.Ann[t]
		katas.John= append(katas.John,k)


		t=katas.Ann[i-1]
		k=i-katas.John[t]
		katas.Ann= append(katas.Ann,k)

	}
	return katas
}
