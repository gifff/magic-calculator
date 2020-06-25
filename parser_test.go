package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckRequiredArgs(t *testing.T) {
	origCommands := commands
	commands = map[string]Command{
		"test": Command{
			RequiredArgs: 5,
		},
	}

	testCases := []struct {
		name     string
		args     []string
		expected bool
	}{
		{
			name:     "args is less than required",
			args:     []string{"1", "3", "5"},
			expected: false,
		},
		{
			name:     "args is more than required",
			args:     []string{"1", "3", "5", "1", "3", "5", "5"},
			expected: true,
		},
		{
			name:     "args is less than required",
			args:     []string{"1", "3", "5", "5", "5"},
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := checkRequiredArgs("test", tc.args)

			assert.Equal(t, tc.expected, got)
		})
	}

	commands = origCommands
}
