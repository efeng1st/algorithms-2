package stack

//参考自： https://gist.github.com/bemasher/1777766

//Stack 栈结构
type Stack struct {
	top  *Element
	size uint
}

// Element 元素
type Element struct {
	item interface{}
	next *Element
}

// Len Return the stack's length
func (s *Stack) Len() uint {
	return s.size
}

//IsEmpty 栈是否为空
func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

// Push a new element onto the stack
func (s *Stack) Push(item interface{}) {
	s.top = &Element{item, s.top}
	s.size++
}

// Pop Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *Stack) Pop() (item interface{}) {
	if s.size > 0 {
		item, s.top = s.top.item, s.top.next
		s.size--
		return
	}
	return nil
}

//---------------

// ArrayStack 能够动态调整数组大小的栈实现
type ArrayStack struct {
	data []interface{}
	size int
}

// NewArrayStack 生成一个数组栈，初始大小为n
func NewArrayStack(n int) *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, n),
		size: 0,
	}
}

// Len Return the stack's length
func (as *ArrayStack) Len() int {
	return as.size
}

//IsEmpty 栈是否为空
func (as *ArrayStack) IsEmpty() bool {
	return as.Len() == 0
}

// 重新调整数组大小
func (as *ArrayStack) resize(max int) {
	temp := make([]interface{}, as.size, max)
	copy(temp, as.data)
	as.data = temp
}

//Push 入栈
func (as *ArrayStack) Push(item interface{}) {
	if as.size == cap(as.data) {
		as.resize(as.size * 2)
	}
	as.size++
	as.data[as.size] = item
}

//Pop 出栈
func (as *ArrayStack) Pop() (item interface{}) {
	item = as.data[as.size-1]
	as.size--

	if (as.size > 0) && (as.size == cap(as.data)/4) {
		as.resize(as.size / 2)
	}

	return item
}
