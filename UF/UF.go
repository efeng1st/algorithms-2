package UF

/***
Union-Find 算法

- 初始化， 初始N个触点
- Union(p int, q int) 在p和q之间添加一条连线
- Find(p int) p所在分量的标识符
- Connected(p int, q int) 如果p和q在同一分量中则返回 true
- Count 联通分量的数目

***/

//UF 基础组件
type UF interface {
	Union(p, q int)          // 归并两个分量
	Find(p int) int          // 分量标识符
	Connected(p, q int) bool // 分量是否连通
	Count() int              // 所有连通分量的数目
}
