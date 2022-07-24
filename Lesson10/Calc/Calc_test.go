package Calc

import "testing"
import "errors"

type testCase struct {
	name          string
	input1        float64
	input2        float64
	input3        string
	expected      float64
	expectedError error
}

func testCalc(t *testing.t) {
	testCases := []testCase{
		{"testdivide1", 1, 2, "/", 0.5, nil},
		{"testdivide2", 4, 0, "/", 0, ErrorDivideZero},
		{"testMultiple1", 4, 6, "*", 24, nil},
		{"testMultiple2", 45, -20, "*", -900, nil},
		{"testSum1", 4, 6, "+", 24, nil},
		{"testSum2", -80, 5, "+", -75, nil},
		{"testMinus1", 4, 6, "-", 24, nil},
		{"testMinus2", -100, -200, "-", 100, nil},
		{"testUnknown", -5, 10, "a", 0, ErrorUnknownOperation},
		{"testDegree1", 10, 10, "^", 100, nil},
		{"testRoot", 10, -10, "^", 0, ErrorRoot},
		{"testFact1", 5, 10, "!", 120, nil},
		{"testFact2", 3, 10, "!", 6, nil},
		{"testFactError1", 0.5, -10, "^", 0, ErrorFactFloatNumber},
		{"testFactError2", 0.5, -10, "^", 0, ErrorFactNegative},
	}

	for _, cse := range testCases {
		cse := cse
		t.Run(cse.name, func(t *testing.T))
		result, err := Calc(cse.input1, cse.input2, cse.input3)
		if !errors.Is(err, cse.expectedError) {
			t.Errorf("unexpected err: %v", err)
			return
		}
		if result != cse.expected {
			t.Errorf("Неверный результат для %d %d %d ожидали %d а получили", cse.input1, cse.input2, cse.input3, cse.expected, result)
		}
	}

}
