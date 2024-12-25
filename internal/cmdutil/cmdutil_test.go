// Copyright (C) 2024 Yota Hamada
// SPDX-License-Identifier: GPL-3.0-or-later

package cmdutil

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitCommandWithQuotes(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		cmd, args, err := SplitCommand("ls -al test/")
		require.NoError(t, err)
		require.Equal(t, "ls", cmd)
		require.Len(t, args, 2)
		require.Equal(t, "-al", args[0])
		require.Equal(t, "test/", args[1])
	})
	t.Run("WithJSON", func(t *testing.T) {
		cmd, args, err := SplitCommand(`echo {"key":"value"}`)
		require.NoError(t, err)
		require.Equal(t, "echo", cmd)
		require.Len(t, args, 1)
		require.Equal(t, `{"key":"value"}`, args[0])
	})
	t.Run("WithQuotedJSON", func(t *testing.T) {
		cmd, args, err := SplitCommand(`echo "{\"key\":\"value\"}"`)
		require.NoError(t, err)
		require.Equal(t, "echo", cmd)
		require.Len(t, args, 1)
		require.Equal(t, `"{\"key\":\"value\"}"`, args[0])
	})
}

func TestSplitCommand(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantCmd   string
		wantArgs  []string
		wantErr   bool
		errorType error
	}{
		{
			name:     "simple command no args",
			input:    "echo",
			wantCmd:  "echo",
			wantArgs: []string{},
		},
		{
			name:     "command with single arg",
			input:    "echo hello",
			wantCmd:  "echo",
			wantArgs: []string{"hello"},
		},
		{
			name:     "command with backtick",
			input:    "echo `echo hello`",
			wantCmd:  "echo",
			wantArgs: []string{"`echo hello`"},
		},
		{
			name:     "command with multiple args",
			input:    "echo hello world",
			wantCmd:  "echo",
			wantArgs: []string{"hello", "world"},
		},
		{
			name:     "command with quoted args",
			input:    `echo "hello world"`,
			wantCmd:  "echo",
			wantArgs: []string{"\"hello world\""},
		},
		{
			name:     "command with pipe",
			input:    "echo foo | grep foo",
			wantCmd:  "echo",
			wantArgs: []string{"foo", "|", "grep", "foo"},
		},
		{
			name:     "complex pipe command",
			input:    "echo foo | grep foo | wc -l",
			wantCmd:  "echo",
			wantArgs: []string{"foo", "|", "grep", "foo", "|", "wc", "-l"},
		},
		{
			name:     "command with quoted pipe",
			input:    `echo "hello|world"`,
			wantCmd:  "echo",
			wantArgs: []string{"\"hello|world\""},
		},
		{
			name:      "empty command",
			input:     "",
			wantErr:   true,
			errorType: ErrCommandIsEmpty,
		},
		{
			name:     "command with escaped quotes",
			input:    `echo "\"hello world\""`,
			wantCmd:  "echo",
			wantArgs: []string{`"\"hello world\""`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCmd, gotArgs, err := SplitCommand(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Errorf("splitCommand() error = nil, want error")
					return
				}
				if tt.errorType != nil && err != tt.errorType {
					t.Errorf("splitCommand() error = %v, want %v", err, tt.errorType)
				}
				return
			}

			if err != nil {
				t.Errorf("splitCommand() error = %v, want nil", err)
				return
			}

			if gotCmd != tt.wantCmd {
				t.Errorf("splitCommand() gotCmd = %v, want %v", gotCmd, tt.wantCmd)
			}

			if len(gotArgs) != len(tt.wantArgs) {
				t.Errorf("splitCommand() gotArgs length = %v, want %v", len(gotArgs), len(tt.wantArgs))
				return
			}

			for i := range gotArgs {
				if gotArgs[i] != tt.wantArgs[i] {
					t.Errorf("splitCommand() gotArgs[%d] = %v, want %v", i, gotArgs[i], tt.wantArgs[i])
				}
			}
		})
	}
}

func TestSplitCommandWithParse(t *testing.T) {
	t.Run("CommandSubstitution", func(t *testing.T) {
		cmd, args, err := SplitCommandWithEval("echo `echo hello`")
		require.NoError(t, err)
		require.Equal(t, "echo", cmd)
		require.Len(t, args, 1)
		require.Equal(t, "hello", args[0])
	})
	t.Run("QuotedCommandSubstitution", func(t *testing.T) {
		cmd, args, err := SplitCommandWithEval("echo `echo \"hello world\"`")
		require.NoError(t, err)
		require.Equal(t, "echo", cmd)
		require.Len(t, args, 1)
		require.Equal(t, "hello world", args[0])
	})
	t.Run("EnvVar", func(t *testing.T) {
		os.Setenv("TEST_ARG", "hello")
		cmd, args, err := SplitCommandWithEval("echo $TEST_ARG")
		require.NoError(t, err)
		require.Equal(t, "echo", cmd)
		require.Len(t, args, 1)
		require.Equal(t, "hello", args[0])
	})
}

func TestSubstituteStringFields(t *testing.T) {
	// Set up test environment variables
	os.Setenv("TEST_VAR", "test_value")
	os.Setenv("NESTED_VAR", "nested_value")
	defer os.Unsetenv("TEST_VAR")
	defer os.Unsetenv("NESTED_VAR")

	type Nested struct {
		NestedField   string
		NestedCommand string
		unexported    string
	}

	type TestStruct struct {
		SimpleField  string
		EnvField     string
		CommandField string
		MultiField   string
		EmptyField   string
		unexported   string
		NestedStruct Nested
	}

	tests := []struct {
		name    string
		input   TestStruct
		want    TestStruct
		wantErr bool
	}{
		{
			name: "basic substitution",
			input: TestStruct{
				SimpleField:  "hello",
				EnvField:     "$TEST_VAR",
				CommandField: "`echo hello`",
				MultiField:   "$TEST_VAR and `echo command`",
				EmptyField:   "",
				NestedStruct: Nested{
					NestedField:   "$NESTED_VAR",
					NestedCommand: "`echo nested`",
					unexported:    "should not change",
				},
				unexported: "should not change",
			},
			want: TestStruct{
				SimpleField:  "hello",
				EnvField:     "test_value",
				CommandField: "hello",
				MultiField:   "test_value and command",
				EmptyField:   "",
				NestedStruct: Nested{
					NestedField:   "nested_value",
					NestedCommand: "nested",
					unexported:    "should not change",
				},
				unexported: "should not change",
			},
			wantErr: false,
		},
		{
			name: "invalid command",
			input: TestStruct{
				CommandField: "`invalid_command_that_does_not_exist`",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SubstituteStringFields(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubstituteStringFields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubstituteStringFields() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestSubstituteStringFields_NonStruct(t *testing.T) {
	_, err := SubstituteStringFields("not a struct")
	if err == nil {
		t.Error("SubstituteStringFields() should return error for non-struct input")
	}
}

func TestSubstituteStringFields_NestedStructs(t *testing.T) {
	type DeepNested struct {
		Field string
	}

	type Nested struct {
		Field      string
		DeepNested DeepNested
	}

	type Root struct {
		Field  string
		Nested Nested
	}

	input := Root{
		Field: "$TEST_VAR",
		Nested: Nested{
			Field: "`echo nested`",
			DeepNested: DeepNested{
				Field: "$NESTED_VAR",
			},
		},
	}

	// Set up environment
	os.Setenv("TEST_VAR", "test_value")
	os.Setenv("NESTED_VAR", "deep_nested_value")
	defer os.Unsetenv("TEST_VAR")
	defer os.Unsetenv("NESTED_VAR")

	want := Root{
		Field: "test_value",
		Nested: Nested{
			Field: "nested",
			DeepNested: DeepNested{
				Field: "deep_nested_value",
			},
		},
	}

	got, err := SubstituteStringFields(input)
	if err != nil {
		t.Fatalf("SubstituteStringFields() error = %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("SubstituteStringFields() = %+v, want %+v", got, want)
	}
}

func TestSubstituteStringFields_EmptyStruct(t *testing.T) {
	type Empty struct{}

	input := Empty{}
	got, err := SubstituteStringFields(input)
	if err != nil {
		t.Fatalf("SubstituteStringFields() error = %v", err)
	}

	if !reflect.DeepEqual(got, input) {
		t.Errorf("SubstituteStringFields() = %+v, want %+v", got, input)
	}
}
