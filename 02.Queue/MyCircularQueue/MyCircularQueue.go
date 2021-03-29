package MyCircularQueue

import "fmt"

// CircularQueue
type CircularQueue interface {
	enQueue(value int) bool
	deQueue() (bool,int)
	Front() int
	Rear() int
	isEmpty() bool
	isFull() bool
}

// Impl1 实现CircularQueue
type Impl1 struct {
	k int
	used int
	front int
	rear int
	//capacity int
	a []int
}

func newImpl1(k int) *Impl1{
	return &Impl1{
		k:     k,
		used:  0,
		front: 0,
		rear:  0,
		a:   make([]int,k),
	}
}

//
func (CQ *Impl1) enQueue(value int) bool{
	// 如果满了,直接返回错误
	if CQ.isFull(){
		return false
	}

	//fmt.Printf("rear:%d\n",CQ.rear)

	// 如果没有放满，用cap[rear]来存放新的原始
	CQ.a[CQ.rear] = value

	// 为了考虑rear已经是最后一个位置的时候
	CQ.rear = (CQ.rear +1) % CQ.k

	// 已经使用的空间
	CQ.used++

	// 存放成功
	return true
}

func (CQ *Impl1) deQueue() (bool,int) {
	// 如果队列是空，直接返回错误
	if CQ.isEmpty() {
		return false, -1
	}

	//fmt.Printf("front:%d\n",CQ.front)

	// 取出第一个原始
	ret := CQ.a[CQ.front]

	// 已使用的空间
	CQ.used--

	CQ.front = (CQ.front + 1) % CQ.k

	// 取出元素
	return true,ret



	panic("implement me")
}

func (CQ *Impl1) Front() (bool,int) {

	// 如果为空，不能取出首元素
	if CQ.isEmpty(){
		return false, -1
	}
	CQ.used--

	// 取出元素
	return true,CQ.a[CQ.front]

}

func (CQ *Impl1) Rear() (bool,int) {
	// 如果为空，不能取出
	if CQ.isEmpty(){
		return false, -1
	}
	CQ.used--
	// 取出末尾元素
	return true, CQ.a[CQ.rear]
}

func (CQ *Impl1) isEmpty() bool {
	if CQ.used == 0{
		return true
	}

	return false

}

func (CQ *Impl1) isFull() bool {
	if CQ.k == CQ.used {
		return true
	}

	fmt.Printf("队列:%d\n",CQ.used)

	return false
}






