func solution(numbers []int, target int) int {
    cnt := add(numbers[1:], numbers[0], target)
    cnt += add(numbers[1:], -numbers[0], target)
    return cnt
}

func add(nums []int, cur, target int) int {
    if len(nums) == 0 {
        if cur == target {
            return 1
        }
        return 0
    }
    cnt := add(nums[1:], cur+nums[0], target)
    cnt += add(nums[1:], cur-nums[0], target)
    return cnt
}