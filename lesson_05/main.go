package main

import "fmt"

func fibonacci(number int) (int, error) {
	if number < 0 {
		return 0, fmt.Errorf("number cannot be <0 : %d", number)
	}

	cache := make(map[int]int)

	return calculate(number, &cache), nil
}

func calculate(number int, cache *map[int]int) int {
	if number <= 1 {
		return number
	}

	cachedValue, hasValue := (*cache)[number]

	if !hasValue {
		cachedValue = calculate(number-1, cache) + calculate(number-2, cache)
		(*cache)[number] = cachedValue
	}
	return cachedValue
}

func main() {
	fmt.Println(fibonacci(50))
	fmt.Println(fibonacci(-1))
}
