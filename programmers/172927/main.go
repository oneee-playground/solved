package main

import (
	"sort"
)

const (
	DIAMOND = 25
	IRON    = 5
	STONE   = 1
)

func solution(picks []int, minerals []string) int {

	var (
		batches  [][]string
		sequence [3]int
	)

	sequence = [3]int{
		DIAMOND, IRON, STONE,
	}

	batches = batchesFrom(minerals, sumAll(picks)*5)

	sort.Slice(batches, func(i, j int) bool {
		return calculateCost(STONE, batches[i]) > calculateCost(STONE, batches[j])
	})

	return simulateBatches(sequence, picks, batches)
}

func batchesFrom(minerals []string, limit int) (batches [][]string) {
	var now, next int

	if limit > len(minerals) {
		limit = len(minerals)
	}

	for next < limit {
		now = next
		next = now + 5

		if next > limit {
			next = limit
		}

		batches = append(batches, minerals[now:next])
	}

	return batches
}

func simulateBatches(sequence [3]int, picks []int, batches [][]string) (sum int) {

	var seqIdx, batchIdx int

	for seqIdx < len(sequence) && batchIdx < len(batches) {
		if picks[seqIdx] == 0 {
			seqIdx++
			continue
		}
		sum += calculateCost(sequence[seqIdx], batches[batchIdx])
		picks[seqIdx]--
		batchIdx++
	}

	return sum
}

func calculateCost(division int, batch []string) (sum int) {

	for _, data := range batch {
		val := findValue(data) / division
		if val == 0 {
			val = 1
		}
		sum += val
	}

	return sum
}

func findValue(mineral string) (val int) {
	switch mineral {
	case "diamond":
		return DIAMOND
	case "iron":
		return IRON
	case "stone":
		return STONE
	}
	return -1
}

func sumAll(arr []int) (sum int) {
	for _, data := range arr {
		sum += data
	}
	return sum
}
