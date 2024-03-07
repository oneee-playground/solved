import (
    "strings"
    "strconv"
    "sort"
    "time"
)

type elem struct {
    t time.Time
    isEnd bool
}

func solution(book_time [][]string) int {
    times := make([]elem, 0, len(book_time)*2)
    for _, t := range book_time {
        start, end := parse(t)
        times = append(times, elem{start, false})
        times = append(times, elem{end.Add(10 * time.Minute), true})
    }
    sort.Slice(times, func(i, j int) bool {
        t1, t2 := times[i], times[j]
        if t1.t.Equal(t2.t) {
            return t1.isEnd
        }
        return t1.t.Before(t2.t)
    })
    
    cur, mx := 0, 0
    for _, e := range times {
        if e.isEnd {
            cur--
        } else {
             cur++ 
        }
        if cur > mx {
            mx = cur
        }
    }
    return mx
}

func parse(strs []string) (start, end time.Time) {
    return parseTime(strs[0]), parseTime(strs[1])
}

func parseTime(s string) time.Time {
    split := strings.Split(s, ":")
    hour, _ := strconv.Atoi(split[0])
    min, _ := strconv.Atoi(split[1])
    return time.Date(2024, 3, 7, hour, min, 0, 0, time.Local)
}