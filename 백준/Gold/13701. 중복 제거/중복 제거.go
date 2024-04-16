package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// word size
//
// 64비트일 때: (^uint(0) >> 63) == 1
// w = 64 = 32 << 1
//
// 32비트일 때: (^uint(0) >> 63) == 0
// w = 32 = 32 << 0
const w = 32 << (^uint(0) >> 63)

type bitarray struct {
	arr []uint // word array
}

// idx의 워드에서의 비트 위치를 가져와 그 비트를 1로 설정한 값을 반환
func mask(idx int) uint { return uint(1 << (idx % w)) }

func (arr *bitarray) check(idx int) bool {
	mask := mask(idx)

	word := arr.arr[idx/w]
	return word&mask != 0
}

func (arr *bitarray) set(idx int, b bool) {
	mask := mask(idx)
	if b {
		arr.arr[idx/w] |= mask
	} else {
		arr.arr[idx/w] &= ^mask
	}
}

func (arr *bitarray) grow(words int) {
	if words <= len(arr.arr) {
		return
	}

	newarr := make([]uint, words)

	// for idx, word := range arr.arr {
	// 	newarr[idx] = word
	// }
	copy(newarr, arr.arr) // == 위 for문

	arr.arr = newarr
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	barr := new(bitarray)
	barr.grow(((1 << 25) / w) + 1) // 2 words

	var n int
	for {
		_, err := fmt.Fscan(in, &n)
		if err == io.EOF {
			break
		}

		if barr.check(n) {
			continue
		}

		barr.set(n, true)

		out.WriteString(strconv.Itoa(n))
		out.WriteByte(' ')
	}
}
