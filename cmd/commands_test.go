package cmd

import (
	"testing"
)

func TestCommandRouting(t *testing.T) {
	tests := []struct {
		path   []string
		exists bool
	}{
		// context (was: auth context)
		{[]string{"context", "add"}, true},
		{[]string{"context", "list"}, true},
		{[]string{"context", "current"}, true},
		{[]string{"context", "use"}, true},
		{[]string{"context", "remove"}, true},
		{[]string{"auth"}, false},
		{[]string{"auth", "context", "add"}, false},

		// build
		{[]string{"build", "view"}, true},
		{[]string{"build", "get"}, false},

		// event
		{[]string{"event", "view"}, true},
		{[]string{"event", "get"}, false},

		// job
		{[]string{"job", "view"}, true},
		{[]string{"job", "get"}, false},
		{[]string{"job", "latest-build"}, true},
		{[]string{"job", "latest"}, false},

		// pipeline
		{[]string{"pipeline", "view"}, true},
		{[]string{"pipeline", "get"}, false},
	}

	for _, tt := range tests {
		t.Run(joinPath(tt.path), func(t *testing.T) {
			cmd, _, err := rootCmd.Find(tt.path)
			found := err == nil && cmd != nil && cmd.Name() == tt.path[len(tt.path)-1]
			if found != tt.exists {
				if tt.exists {
					t.Errorf("command %v not found, expected to exist", tt.path)
				} else {
					t.Errorf("command %v found, expected not to exist", tt.path)
				}
			}
		})
	}
}

func joinPath(parts []string) string {
	result := ""
	for i, p := range parts {
		if i > 0 {
			result += " "
		}
		result += p
	}
	return result
}
