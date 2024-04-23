import (
    "bytes"
    "strconv"
)

func solution(food []int) string {
    buf := bytes.NewBuffer(nil)
    for i := 1; i < len(food); i++ {
        buf.Write(bytes.Repeat([]byte(strconv.Itoa(i)), food[i]/2))
    }
    buf.WriteByte('0')
    
    b := make([]byte, buf.Len()-1)
    copy(b, buf.Bytes()[:buf.Len()-1])
    reverse(b)
    
    buf.Write(b)
    return buf.String()
}

func reverse(r []byte) {
    for i := 0; i < len(r)/2; i++ {
        r[i], r[len(r)-i-1] = r[len(r)-i-1], r[i]
    }
}