package main

import (
	"reflect"
	"testing"
)

func TestStringValidate(t *testing.T) {
	testTable := []struct {
		inputStr 	string
		dFlag 		string
		fFlag  		fieldsFlag
		sFlag		bool
		expected	[]string
		expPermit	bool
	} {
		{
			inputStr: "Selecting bytes 1-6 yields the first two characters",
			dFlag: " ",
			fFlag: []string{"1,2"},
			sFlag: false,
			expected: []string{"Selecting", "bytes"},
			expPermit: true,
		},
		{
			inputStr: "Selecting bytes 1-6 yields the first two characters",
			dFlag: ":",
			fFlag: []string{"1,2"},
			sFlag: false,
			expected: []string{"Selecting bytes 1-6 yields the first two characters"},
			expPermit: true,
		},
		{
			inputStr: "Selecting bytes 1-6 yields the first two characters",
			dFlag: ":",
			fFlag: []string{"1,2"},
			sFlag: true,
			expected: nil,
			expPermit: false,
		},
		{
			inputStr: ":Selecting:all:three:characters:with:the:-c:switch:doesn’t:work.",
			dFlag: ":",
			fFlag: []string{"1,2"},
			sFlag: false,
			expected: []string{"", "Selecting"},
			expPermit: true,
		},
	}

	for _, testCase := range testTable {
		s = testCase.sFlag
		d = testCase.dFlag
		f = testCase.fFlag

		result, permit := stringValidateForTest(testCase.inputStr)
		if permit != testCase.expPermit {
			t.Errorf("expected: %v, got: %v", testCase.expPermit, permit)
		}

		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("expected: %v, got: %v", testCase.expected, result)
		}
	}
}

var content = []string{"LC_ALL=ja_JP.UTF-8 cut -c1-3 kanji.utf-8.txt",
":Selecting:all:three:characters:with:the:-c:switch:doesn’t:work.",
"In this case, an illegal UTF-8 string is produced.},",
}

func TestCustomCut(t *testing.T) {
	testTable := []struct {
		dFlag 		string
		fFlag  		fieldsFlag
		sFlag		bool
		expected	[]string
	} {
		{
			dFlag: " ",
			fFlag: []string{"1,2"},
			sFlag: true,
			expected: []string{
				"LC_ALL=ja_JP.UTF-8", "cut",
				"In", "this",
			},
		},
		{
			dFlag: ":",
			fFlag: []string{"1,2"},
			sFlag: true,
			expected: []string{
				"" ,"Selecting",
			},
		},
		{
			dFlag: ":",
			fFlag: []string{"1,2"},
			sFlag: false,
			expected: []string{
				"LC_ALL=ja_JP.UTF-8 cut -c1-3 kanji.utf-8.txt",
				"" ,"Selecting",
				"In this case, an illegal UTF-8 string is produced.},",
			},
		},
	}

	for _, testCase := range testTable {
		s = testCase.sFlag
		d = testCase.dFlag
		f = testCase.fFlag

		result := customCutForTest(content)

		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("expected: %v, got: %v", testCase.expected, result)
		}
	}
} 