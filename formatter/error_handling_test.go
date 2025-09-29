package formatter

import (
	"strings"
	"testing"
)

func Test_formatter_returns_error_on_invalid_syntax(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr string
	}{
		{
			name:    "Missing semicolon",
			input:   "public class Test { void method() { String s = 'test' } }",
			wantErr: "line",
		},
		{
			name:    "Invalid token",
			input:   "public class { }",
			wantErr: "line",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			// Use empty filename to use the reader
			f := NewFormatter("", reader)
			err := f.Format()

			if err == nil {
				t.Errorf("Expected error for invalid syntax, got nil")
				return
			}

			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("Expected error to contain '%s', got '%s'", tt.wantErr, err.Error())
			}
		})
	}
}

func Test_formatter_no_error_on_valid_syntax(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "Simple valid class",
			input: "public class Test { }",
		},
		{
			name:  "Class with method",
			input: "public class Test { void method() { String s = 'test'; } }",
		},
		{
			name: "Class with SOQL",
			input: `public class Test {
				void method() {
					List<Account> accts = [SELECT Id FROM Account];
				}
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			// Use empty filename to use the reader
			f := NewFormatter("", reader)
			err := f.Format()

			if err != nil {
				t.Errorf("Expected no error for valid syntax, got: %v", err)
			}
		})
	}
}

func Test_errorListener_accumulates_multiple_errors(t *testing.T) {
	listener := &errorListener{filename: "test.cls", suppressStderr: true}

	// Simulate multiple syntax errors
	listener.SyntaxError(nil, nil, 1, 10, "first error", nil)
	listener.SyntaxError(nil, nil, 2, 20, "second error", nil)
	listener.SyntaxError(nil, nil, 3, 30, "third error", nil)

	if !listener.HasErrors() {
		t.Error("Expected HasErrors() to return true")
	}

	err := listener.GetError()
	if err == nil {
		t.Error("Expected GetError() to return an error")
		return
	}

	errStr := err.Error()
	if !strings.Contains(errStr, "first error") {
		t.Error("Expected error to contain 'first error'")
	}
	if !strings.Contains(errStr, "second error") {
		t.Error("Expected error to contain 'second error'")
	}
	if !strings.Contains(errStr, "third error") {
		t.Error("Expected error to contain 'third error'")
	}
}

func Test_errorListener_no_errors_returns_nil(t *testing.T) {
	listener := &errorListener{filename: "test.cls", suppressStderr: true}

	if listener.HasErrors() {
		t.Error("Expected HasErrors() to return false for new listener")
	}

	err := listener.GetError()
	if err != nil {
		t.Errorf("Expected GetError() to return nil, got: %v", err)
	}
}
