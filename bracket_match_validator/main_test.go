package main

import (
	"io"
	"strings"
	"testing"
)

func TestIsValidSequence(t *testing.T) {
	testCases := []struct {
		bSeq    string
		isValid bool
	}{
		{"{([])}", true},
		{"[({}]", false},
		{"[()", false},
		{")[]", false},
		{"", true},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := isValidSequence(tc.bSeq)
			if result != tc.isValid {
				t.Errorf("got %t for sequence '%s'", result, tc.bSeq)
			}
		})
	}
}

func TestGetBracketSequence(t *testing.T) {
	fooReader := strings.NewReader("Foo\nBar\n")

	testCases := []struct {
		reader    io.Reader
		shouldErr bool
		expected  string
	}{
		{fooReader, false, "Foo"},
		{strings.NewReader(""), true, ""},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result, err := getBracketSequence(tc.reader)
			if tc.shouldErr && err == nil {
				t.Errorf("expected err, got nil")
			}
			if !tc.shouldErr && err != nil {
				t.Error(err)
			}
			if result != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, result)
			}
		})
	}
}
