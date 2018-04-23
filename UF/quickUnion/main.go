package main

import (
	"fmt"

	"github.com/hebl/algorithms/UF"
	"github.com/hebl/algorithms/stdlib"
)

type quickUnion struct {
	count int   //分量的数目
	id    []int //分量的Id
}

//newUF
func newUF(n int) UF.UF {
	uf := quickUnion{
		count: n,
		id:    make([]int, n),
	}
	for k := range uf.id {
		uf.id[k] = k
	}
	return &uf
}

// Find 找出根节点
func (qu *quickUnion) Find(p int) int {
	for p != qu.id[p] {
		p = qu.id[p]
	}
	return p
}

// Union 在p和q之间添加一条连线
func (qu *quickUnion) Union(p, q int) {
	pRoot := qu.Find(p)
	qRoot := qu.Find(q)
	if pRoot == qRoot {
		return
	}

	qu.id[pRoot] = qRoot
	qu.count--
}

func (qu *quickUnion) Connected(p, q int) bool {
	return qu.id[p] == qu.id[q]
}

func (qu *quickUnion) Count() int {
	return qu.count
}

//---------
func main() {
	var N int
	stdin := stdlib.NewStdIn()
	if stdin.Scan() {
		N, _ = stdin.ReadInt()
	}

	qu := newUF(N)

	var p, q int
	for stdin.Scan() {
		p, _ = stdin.ReadInt()
		stdin.Scan()
		q, _ = stdin.ReadInt()

		if qu.Connected(p, q) {
			continue
		}
		qu.Union(p, q)
		fmt.Println(p, q)
	}
	fmt.Println(qu.Count())
}
