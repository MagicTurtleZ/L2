package main

import "testing"

func TestUnpuck(t *testing.T) {
	testTable := []struct {
		str 	 	string
		expectedStr string
		expectedErr string
	} {
		{
			str: "a4bc2d5e",
			expectedStr: "aaaabccddddde",
		},
		{
			str: "abcd",
			expectedStr: "abcd",
		}, 
		{
			str: "45",
			expectedStr: "",
			expectedErr: incorrectString,
		}, 
		{
			str: "",
			expectedStr: "",
		}, 
		{
			str: `qwe\4\5`,
			expectedStr: `qwe45`,
		},
		{
			str: `qwe\45`,
			expectedStr: `qwe44444`,
		},
		{
			str: `qwe\\5`,
			expectedStr: `qwe\\\\\`,
		},
	}

	for _, testCase := range testTable {
		result, err := Unpack(testCase.str)

		t.Logf("Calling Unpack(%v), result %s %v", testCase.str, result, err)
		if err != nil {
			if err.Error() != testCase.expectedErr {
				t.Errorf("Incorrect result. Expect %s, got %s", testCase.expectedErr, err.Error())
			}
		}

		if result != testCase.expectedStr {
			t.Errorf("Incorrect result. Expect %s, got %s", testCase.expectedStr, result)
		}
	}
}
