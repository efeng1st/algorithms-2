package sort

const (
	//SELECT 选择排序
	SELECT uint8 = 1 << iota
	//INSERT 插入排序
	INSERT
	//SHELL 希尔排序
	SHELL
)

//FuncAPI 排序函数
type FuncAPI func(data Interface)

//Sort sort func
func Sort(data Interface, algo uint8) {
	sorter := map[uint8]FuncAPI{
		SELECT: selectSort,
		INSERT: insertSort,
		SHELL:  shellSort,
	}
	sorter[algo](data)
}

//selectSort 选择排序
func selectSort(data Interface) {
	N := data.Len()
	for i := 0; i < N; i++ {
		min := i
		for j := i + 1; j < N; j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		data.Swap(i, min)
	}
}

func insertSort(data Interface) {
	N := data.Len()
	for i := 0; i < N; i++ {
		for j := i; j > 0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

func shellSort(data Interface) {
	N := data.Len()
	h := 1
	for h < N/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < N; i++ {
			for j := i; j >= h && data.Less(j, j-h); j -= h {
				data.Swap(j, j-h)
			}
		}
		h = h / 3
	}
}
