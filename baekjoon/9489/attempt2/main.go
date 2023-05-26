package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		nodeCnt   uint32
		targetNum uint32

		inputList   []uint32
		splitted    [][]uint32
		generations [][]uint32
	)

	for {
		fmt.Scanf("%d %d", &nodeCnt, &targetNum)

		if nodeCnt == 0 && targetNum == 0 {
			return
		}

		inputList = getInput(nodeCnt)
		splitted = split(inputList)
		generations = getGenerations(inputList, splitted)

		genHeight := getGenHeight(targetNum, generations) // 자식의 height 찾기

		if genHeight < 2 {
			fmt.Println(0)
			continue
		}
		parent, brothers := findParentAndList(targetNum, genHeight-1, inputList, splitted, generations)
		count := findSiblingCount(parent, brothers, inputList, splitted)

		fmt.Println(count)
	}
}

func getInput(nodeCnt uint32) []uint32 {

	r := bufio.NewReader(os.Stdin)
	inputList := make([]uint32, nodeCnt)

	for i := 0; i < int(nodeCnt); i++ {
		fmt.Fscan(r, &inputList[i])
	}
	return inputList
}

func split(src []uint32) [][]uint32 {

	var (
		splitted [][]uint32
		temp     []uint32
	)

	for idx, data := range src {
		temp = append(temp, data)
		if len(src) <= idx+1 || src[idx+1]-data > 1 {
			splitted = append(splitted, temp)
			temp = []uint32{}
		}
	}

	return splitted
}

func getGenerations(originList []uint32, splitted [][]uint32) [][]uint32 {

	var generations [][]uint32

	capacity := 1
	totalLen := 0

	for cursor := 0; totalLen < len(originList) && cursor < len(splitted); {
		tempList := []uint32{}

		for i := 0; i < capacity && cursor < len(splitted); i++ {
			totalLen += len(splitted[cursor])
			tempList = append(tempList, splitted[cursor]...)
			cursor++
		}
		generations = append(generations, tempList)
		capacity = len(generations[len(generations)-1])
	}

	return generations
}

func getGenHeight(targetNum uint32, generations [][]uint32) int {
	// could optimize it
	for idx, list := range generations {
		for _, data := range list {
			if data == targetNum {
				return idx
			}
		}
	}

	return -1
}

func findParentAndList(targetNum uint32, targetGen int, originList []uint32, splitted [][]uint32, generations [][]uint32) (uint32, []uint32) {

	for _, target := range generations[targetGen] {
		for idx, data := range originList {
			if idx+1 == len(splitted) {
				break
			}
			if data == target && contains(targetNum, splitted[idx+1]) {
				for _, list := range splitted {
					if contains(data, list) {
						return data, list
					}
				}
			}
		}
	}
	return 0, nil
}

func findSiblingCount(ignoreNode uint32, list []uint32, originList []uint32, splitted [][]uint32) int {
	cnt := 0
	for _, node := range list {
		if node != ignoreNode {
			idx := findIndex(node, originList)

			if idx+1 < len(splitted) {
				cnt += len(splitted[idx+1])
			}
		}
	}

	return cnt
}

func contains(targetNum uint32, list []uint32) bool {
	for _, data := range list {
		if data == targetNum {
			return true
		}
	}
	return false
}

func findIndex(target uint32, list []uint32) int {
	for idx, data := range list {
		if data == target {
			return idx
		}
	}
	return -1
}
