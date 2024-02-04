package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var (
		n int
	)

	fmt.Fscan(r, &n)
	fmt.Fprintln(w, "어느 한 컴퓨터공학과 학생이 유명한 교수님을 찾아가 물었다.")
	recursion(w, 0, n)
}

func recursion(w io.Writer, cur, n int) {
	prefix := strings.Repeat("_", cur*4)
	fmt.Fprintf(w, "%s%s\n", prefix, "\"재귀함수가 뭔가요?\"")
	if cur == n {
		fmt.Fprintf(w, "%s%s\n", prefix, "\"재귀함수는 자기 자신을 호출하는 함수라네\"")
		fmt.Fprintf(w, "%s%s\n", prefix, "라고 답변하였지.")
		return
	}
	fmt.Fprintf(w, "%s%s\n", prefix, "\"잘 들어보게. 옛날옛날 한 산 꼭대기에 이세상 모든 지식을 통달한 선인이 있었어.")
	fmt.Fprintf(w, "%s%s\n", prefix, "마을 사람들은 모두 그 선인에게 수많은 질문을 했고, 모두 지혜롭게 대답해 주었지.")
	fmt.Fprintf(w, "%s%s\n", prefix, "그의 답은 대부분 옳았다고 하네. 그런데 어느 날, 그 선인에게 한 선비가 찾아와서 물었어.\"")
	recursion(w, cur+1, n)
	fmt.Fprintf(w, "%s%s\n", prefix, "라고 답변하였지.")
}
