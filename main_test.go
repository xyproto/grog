package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/xyproto/autoimport"
)

func TestFix(t *testing.T) {
	tests := []struct {
		name         string
		filename     string
		expectedFile string
	}{
		{
			name:         "Kotlin Example1",
			filename:     "testdata/Example1.kt",
			expectedFile: "testdata/ExpectedExample1.kt",
		},
		{
			name:         "Kotlin Example2",
			filename:     "testdata/Example2.kt",
			expectedFile: "testdata/ExpectedExample2.kt",
		},
		{
			name:         "Java Example",
			filename:     "testdata/Example.java",
			expectedFile: "testdata/ExpectedExample.java",
		},
		// Add more test cases if needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expectedData, err := os.ReadFile(test.expectedFile)
			if err != nil {
				t.Fatalf("Failed to read expected file: %s, error: %v", test.expectedFile, err)
			}

			output, err := autoimport.Fix(test.filename, true, true, false)
			if err != nil {
				t.Errorf("fix returned an error for %s: %v", test.filename, err)
			}

			if string(bytes.TrimSpace(output)) != string(bytes.TrimSpace(expectedData)) {
				t.Errorf("fix output for %s does not match expected output.\nExpected:\n%s\nGot:\n%s", test.filename, string(expectedData), string(output))
			}
		})
	}
}
