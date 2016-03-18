package queue

// Element 元素
type Element struct {
	item interface{}
	next *Element
}

// Queue 队列
type Queue struct {
	first, last *Element
	size        uint
}

// Len Return the stack's length
func (q *Queue) Len() uint {
	return q.size
}

//IsEmpty 栈是否为空
func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

// EnQueue 队尾添加元素
func (q *Queue) EnQueue(item interface{}) {
	oldlast := q.last
	q.last = &Element{item, nil}
	if q.IsEmpty() {
		q.first = q.last
	} else {
		oldlast.next = q.last
	}
	q.size++
}

// DeQueue 队首取出元素
func (q *Queue) DeQueue() (item interface{}) {
	if q.IsEmpty() {
		item = nil
	} else {
		item = q.first.item
		q.first = q.first.next
		q.size--
	}
	return item
}
