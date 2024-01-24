import "strings"

type person struct {
    earns int
    received int
    gave int
    friends map[string]int
}

func solution(friends []string, gifts []string) int {
    friendMap := make(map[string]*person)
    for _, friend := range friends {
        friendMap[friend] = &person{
            friends: make(map[string]int),
        }
    }
    for _, gift := range gifts {
        s := strings.Split(gift, " ")
        giverN, receiverN := s[0], s[1]
        giver, receiver := friendMap[giverN], friendMap[receiverN]
        
        giver.gave++
        receiver.received++
        giver.friends[receiverN]++
    }
    for nameP, person := range friendMap {
        for nameF, friend := range friendMap {
            if nameP == nameF {
                continue
            }
            gave := person.friends[nameF]
            got := friend.friends[nameP]
            if gave == got {
                rateP := person.gave - person.received
                rateF := friend.gave - friend.received
                if rateP > rateF {
                    person.earns++
                }
                continue
            }
            if gave > got {
                person.earns++
            }
        }
    }
    max := 0
    for _, p := range friendMap {
        if p.earns > max {
            max = p.earns
        }
    }
    return max
}