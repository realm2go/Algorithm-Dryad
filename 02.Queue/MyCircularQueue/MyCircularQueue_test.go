package MyCircularQueue

import (
	"fmt"
	"testing"
)

func TestMyCicularQueue(t *testing.T) {

	p1 := newImpl1(8)

	p1.enQueue(1)
	p1.enQueue(2)
	p1.enQueue(3)
	p1.enQueue(4)
	p1.enQueue(5)
	p1.enQueue(6)
	p1.enQueue(7)

	p1.enQueue(8)

	fmt.Println("当前队列长度:",p1.used)

	fmt.Println("队列是否满了:",p1.isFull())

	fmt.Println(p1.deQueue())
	fmt.Println(p1.deQueue())
	fmt.Println(p1.deQueue())
	fmt.Println(p1.deQueue())
	fmt.Println(p1.deQueue())
	fmt.Println(p1.deQueue())
	fmt.Println(p1.deQueue())
	fmt.Println(p1.deQueue())

	fmt.Println(p1.isFull(),len(p1.a))

	fmt.Println(p1.deQueue())
	fmt.Println(p1.deQueue())
	fmt.Println(p1.deQueue())
	fmt.Println(p1.deQueue())

	fmt.Println(p1.isEmpty())

}