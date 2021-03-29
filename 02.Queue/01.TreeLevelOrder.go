package QueueExam

import "github.com/golang-collections/collections/queue"

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
       // 当前层的结果存放于tmp
       curLevel := []int{0}
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