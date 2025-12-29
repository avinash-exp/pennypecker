package tests

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/avinash-exp/pennypecker/provider"
)

const expectedVersion = "0.1.0-alpha.1"

func TestVersionFlag(t *testing.T) {
	tests := []struct {
		name     string
		flag     string
		expected string
	}{
		{
			name:     "long flag --version",
			flag:     "--version",
			expected: "pennypecker version",
		},
		{
			name:     "short flag -v",
			flag:     "-v",
			expected: "pennypecker version",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Build the binary first
			cmd := exec.Command("go", "build", "-o", "../pennypecker_test", "..")
			if err := cmd.Run(); err != nil {
				t.Fatalf("failed to build binary: %v", err)
			}
			defer os.Remove("../pennypecker_test")

			// Run with version flag
			cmd = exec.Command("../pennypecker_test", tt.flag)
			var out bytes.Buffer
			cmd.Stdout = &out

			if err := cmd.Run(); err != nil {
				t.Fatalf("failed to run binary with %s: %v", tt.flag, err)
			}

			output := out.String()
			if !strings.Contains(output, tt.expected) {
				t.Errorf("expected output to contain %q, got %q", tt.expected, output)
			}

			if !strings.Contains(output, expectedVersion) {
				t.Errorf("expected output to contain version %q, got %q", expectedVersion, output)
			}
		})
	}
}

func TestProviderNew(t *testing.T) {
	p := provider.New()

	if p == nil {
		t.Fatal("provider should not be nil")
	}

	if p.Schema == nil {
		t.Error("provider schema should not be nil")
	}

	if p.ResourcesMap == nil {
		t.Error("provider resources map should not be nil")
	}

	if p.DataSourcesMap == nil {
		t.Error("provider data sources map should not be nil")
	}
}
