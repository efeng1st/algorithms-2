package main

import (
	"fmt"

	"github.com/hebl/algorithms/UF"
	"github.com/hebl/algorithms/stdlib"
)

type weightwquickUnion struct {
	count int   //分量的数目
	id    []int //分量的Id
	sz    []int
}

//newUF
func newUF(n int) UF.UF {
	uf := weightwquickUnion{
		count: n,
		id:    make([]int, n),
		sz:    make([]int, n),
	}
	for k := range uf.id {
		uf.id[k] = k
	}
	for k := range uf.sz {
		uf.sz[k] = 1
	}
	return &uf
}

// Find 找出根节点
func (wqu *weightwquickUnion) Find(p int) int {
	for p != wqu.id[p] {
		p = wqu.id[p]
	}
	return p
}

// Union 在p和q之间添加一条连线
func (wqu *weightwquickUnion) Union(p, q int) {
	pRoot := wqu.Find(p)
	qRoot := wqu.Find(q)
	if pRoot == qRoot {
		return
	}

	// 将小树的根节点连接到大树的跟节点
	if wqu.sz[pRoot] < wqu.sz[qRoot] {
		wqu.id[pRoot] = qRoot
		wqu.sz[qRoot] += wqu.sz[pRoot]
	} else {
		wqu.id[qRoot] = pRoot
		wqu.sz[pRoot] += wqu.sz[qRoot]
	}

	wqu.count--
}

func (wqu *weightwquickUnion) Connected(p, q int) bool {
	return wqu.id[p] == wqu.id[q]
}

func (wqu *weightwquickUnion) Count() int {
	return wqu.count
}

//---------
func main() {
	var N int
	stdin := stdlib.NewStdIn()
	if stdin.Scan() {
		N, _ = stdin.ReadInt()
	}

	wqu := newUF(N)

	var p, q int
	for stdin.Scan() {
		p, _ = stdin.ReadInt()
		stdin.Scan()
		q, _ = stdin.ReadInt()

		if wqu.Connected(p, q) {
			continue
		}
		wqu.Union(p, q)
		fmt.Println(p, q)
	}
	fmt.Println(wqu.Count())
}
