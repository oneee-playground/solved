import "bytes"

const diff = 'a' - 'A'

func solution(s string) string {
    b := []byte(s)
    barr := bytes.Split(b, []byte(" "))
    for _, word := range barr {
        for idx, c := range word {
            if idx % 2 == 0 {
                if c >= 'a' && c <= 'z' {
                    word[idx] = c - diff
                }
                continue
            }
            if c >= 'A' && c <= 'Z' {
                word[idx] = c + diff
            }
        }
    }
    b = bytes.Join(barr, []byte(" "))
    return string(b)
}