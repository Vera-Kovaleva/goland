package main

import "fmt"

func fibonacci(number int) (int, error) {

	if number < 0 {
		return 0, fmt.Errorf("number cannot be negative: %d", number)
	}

	cache := make(map[int]int)

	return calcImpl(number, &cache), nil
}

func calcImpl(number int, cache *map[int]int) int {
	if number <= 1 {
		return number
	}

	cachedValue, hasValue := (*cache)[number]
	if !hasValue {
		cachedValue = calcImpl(number-1, cache) + calcImpl(number-2, cache)
		(*cache)[number] = cachedValue
	}

	return cachedValue
}

func main() {
	fmt.Println(fibonacci(0))
	fmt.Println(fibonacci(2))
	fmt.Println(fibonacci(9))
	fmt.Println(fibonacci(50))
	fmt.Println(fibonacci(-1))
}
