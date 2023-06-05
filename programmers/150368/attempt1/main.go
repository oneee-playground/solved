package main

import (
	"log"
)

func solution(users [][]int, emoticons []int) []int {

	var (
		maxEmotePlus   int
		maxEmoteSales  int
		emoteDiscounts []int
	)

	discountRate := [4]int{
		10, 20, 30, 40,
	}

	maxEmotePlus, maxEmoteSales, emoteDiscounts = findMax(users, emoticons, discountRate)
	log.Println(maxEmotePlus, emoteDiscounts)

	for idx := len(emoticons) - 1; idx >= 0; idx-- {
		for {
			rate := emoteDiscounts[idx]
			if rate == discountRate[0] {
				break
			}
			emoteDiscounts[idx] -= 10
			if newEmotePlus, newEmoteSales := calculate(users, emoticons, emoteDiscounts); newEmotePlus < maxEmotePlus || newEmoteSales < maxEmoteSales {
				log.Println(newEmoteSales, maxEmoteSales)
				emoteDiscounts[idx] += 10
				break
			} else {
				maxEmoteSales = newEmoteSales
			}
		}
		log.Println(emoteDiscounts)
	}

	maxEmotePlus, maxEmoteSales = calculate(users, emoticons, emoteDiscounts)

	return []int{
		maxEmotePlus, maxEmoteSales,
	}
}

func discount(val int, rate int) int {
	return val - (val * rate / 100)
}

func fill(with int, len int) []int {
	arr := make([]int, len)
	for i := range arr {
		arr[i] = with
	}
	return arr
}

func findMax(users [][]int, emoticons []int, discountRate [4]int) (emotePlusCnt, emoteSales int, emoteDiscounts []int) {

	var (
		rateIdx int = len(discountRate) - 1
	)

	emoteDiscounts = make([]int, len(emoticons))
	tempDiscounts := fill(40, len(emoticons))

	for rateIdx >= 0 {
		for idx := range emoticons {
			tempDiscounts[idx] = discountRate[rateIdx]
			newEmotePlusCnt, newEmoteSales := calculate(users, emoticons, tempDiscounts)

			if newEmotePlusCnt > emotePlusCnt {
				emotePlusCnt = newEmotePlusCnt
				emoteSales = newEmoteSales
				copy(emoteDiscounts, tempDiscounts)
			}
		}
		rateIdx--
	}

	return
}

func calculate(users [][]int, emoticons []int, discountRates []int) (maxEmotePlus, maxEmoteSales int) {

	var (
		salePrice, sumPrice int
		budgetExceeded      bool
	)

	for _, user := range users {

		for idx, price := range emoticons {
			if rate := discountRates[idx]; rate >= user[0] {
				salePrice = discount(price, rate)
				sumPrice += salePrice
				if exchange := sumPrice - user[1]; exchange >= 0 {
					maxEmotePlus++
					budgetExceeded = true
					break
				}
			}
		}

		if !budgetExceeded {
			maxEmoteSales += sumPrice
		}

		sumPrice = 0
		budgetExceeded = false
	}

	return
}
