package QueueExam

import (
	"testing"
)

func TestMyTreeWithNext(t *testing.T) {

	root := newMyTree(3)
	root.left = newMyTree(8)
	root.right = newMyTree(9)

	root.right.left = newMyTree(6)
	root.right.right = newMyTree(7)
	//newRoot := transformAndModifyNextPtr(root)
	//fmt.Println(newRoot.left.next.value)
}