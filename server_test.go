package main

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// generate test cases for file server.go
func TestNewServer(t *testing.T) {
	testCases := []struct {
		name     string
		options  ServerOptions
		expected bool
	}{
		{
			name:     "new server",
			options:  ServerOptions{Path: "/test", Handler: func(*fiber.Ctx) error { return nil }},
			expected: true,
		},
	}

	// run test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			app := NewServer(tc.options)
			assert.Equal(t, tc.expected, app != nil)
		})
	}
}
