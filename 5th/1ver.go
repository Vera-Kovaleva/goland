package main

import "fmt"

func fibonacci2(number int) (int, error) {
	if number < 0 {
		return 0, fmt.Errorf("number cannot be <0 : %d", number)
	}

	cache2 := make(map[int]int)

	return calculate(number, &cache2), nil
}

func calculate(number int, cache2 *map[int]int) int {
	if number <= 1 {
		return number
	}

	cachedValue, hasValue := (*cache2)[number]

	if !hasValue {
		cachedValue = calculate(number-1, cache2) + calculate(number-2, cache2)
		(*cache2)[number] = cachedValue
	}
	return cachedValue
}

func main() {
	fmt.Println(fibonacci2(50))
}
