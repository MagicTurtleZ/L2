package main

import (
	"bufio"
	"regexp"
	"strings"
	"testing"
)

func TestCountMatch(t *testing.T) {
	testTable := []struct {
		regPattern string
		input      []string
		invertFlag bool
		expected   int
	}{
		{
			regPattern: "foo",
			input:      []string{"foo", "bar", "baz", "foo", "foo"},
			expected:   3,
		},
		{
			regPattern: "foo",
			input:      []string{"foo", "bar", "baz", "foo", "foo"},
			invertFlag: true,
			expected:   2,
		},
	}

	for _, testCase := range testTable {
			invertFlag = testCase.invertFlag
			reg := regexp.MustCompile(testCase.regPattern)
			scanner := bufio.NewScanner(strings.NewReader(strings.Join(testCase.input, "\n")))
			count := countMatch(reg, scanner)
			if count != testCase.expected {
				t.Errorf("expected: %d, got: %d", testCase.expected, count)
			}
	}
}

func TestC(t *testing.T) {
	testTable := []struct {
		expectedErr string
	}{
		{
			expectedErr: "canno`t open file: open : The system cannot find the file specified.",
		},
	}

	for _, testCase := range testTable {
		err := customGrep()

		if err.Error() != testCase.expectedErr {
			t.Errorf("expected: %v, got: %v", testCase.expectedErr, err)
		}
	}
}
