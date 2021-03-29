package stackExam

import (
	"github.com/golang-collections/collections/stack"
)


func Soluction(fishSize []int,fishDirection []int) int{

	// 获取鱼的数量
	fishNum := len(fishSize)
	if fishNum <= 1 {
		return fishNum
	}

	// 0 表示左移，1 表示右移
	left  := 0
	right := 1

	//初始化一个栈
	t := stack.New()
	for i := 0; i < fishNum ; i++  {
		// 当前鱼的情况：游动方向和大小
		curFishDirection := fishDirection[i]
		curFishSize := fishSize[i]

		// 当前的鱼是否被栈中的鱼吃掉
		hasEat := false
		// 如果栈中有鱼，并且栈中鱼向右，当前的鱼向左游，那么就会有相遇的可能性
		for t.Len()>0 && fishDirection[t.Peek().(int)] == right && curFishDirection == left{
			// 如果栈顶的鱼比较大，那么把新来的吃掉
			if fishSize[t.Peek().(int)] > curFishSize {
				hasEat = true
				break
			}
			// 如果栈中的鱼较小，那么会把栈中的鱼吃掉，栈中的鱼被消除，所以需要弹栈。
			t.Pop()
		}
		// 如果新来的鱼，没有被吃掉，那么压入栈中。
		if !hasEat {
			t.Push(i)
		}
	}
	return t.Len()
}
