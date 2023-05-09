package main

import "testing"

func TestRollReturnsValidResult(t *testing.T) {
	t.Run("D6", func(t *testing.T) {
		rollDie1000(t, 6)
	})

	t.Run("D20", func(t *testing.T) {
		rollDie1000(t, 20)
	})

	t.Run("invalid die", func(t *testing.T) {
		for _, n := range []int{1, 0, -1} {
			_, err := roll(n)
			if err == nil {
				t.Fail()
			}
		}
	})
}

func TestRollManyDice(t *testing.T) {
	testCases := []struct {
		amount int
		sides  int
		err    error
	}{
		{5, 6, nil},
		{1, 20, nil},
		{999, 4, nil},
		{0, 2, errInvalidAmount},
		{-1, 6, errInvalidAmount},
		{5, 1, errInvalidSides},
		{5, 0, errInvalidSides},
		{5, -1, errInvalidSides},
	}

	for _, tt := range testCases {
		results, err := rollMany(tt.amount, tt.sides)
		if err != tt.err {
			t.Errorf("wrong error. want %s, got %s", err, tt.err)
			continue
		}
		if tt.err == nil && len(results) != tt.amount {
			t.Errorf("did not receive the right amount of dice roll results, want %d, got %d", tt.amount, len(results))
		}
	}
}

func rollDie1000(t *testing.T, sides int) {
	t.Helper()
	results := map[int]int{}
	for i := 0; i < 1000; i++ {
		result, _ := roll(sides)
		if result < 1 || result > sides {
			t.Fatalf("result (%d) not within range of valid values", result)
		}
		results[result]++
	}
	for i := 1; i <= sides; i++ {
		if _, ok := results[i]; !ok {
			t.Errorf("%d was never rolled", i)
		}
	}
}
