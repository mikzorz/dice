package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
)

var errInvalidSides = fmt.Errorf("sides must be >= 2")
var errInvalidAmount = fmt.Errorf("amount of dice must be > 0")

func roll(sides int) (int, error) {
	if sides < 2 {
		return 0, errInvalidSides
	}
	return rand.Intn(sides) + 1, nil
}

func rollMany(amount, sides int) ([]int, error) {
	if amount < 1 {
		return nil, errInvalidAmount
	}
	results := []int{}
	for i := 0; i < amount; i++ {
		r, err := roll(sides)
		if err != nil {
			return nil, err
		}
		results = append(results, r)
	}
	return results, nil
}

func main() {
	var amount = flag.Int("n", 1, "amount of dice to roll")
	var sides = flag.Int("d", 6, "dice have this many sides")
	flag.Parse()

	result, err := rollMany(*amount, *sides)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)
}
