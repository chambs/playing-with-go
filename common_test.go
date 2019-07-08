package main

import "testing"

func TestSum(t *testing.T) {

	cases := []Test{
		{int32(sum(5, 5)), 10},
		{int32(sub(5, 5)), 10},
	}

	for index, testCase := range cases {
		if testCase.expected != testCase.result {
			t.Errorf("Got %d but expected %d", testCase.result, testCase.expected)
		} else {
			t.Logf("Success %d", index)
		}
	}

}
