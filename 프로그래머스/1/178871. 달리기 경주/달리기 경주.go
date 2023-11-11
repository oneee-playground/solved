func solution(players []string, callings []string) []string {
    hashmap := make(map[string]int, len(players))
    for idx, player := range players {
        hashmap[player] = idx
    }
    
    for _, calling := range callings {
        ahead := hashmap[calling]
        behind := hashmap[players[ahead-1]]
                
        hashmap[players[ahead]] = behind
        hashmap[players[behind]] = ahead
        
        players[ahead], players[behind] = players[behind], players[ahead]
    }
    
    return players
}