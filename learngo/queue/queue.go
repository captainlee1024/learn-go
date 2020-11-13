package queue

// Queue int 类型队列
// type Queue []int

// Queue 支持任何类型的队列
type Queue []interface{}

// Push 添加任何类型元素至队列
func (q *Queue) Push(value interface{})  {
	*q = append(*q, value)
	// 如果我们没有在定义方法和队列时限定类型，又想确保只有 int 类型的元素，可以在函数里判断
	// 这里是一个运行时检查，不是编译时检查
	// *q = append(*q, value.(int))
}

// PushInt 只能添加 int 类型
func (q *Queue) PushInt(v int)  {
	*q = append(*q, v)
}

// Popo 删除任何类型，并返回删除元素
func (q *Queue) Popo() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// PopoInt 删除 int 并返回 int
func (q *Queue) PopoInt() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

// IsEmpty 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}