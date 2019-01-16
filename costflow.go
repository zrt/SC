package SC

type Networkflow struct {
	S    int
	T    int
	H    []int
	X    []int
	P    []int
	flow []int
	cost []float64
	from []int
	tot  int
	pnum int
}

func NewCostFlow(pnum int) Networkflow {
	pnum += 2
	enum := pnum * 10
	nw := Networkflow{pnum - 1, pnum, make([]int, pnum+1), make([]int, enum), make([]int, enum), make([]int, enum), make([]float64, enum), make([]int, enum), 1, pnum}
	return nw
}

func (nw *Networkflow) AddEdge(x, y, z int, c float64) {
	nw.tot++
	nw.P[nw.tot] = y
	nw.X[nw.tot] = nw.H[x]
	nw.H[x] = nw.tot
	nw.flow[nw.tot] = z
	nw.cost[nw.tot] = c
	nw.from[nw.tot] = x
	nw.tot++
	nw.P[nw.tot] = x
	nw.X[nw.tot] = nw.H[y]
	nw.H[y] = nw.tot
	nw.flow[nw.tot] = 0
	nw.cost[nw.tot] = -c
	nw.from[nw.tot] = y
	if nw.tot > len(nw.X) {
		panic("slice so small")
	}
}

func (nw *Networkflow) Check(x, y int) bool {
	for i := nw.H[x]; i != 0; i = nw.X[i] {
		if nw.P[i] == y {
			if nw.flow[i] == 0 {
				return true
			} else {
				return false
			}
		}
	}
	panic("not found")
}

func (nw *Networkflow) spfa(num int) (bool, int) {
	//println("in spfa()")
	d := make([]float64, nw.pnum+1)
	a := make([]int, nw.pnum+1)
	p := make([]int, nw.pnum+1)
	inq := make([]bool, nw.pnum+1)
	q := make([]int, nw.pnum)
	h := 0
	t := 0
	//println("init spfa")
	for i := 1; i <= nw.pnum; i++ {
		d[i] = 1e20
	}
	d[nw.S] = 0
	a[nw.S] = 1000000 // 每次流1000000
	q[t] = nw.S
	t = (t + 1) % nw.pnum
	inq[nw.S] = true
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	//println("new spfa")
	for h != t {
		x := q[h]
		//println(x, d[x])
		h = (h + 1) % nw.pnum
		inq[x] = false
		for i := nw.H[x]; i != 0; i = nw.X[i] {
			if nw.flow[i] > 0 && d[nw.P[i]] > d[x]+nw.cost[i]+1e-8 {
				d[nw.P[i]] = d[x] + nw.cost[i]
				a[nw.P[i]] = min(a[x], nw.flow[i])
				p[nw.P[i]] = i
				if !inq[nw.P[i]] {
					inq[nw.P[i]] = true
					q[t] = nw.P[i]
					t = (t + 1) % nw.pnum
				}
			}
		}
	}
	//println(d[nw.T])
	if d[nw.T] >= 1e19 {
		return false, 0
	}
	a[nw.T] = min(a[nw.T], num)
	x := nw.T
	for x != nw.S {
		//println(x)
		nw.flow[p[x]] -= a[nw.T]
		nw.flow[p[x]^1] += a[nw.T]
		x = nw.from[p[x]]
	}
	return true, a[nw.T]
}

func (nw *Networkflow) MincostMaxflow(num int) {

	for i := 0; i < num; {
		println("start", i, "/", num)
		ok, f := nw.spfa(num - i)
		if !ok {
			panic("spfa failed")
		}
		i += f
		println("done", i, "/", num)
	}
}
