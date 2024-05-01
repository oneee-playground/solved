def solution(dartResult):
    arr = [0 for i in range(3)]
    idx = 0
    cur = 0
    stage = 0
    
    while idx < len(dartResult):
        if stage == 0:
            numstr = ""
            while dartResult[idx] >= '0' and dartResult[idx] <= '9':
                numstr += dartResult[idx]
                idx += 1
            arr[cur] = int(numstr) 
            stage += 1
        elif stage == 1:
            c = dartResult[idx]
            if c == 'D':
                arr[cur] **= 2
            if c == 'T':
                arr[cur] **= 3
            idx += 1
            stage += 1
        elif stage == 2:
            c = dartResult[idx]
            if c != '*' and c != '#':
                cur += 1
                stage = 0
                continue
            if c == '*':
                if cur != 0:
                    arr[cur-1] *= 2
                arr[cur] *= 2
            if c == '#':
                arr[cur] *= -1
            idx += 1
            cur += 1
            stage = 0
    
    return sum(arr)