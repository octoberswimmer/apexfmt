package formatter

import (
	"strings"
	"testing"
)

func Test_soql_formatter_set_source(t *testing.T) {
	f := NewSOQLFormatter()

	testQuery := "SELECT Id FROM Account"
	f.SetSource(testQuery)

	// Format should use the set source
	err := f.Format()
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	formatted, err := f.Formatted()
	if err != nil {
		t.Errorf("Expected no error from Formatted(), got: %v", err)
	}

	// Should format a simple query
	if !strings.Contains(formatted, "SELECT") {
		t.Errorf("Expected formatted output to contain SELECT, got: %s", formatted)
	}
}

func Test_soql_formatter_set_filename_for_error_reporting(t *testing.T) {
	f := NewSOQLFormatter()
	f.SetFilename("query.soql")

	// Test with valid SOQL - should not error
	f.SetSource("SELECT Id FROM Account")

	err := f.Format()
	if err != nil {
		t.Errorf("Expected no error for valid SOQL, got: %v", err)
	}
}

func Test_soql_formatter_returns_error_on_invalid_syntax(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		filename string
		wantErr  bool
	}{
		{
			name:     "Valid simple SOQL",
			input:    "SELECT Id FROM Account",
			filename: "test.soql",
			wantErr:  false,
		},
		{
			name:     "Valid SOQL with WHERE",
			input:    "SELECT Id, Name FROM Account WHERE Name = 'Test'",
			filename: "test.soql",
			wantErr:  false,
		},
		{
			name:     "Valid complex SOQL",
			input:    "SELECT Id, Name FROM Account WHERE CreatedDate = TODAY ORDER BY Name",
			filename: "complex.soql",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewSOQLFormatter()
			f.SetFilename(tt.filename)
			f.SetSource(tt.input)

			err := f.Format()

			if tt.wantErr {
				if err == nil {
					t.Error("Expected error for invalid SOQL, got nil")
				} else if tt.filename != "" && !strings.Contains(err.Error(), tt.filename) {
					t.Errorf("Expected error to contain filename '%s', got: %s", tt.filename, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error for valid SOQL, got: %v", err)
				}
			}
		})
	}
}

func Test_soql_formatter_formatted_method(t *testing.T) {
	f := NewSOQLFormatter()

	// Calling Formatted() before Format() should trigger formatting
	f.SetSource("SELECT Id, Name FROM Account WHERE CreatedDate = TODAY")

	formatted, err := f.Formatted()
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Should be properly formatted
	if !strings.Contains(formatted, "SELECT\n") {
		t.Error("Expected formatted output to have SELECT on its own line")
	}
	if !strings.Contains(formatted, "FROM\n") {
		t.Error("Expected formatted output to have FROM on its own line")
	}
	if !strings.Contains(formatted, "WHERE\n") {
		t.Error("Expected formatted output to have WHERE on its own line")
	}
}

func Test_soql_formatter_methods_with_valid_input(t *testing.T) {
	f := NewSOQLFormatter()
	f.SetFilename("complex.soql")

	// Test all methods with valid SOQL
	f.SetSource("SELECT Id, Name FROM Account WHERE Name = 'Test'")

	err := f.Format()
	if err != nil {
		t.Errorf("Expected no error for valid SOQL, got: %v", err)
		return
	}

	// Test Formatted method
	formatted, err := f.Formatted()
	if err != nil {
		t.Errorf("Expected no error from Formatted(), got: %v", err)
		return
	}

	// Should be properly formatted
	if !strings.Contains(formatted, "SELECT") {
		t.Error("Expected formatted output to contain SELECT")
	}
}
