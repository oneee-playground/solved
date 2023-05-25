package main

import (
	"fmt"
)

// stack implementation
type stack []*tree

func (s *stack) push(node *tree) {
	*s = append(*s, node)
}

func (s *stack) pop() {
	*s = (*s)[:len(*s)-1]
}

func (s *stack) peek() *tree {
	return (*s)[s.height()-1]
}

func (s *stack) height() int {
	return len(*s)
}

// tree implementation
type tree struct {
	data  uint32
	nodes []*tree
}

func main() {
	var (
		nodeCnt   uint32
		targetNum uint32

		myTree *tree
	)

	for {
		fmt.Scanf("%d %d", &nodeCnt, &targetNum)

		if nodeCnt == 0 && targetNum == 0 {
			return
		}

		myTree = createTreeWithInput(nodeCnt)
		fmt.Println(findSiblingCount(myTree, targetNum))
	}
}

func createTreeWithInput(nodeCnt uint32) *tree {
	var myTree *tree
	myStack := new(stack)

	var (
		input    uint32
		prevData uint32
	)

	for i := 0; i < int(nodeCnt); i++ {
		fmt.Scan(&input)

		if i == 0 {
			// create root
			myTree = new(tree)
			myTree.data = input
			prevData = input + 1
			myStack.push(myTree)
			continue
		}

		insertData(input, prevData, myStack)
		prevData = input
	}

	return myTree
}

func insertData(data, prevData uint32, myStack *stack) {
	curNode := myStack.peek()

	myTree := new(tree)
	myTree.data = data

	if data-prevData == 1 {
		curNode.nodes = append(curNode.nodes, myTree)
		return
	}

	newCursor := findNextNode(curNode, myStack)
	newCursor.nodes = append(newCursor.nodes, myTree)
	myStack.push(newCursor)

}

func findNextNode(curNode *tree, myStack *stack) *tree {

	i := 0

	for {
		if myStack.height() != 1 {
			myStack.pop()
		}
		parent := myStack.peek()
		curIdx := findIndex(parent.nodes, curNode)

		if curIdx == len(parent.nodes)-1 && myStack.height() != 1 {
			curNode = parent
			myStack.pop()
			parent = myStack.peek()
			curIdx = findIndex(parent.nodes, curNode)

			i++
		}
		curNode = parent.nodes[curIdx+1]

		for j := 0; j < i; j++ {
			myStack.push(curNode)
			curNode = curNode.nodes[0]
		}

		if len(curNode.nodes) == 0 {
			return curNode
		}
	}

}

func findIndex(nodes []*tree, target *tree) int {
	for idx, node := range nodes {
		if node == target {
			return idx
		}
	}
	return -1
}

func findSiblingCount(myTree *tree, target uint32) int {
	myStack := new(stack)
	targetParent, height := findParentAndHeight(myTree, myStack, target)

	myStack = new(stack)

	return countSiblingForNode(myTree, targetParent, myStack, height)
}

func findParentAndHeight(myTree *tree, myStack *stack, targetNum uint32) (*tree, int) {

	myStack.push(myTree)
	defer myStack.pop()

	for _, node := range myTree.nodes {
		if node.data == targetNum {
			return myTree, myStack.height()
		}

		t, height := findParentAndHeight(node, myStack, targetNum)
		if t != nil {
			return t, height
		}
	}

	return nil, -1
}

func countSiblingForNode(curNode *tree, ignoreNode *tree, myStack *stack, targetHeight int) int {
	myStack.push(curNode)
	defer myStack.pop()

	cnt := 0

	if myStack.height()+1 == targetHeight {
		for _, node := range curNode.nodes {
			if node != ignoreNode {
				cnt += len(node.nodes)
			}
		}
	}

	for _, node := range curNode.nodes {
		cnt += countSiblingForNode(node, ignoreNode, myStack, targetHeight)
	}

	return cnt
}
