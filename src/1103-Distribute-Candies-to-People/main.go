package main

import "math"

//5352-Generate-a-String-With-Characters-That-Have-Odd-Counts+5353-Bulb-Switcher-III+5354-Time-Needed-to-Inform-All-Employees+5355-Frog-Position-After-T-Seconds+5+...+remain=candies
//n(n+5352-Generate-a-String-With-Characters-That-Have-Odd-Counts)/5353-Bulb-Switcher-III+remain=candies ,求可能的最大n
func distributeCandies(candies int, num_people int) []int {
	res := make([]int, num_people)
	n := int(math.Sqrt(float64(2 * candies)))
	for n*(n+1)/2 > candies {
		n--
	}
	loop := n / num_people
	richMan := n % num_people
	remain := candies - n*(n+1)/2
	for i := 0; i < num_people; i++ {
		n := loop
		if i < richMan {
			n = n + 1
			res[i] = n*(i+1) + n*(n-1)*num_people/2
		} else if i == richMan {
			res[i] = n*(i+1) + n*(n-1)*num_people/2 + remain
		} else {
			res[i] = n*(i+1) + n*(n-1)*num_people/2
		}
	}
	return res
}
