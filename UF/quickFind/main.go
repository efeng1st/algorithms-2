package main

import (
	"fmt"

	"github.com/hebl/algorithms/UF"
	"github.com/hebl/algorithms/stdlib"
)

type quickFind struct {
	count int   //分量的数目
	id    []int //分量的Id
}

//newUF
func newUF(n int) UF.UF {
	uf := quickFind{
		count: n,
		id:    make([]int, n),
	}
	for k := range uf.id {
		uf.id[k] = k
	}
	return &uf
}

func (qf *quickFind) Union(p, q int) {
	pID := qf.id[p]
	qID := qf.id[q]

	if pID == qID {
		return
	}
	for i := 0; i < len(qf.id); i++ {
		if qf.id[i] == pID {
			qf.id[i] = qID
		}
	}
	qf.count--
}

func (qf *quickFind) Find(p int) int {
	return qf.id[p]
}

func (qf *quickFind) Connected(p, q int) bool {
	return qf.id[p] == qf.id[q]
}

func (qf *quickFind) Count() int {
	return qf.count
}

func main() {
	var N int
	stdin := stdlib.NewStdIn()
	if stdin.Scan() {
		N, _ = stdin.ReadInt()
	}

	qf := newUF(N)

	var p, q int
	for stdin.Scan() {
		p, _ = stdin.ReadInt()
		stdin.Scan()
		q, _ = stdin.ReadInt()

		if qf.Connected(p, q) {
			continue
		}
		qf.Union(p, q)
		fmt.Println(p, q)
	}
	fmt.Println(qf.Count())
}
