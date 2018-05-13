package practice

import (
	"fmt"
	"strconv"
)

type BinaryTree struct {
	value int
	left  *BinaryTree
	right *BinaryTree
}

func contains(root, a *BinaryTree) bool {
	if root == a {
		return true
	}
	if root == nil {
		return false
	}
	return contains(root.right, a) || contains(root.left, a)
}

func CommonAncestor(root, a, b *BinaryTree) *BinaryTree {
	if !contains(root, a) || !contains(root, b) {
		return nil
	}
	return commonAncestorRec(root, a, b)
}

func commonAncestorRec(root, a, b *BinaryTree) *BinaryTree {
	if contains(root.left, a) && contains(root.left, b) {
		return commonAncestorRec(root.left, a, b)
	}
	if contains(root.right, a) && contains(root.right, b) {
		return commonAncestorRec(root.right, a, b)
	}
	return root
}

// Prove array is constant time append
// integral(2^-x) > 0 for all values of x
// integral(2^-x) = 2^-x/ln(x)
// 2^-x = (e^ln(2))^x = e^(ln(2)*x)
// u = ln(2)* x
// integral e^u du/ln(2)
// 1/ln(2) * integral(e^u)
// 1/ln(2) * e^u +c
// e^(ln(2)*x)/ln2 + c
// 2^-x/ln(2)

func isomorphic(a, b *BinaryTree) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.value != b.value {
		return false
	}
	return isomorphic(a.right, b.right) && isomorphic(a.left, b.left)
}

func isSubtree(root, smaller *BinaryTree) bool {
	if smaller == nil {
		return true
	}
	if root == nil {
		return false
	}
	if isomorphic(root, smaller) {
		return true
	}
	return isSubtree(root.left, smaller) || isSubtree(root.right, smaller)
}

// Assume all values are greater than zero
func pathSumEquals(root *BinaryTree, n int) {
	if n == 0 {
		fmt.Println("")
		return
	}
	pathSumEqualsRec(root, n, make([]int, 0), "")
}

func pathSumEqualsRec(root *BinaryTree, n int, prefixSum []int, path string) {
	if root == nil {
		return
	}

	last := prefixSum[len(prefixSum)-1]
	sum := last + root.value
	currentPath := path + strconv.Itoa(root.value)
	nextPrefixSum := make([]int, len(prefixSum)+1)
	copy(nextPrefixSum, prefixSum)
	nextPrefixSum = append(nextPrefixSum, sum)
	for i, v := range nextPrefixSum {
		if sum-v == n {
			fmt.Println(currentPath[i:])
		}
	}
	pathSumEqualsRec(root.left, n, nextPrefixSum, currentPath)
	pathSumEqualsRec(root.right, n, nextPrefixSum, currentPath)
}

func pathSum2(root *BinaryTree, n int, values []int) {
	if root == nil {
		return
	}
	values = append(values, root.value)
	backwardsSum := 0
	for i := len(values) - 1; i >= 0; i-- {
		backwardsSum += values[i]
		if backwardsSum == n {
			fmt.Println(values[i:])
		}
	}
	pathSum2(root.left, n, values)
	pathSum2(root.right, n, values)
	values = values[:len(values)-1]
}
