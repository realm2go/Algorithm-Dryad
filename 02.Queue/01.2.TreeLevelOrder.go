package QueueExam

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
	"reflect"
)

type myTree struct {
	value int
	left *myTree
	right *myTree
}


func newMyTree(value int) *myTree {
	return &myTree{
		value: value,
		left:  nil,
		right: nil,
	}
}

func (a myTree) IsEmpty() bool {
	return reflect.DeepEqual(a, &myTree{})
}

func Transform(root *myTree) (result [][]int){

	// 存放每一层的结果
	result = [][]int{}

	// 记录当前层
	var cur []*myTree

	// 记录当前层，注意不要把空的root放进去
	if !root.IsEmpty()   {
		cur = append(cur,root)
	}

	for len(cur) > 0 {
		// 存放下一级别
		nextLevel := make([]*myTree,0)

		// curResult用来存放当前层的值
		curResult := make([]int,0)

		for _,c := range cur{
			// 把当前层的值放到curResult里面
			fmt.Printf("c:%v\n",c.left)
			curResult = append(curResult,c.value)
			// 生成下一层
			if c.left != nil{
				nextLevel = append(nextLevel,c.left)
			}
			if c.right != nil {
				nextLevel = append(nextLevel,c.right)
			}

		}
		// 指向下一层
		cur = nextLevel

		result = append(result,curResult)

	}

	return result
}


func levelOrder(root *myTree) (result [][]int){

	// 初始化一个FIFO队列
	Q := queue.New()

	// 如果结点不为空，那么加入FIFO队列
	if !root.IsEmpty(){
		Q.Enqueue(root)
	}

	// 利用FIFO队列进行层次遍历
	for Q.Len()>0 {
		// 取出当前层里面的元素个数
		qSize := Q.Len()
		fmt.Printf("qSize:%d\n",qSize)
		// 当前层的结果存放于tmp
		var curLevel []int
		// 遍历当前层的每个节点
		for i := 0; i < qSize ; i++  {
			// 出队,把接口类型强制转成自定义的类型
			cur := Q.Dequeue().(*myTree)
			// 把结果存放当于当前层中
			curLevel = append(curLevel,cur.value)
			// 把下一层的结点入队，注意入队时需要非空才可以入队。
			if cur.left != nil {
				Q.Enqueue(cur.left)
			}
			if cur.right != nil {
				Q.Enqueue(cur.right)
			}
		}
		result = append(result,curLevel)

	}

	return result
}

func main() {
	// 初始化一个树

	root := newMyTree(3)
	root.left = newMyTree(8)
	root.right = newMyTree(9)

	root.right.left = newMyTree(6)
	root.right.right = newMyTree(7)

	result2 := Transform(root)

	result1 := levelOrder(root)

	fmt.Printf("result1:%v\n",result1)
	fmt.Printf("result2:%v\n",result2)

}