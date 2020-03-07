package main

import (
	"fmt"
	"testing"
)

type TestFileData struct {
	name  string
	isDir bool
}

func (tfd *TestFileData) Name() string {
	return tfd.name
}

func (tfd *TestFileData) IsDir() bool {
	return tfd.isDir
}

func Test_isImage(t *testing.T) {
	testCases := []struct {
		in     FileData
		expect bool
	}{
		{
			in:     &TestFileData{name: "test.jpg", isDir: false},
			expect: true,
		},
		{
			in:     &TestFileData{name: "test.Jpg", isDir: false},
			expect: true,
		},
		{
			in:     &TestFileData{name: "test.png", isDir: false},
			expect: true,
		},
		{
			in:     &TestFileData{name: "test.JPEG", isDir: false},
			expect: true,
		},
		{
			in:     &TestFileData{name: "test.doc", isDir: false},
			expect: false,
		},
		{
			in:     &TestFileData{name: "folder", isDir: true},
			expect: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.in), func(t *testing.T) {
			// TODO: make image structs to run these tests
			result := isImage(tc.in)
			if result != tc.expect {
				t.Errorf("Expected %t, got %t", tc.expect, result)
			}
		})
	}
}
