package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCacheImage(t *testing.T) {
	testCases := []struct {
		name     string
		url      string
		content  []byte
		expected bool
	}{
		{
			name:     "cache image",
			url:      "https://avatars.githubusercontent.com/u/5659229?v=4",
			content:  []byte{},
			expected: true,
		},
	}

	// run test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cacheImage(tc.url, tc.content)
			_, ok := getImageCache(tc.url)
			assert.Equal(t, tc.expected, ok)
		})
	}
}

func TestGetImageCache(t *testing.T) {
	testCases := []struct {
		name     string
		url      string
		expected bool
	}{
		{
			name:     "get image cache",
			url:      "https://avatars.githubusercontent.com/u/5659229?v=4",
			expected: true,
		},
	}

	// run test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, ok := getImageCache(tc.url)
			assert.Equal(t, tc.expected, ok)
		})
	}
}

func TestGetImageCacheNotFound(t *testing.T) {
	testCases := []struct {
		name     string
		url      string
		expected bool
	}{
		{
			name:     "get image cache not found",
			url:      "https://avatars.githubusercontent.com/u/5659229?v=5",
			expected: false,
		},
	}

	// run test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, ok := getImageCache(tc.url)
			assert.Equal(t, tc.expected, ok)
		})
	}
}

func TestGetImageCacheEmpty(t *testing.T) {
	testCases := []struct {
		name     string
		url      string
		expected bool
	}{
		{
			name:     "get image cache empty",
			url:      "https://avatars.githubusercontent.com/u/5659229?v=6",
			expected: false,
		},
	}

	// run test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, ok := getImageCache(tc.url)
			assert.Equal(t, tc.expected, ok)
		})
	}
}

func TestGetImageCacheEmptyNotFound(t *testing.T) {
	testCases := []struct {
		name     string
		url      string
		expected bool
	}{
		{
			name:     "get image cache empty not found",
			url:      "https://avatars.githubusercontent.com/u/5659229?v=7",
			expected: false,
		},
	}

	// run test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, ok := getImageCache(tc.url)
			assert.Equal(t, tc.expected, ok)
		})
	}
}
