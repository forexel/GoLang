package fib

import "testing"
import "errors"

type testCase struct {
	name          string
	input        float64
	expected      float64
	expectedError error
}

func testfiboptimized(t *testing.T) {
	testCases := []testCase{
		{"Test10", 10, 55, nil},
		{"Test9", 9, 34, nil},
		{"Test1", 2, 1, nil},
		{"Test-1", -2, 0, ErrorNegative},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.name, func(t *testing.T))
		result, err := fibOptimized(cse.input, make(map[int]int))
		if !errors.Is(err, cse.expectedError) {
			t.Errorf("unexpected err: %v", err)
			return
		}
		if result != cse.expected {
			t.Errorf("Неверный результат для %d ожидали %d а получили", cse.input, cse.expected, result)
		}
	}

}
