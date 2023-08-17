package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// generate tests case for file fetch.go
func TestFetchImage(t *testing.T) {
	testCases := []struct {
		name     string
		url      string
		expected []byte
	}{
		{
			name:     "fetch image",
			url:      "https://avatars.githubusercontent.com/u/5659229?v=4",
			expected: []byte{},
		},
	}

	// run test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := fetchImage(tc.url)
			assert.NoError(t, err)
			assert.NotNil(t, actual)
		})
	}
}
