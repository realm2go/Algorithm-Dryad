package main

import (
    "fmt"
    "reflect"
)

type myTreeWithNext struct {
    value int
    left *myTreeWithNext
    right *myTreeWithNext
    next *myTreeWithNext
}


func newMyTree(value int) *myTreeWithNext {
    return &myTreeWithNext{
        value: value,
        left:  nil,
        right: nil,
        next: nil,
    }
}

func (a *myTreeWithNext) IsEmpty() bool {
    return reflect.DeepEqual(a, &myTreeWithNext{})
}

func transformAndModifyNextPtr(root *myTreeWithNext) *myTreeWithNext {

    var Q *myTreeWithNext

    if  !root.IsEmpty(){
        Q = root
    }

    for  !Q.IsEmpty() {
         //下一层节点的节点的前一个节点(节点左边的节点)
        // 不能使用 var nextLevelPreNode *myTreeWithNext,否则对象没有初始化
        nextLevelPreNode := &myTreeWithNext{}
        // 下一层结点的头结点
        nextLevelHead := &myTreeWithNext{}
        // Q是当前层的head

        p := Q
        // [当前层处理]按照顺序依次遍历当前层
        for p != nil {
            //# 拿到下一层的结点
            if p.left != nil {
                //# 设置下一层的头
                if nextLevelHead != nil{
                    nextLevelHead = p.left
                }
                //# 前下一层的前驱结点的next指向下一层的结点
                if nextLevelPreNode != nil{

                    nextLevelPreNode.next = p.left
                    nextLevelPreNode = p.left
                }
            }

            if p.right != nil {
                if nextLevelHead != nil {
                    nextLevelHead = p.right
                }
                if nextLevelPreNode != nil {
                    nextLevelPreNode.next = p.right
                    nextLevelPreNode = p.right

                }
            }

            // 下一个元素，从左指向右
            p = p.next

        }

        //# 指向下一层的头
        Q = nextLevelHead

    }

    return root
}

func main() {
    root := newMyTree(3)
    root.left = newMyTree(8)
    root.right = newMyTree(9)

    root.right.left = newMyTree(6)
    root.right.right = newMyTree(7)
    newRoot := transformAndModifyNextPtr(root)

    fmt.Println(root.left)
    fmt.Println(newRoot.left)
    fmt.Println(root.left.next.value)
}