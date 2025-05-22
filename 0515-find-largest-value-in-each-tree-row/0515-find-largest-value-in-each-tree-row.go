/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
    if root == nil {
    return []int{}
}
    var result []int
    queue:=[]*TreeNode{root}

    for len(queue)>0{
    lvlSize:=len(queue)
    max:=math.MinInt64

    for i:=0;i<lvlSize;i++{
        node:=queue[0]
        queue=queue[1:]

        if node.Val>max{
            max=node.Val
        }
        if node.Left!=nil{
            queue=append(queue,node.Left)
        }
        if node.Right!=nil{
            queue=append(queue,node.Right)
        }

    }
    result=append(result,max)
    }
    return result
}