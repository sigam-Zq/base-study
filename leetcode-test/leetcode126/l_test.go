package leetcode125

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

/*
1022. 从根到叶的二进制数之和
https://leetcode.cn/problems/sum-of-root-to-leaf-binary-numbers/
*/

func TestCreateTree(t *testing.T) {
	printTree(createTree([]int{1, 0, 1, 0, 1, 0, 1}))

	fmt.Println("-----")

	printTree(createTree([]int{1, 1}))
}

func TestBinaryToBase10(t *testing.T) {
	fmt.Println(binaryToBase10([]int{1, 0, 1, 0, 1, 0, 1}))

	fmt.Println("-----")

	fmt.Println(binaryToBase10([]int{1, 1}))
	fmt.Println("-----")

	fmt.Println(binaryToBase10([]int{1, 0}))
}

func TestXxx(t *testing.T) {

	for i, v := range []struct {
		f    func(root *TreeNode) int
		root *TreeNode
		want int
	}{
		{
			f:    sumRootToLeafFix,
			root: createTree([]int{1, 0, 1, 0, 1, 0, 1}),
			want: 22,
		},
		{
			f:    sumRootToLeafFix,
			root: createTree([]int{1, 1}),
			want: 3,
		},
	} {

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := v.f(v.root)
			if got != v.want {
				t.Errorf(" got %v want %v \n", got, v.want)
			}
		})
	}
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, val int) int {
	if node == nil {
		return 0
	}
	val = val<<1 | node.Val
	if node.Left == nil && node.Right == nil {
		return val
	}
	return dfs(node.Left, val) + dfs(node.Right, val)
}

func sumRootToLeafLeetCode(root *TreeNode) int {
	return dfs(root, 0)
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/sum-of-root-to-leaf-binary-numbers/solutions/1521470/cong-gen-dao-xie-de-er-jin-zhi-shu-zhi-h-eqss/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func sumRootToLeafCoderAI(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, pathValue int) int {
		if node == nil {
			return 0
		}
		// 更新当前路径的二进制值
		pathValue = pathValue*2 + node.Val
		// 如果是叶子节点，返回该路径表示的十进制值
		if node.Left == nil && node.Right == nil {
			return pathValue
		}
		// 递归计算左右子树的路径和
		return dfs(node.Left, pathValue) + dfs(node.Right, pathValue)
	}
	return dfs(root, 0)
}

// dfs
func sumRootToLeafFix(root *TreeNode) int {

	var dfs func(n *TreeNode, binaryList []int) int
	dfs = func(n *TreeNode, bList []int) int {
		if n == nil {
			return 0
		}

		bList = append(bList, n.Val)
		if n.Left == nil && n.Right == nil {
			fmt.Println("-", bList)
			fmt.Println("--", binaryToBase10(bList))
			return binaryToBase10(bList)
		}
		return dfs(n.Left, bList) + dfs(n.Right, bList)
	}

	return dfs(root, make([]int, 0))
}

func binaryToBase10(binaryList []int) int {
	if len(binaryList) == 0 {
		return 0
	}
	s := 0
	for i := 0; i < len(binaryList); i++ {
		if binaryList[i] == 1 {
			s += int(math.Pow(float64(2), float64(len(binaryList)-i-1)))
		}
	}
	return s
}

// 当前树如果不是平衡树的情况下这里没用了
func sumRootToLeaf(root *TreeNode) int {
	// 树的高
	h := 0
	for v := root; v != nil; v = v.Left {
		h++
	}
	var fSumNodeVal func(h int, node *TreeNode) int
	fSumNodeVal = func(h int, node *TreeNode) int {
		if node == nil || h == 0 {
			return 0
		}
		// 求出当前节点的最底层的子树的宽
		// 当前树如果不是平衡树的情况下这里没用了
		w := int(math.Pow(float64(2), float64(h-1)))
		var val int
		if node.Val == 1 {
			val = int(math.Pow(float64(2), float64(h-1))) * w
		}
		fmt.Println("val", val)
		return val + fSumNodeVal(h-1, node.Left) + fSumNodeVal(h-1, node.Right)
	}

	return fSumNodeVal(h, root)
}

// TODO
func arrayToTree(list []int) *TreeNode {
	root := &TreeNode{}

	// i, j := 0, 1
	// for j < len(list) {

	// }

	return root
}

// createTree creates a binary tree from the given array representation. Coder
func createTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	// Create root node
	root := &TreeNode{Val: nums[0]}

	// Use a queue to keep track of nodes for level order traversal
	queue := []*TreeNode{root}
	index := 1

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		if index < len(nums) && nums[index] != -1 { // -1 can be used as a placeholder for nil nodes
			currentNode.Left = &TreeNode{Val: nums[index]}
			queue = append(queue, currentNode.Left)
		}
		index++

		if index < len(nums) && nums[index] != -1 {
			currentNode.Right = &TreeNode{Val: nums[index]}
			queue = append(queue, currentNode.Right)
		}
		index++
	}

	return root
}
func inOrderTraversal(root *TreeNode) {
	if root != nil {
		inOrderTraversal(root.Left)
		fmt.Printf("%d ", root.Val)
		inOrderTraversal(root.Right)
	}
}

// printTree prints the binary tree in a structured format.
func printTree(root *TreeNode) {
	if root == nil {
		return
	}

	// Find the depth of the tree to determine the height of the output
	depth := maxDepth(root)
	width := (1 << depth) - 1 // Total number of nodes in the last level, which gives us the width
	output := make([][]string, depth)
	for i := range output {
		output[i] = make([]string, width)
	}

	// Fill the output array with spaces to create a "canvas" for the tree
	for row := 0; row < depth; row++ {
		for col := 0; col < width; col++ {
			output[row][col] = "  " // Two spaces as padding
		}
	}

	// Fill in the actual node values into the canvas, starting from the root at position (0, (width-1)/2)
	fillOutput(output, root, 0, 0, width-1)

	// Print the tree structure
	for _, row := range output {
		for _, cell := range row {
			fmt.Print(cell, " ")
		}
		fmt.Println()
	}
}

// fillOutput recursively fills the output array with node values at appropriate positions.
func fillOutput(output [][]string, root *TreeNode, row, left, right int) {
	if root == nil {
		return
	}

	// Calculate the middle position for the current node in the current row
	middle := (left + right) / 2
	output[row][middle] = fmt.Sprintf("%d", root.Val)

	// Recursively fill left and right subtrees, moving to lower rows
	fillOutput(output, root.Left, row+1, left, middle-1)
	fillOutput(output, root.Right, row+1, middle+1, right)
}

// maxDepth calculates the maximum depth of a binary tree.
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}
