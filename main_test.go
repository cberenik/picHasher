package main

import (
	"fmt"
	"testing"
)

func Test_isImage(t *testing.T) {
	testCases := []struct {
		in     string
		expect bool
	}{
		{
			in:     "test.jpg",
			expect: true,
		},
		{
			in:     "test.Jpg",
			expect: true,
		},
		{
			in:     "test.png",
			expect: true,
		},
		{
			in:     "test.JPEG",
			expect: true,
		},
		{
			in:     "test.doc",
			expect: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.in), func(t *testing.T) {
			result := isImage(tc.in)
			if result != tc.expect {
				t.Errorf("Expected %t, got %t", tc.expect, result)
			}
		})
	}
}
